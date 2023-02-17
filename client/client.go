package client

import (
	"encoding/json"
	"fmt"

	"github.com/go-zoox/feishu/access_token"
	"github.com/go-zoox/fetch"
	"github.com/go-zoox/logger"
)

type Client interface {
	Request(resource string, request *fetch.Config, response interface{}, dataKey ...string) error
}

type Config struct {
	BaseURI   string
	AppID     string
	AppSecret string
}

func New(cfg *Config) Client {
	if cfg.BaseURI == "" {
		cfg.BaseURI = "https://open.feishu.cn"
	}

	return &client{
		cfg: cfg,
	}
}

type client struct {
	cfg *Config
	//
	accessToken string
}

func (c *client) Request(resource string, request *fetch.Config, response interface{}, dataKey ...string) error {
	dataKeyX := "data"
	if len(dataKey) > 0 {
		dataKeyX = dataKey[0]
	}

	if err := c.refreshAccessToken(nil); err != nil {
		return fmt.Errorf("refresh access token failed(1): %s", err)
	}

	client := fetch.
		Create(c.cfg.BaseURI).
		SetConfig(request).
		SetMethod(request.Method).
		SetURL(resource).
		SetContentType("application/json; charset=utf-8").
		SetBearToken(c.accessToken)

	resp, err := client.Send()
	if err != nil {
		return fmt.Errorf("fetch failed: %s", err)
	}

	code := resp.Get("code").Int()
	// https://open.feishu.cn/document/ukTMukTMukTM/ugjM14COyUjL4ITN
	// 99991663 | tenant token invalid | tenant_access_token 无效，更新token
	// 99991668 | token invalid | user_access_token无效，更新token。详情可参考API访问凭证概述
	// 99991668 | user access token not support | 当前请求不支持user_access_token，请检查后重试

	switch code {
	case 99991663, 99991668:
		err = c.refreshAccessToken(fmt.Errorf("invalid access token or token expired"))
		if err != nil {
			return fmt.Errorf("refresh access token failed(2): %s", err)
		}

		resp, err = client.Retry(func(f *fetch.Fetch) {
			f.SetBearToken(c.accessToken)
		})
		if err != nil {
			return fmt.Errorf("retry failed: %s", err)
		}

		code = resp.Get("code").Int()
	}

	if code != 0 {
		msg := resp.Get("msg").String()
		var detail string
		errorFieldViolations := resp.Get("error.field_violations").String()
		if errorFieldViolations != "" {
			var fieldViolations []struct {
				Field       string `json:"field"`
				Description string `json:"description"`
			}

			if err := json.Unmarshal([]byte(errorFieldViolations), &fieldViolations); err == nil {
				for _, fieldViolation := range fieldViolations {
					detail += fmt.Sprintf("%s: %s; ", fieldViolation.Field, fieldViolation.Description)
				}
			}

			return fmt.Errorf("[%d] %s(detail: %s)", code, msg, detail)
		}

		return fmt.Errorf("[%d] %s", code, msg)
	}

	if resp.Status != 200 {
		return fmt.Errorf("[status: %d] %s", resp.Status, resp.String())
	}

	if err := json.Unmarshal([]byte(resp.Get(dataKeyX).String()), response); err != nil {
		return fmt.Errorf("parse response failed: %s", err)
	}

	return nil
}

func (c *client) refreshAccessToken(reason error) (err error) {
	if reason == nil && c.accessToken != "" {
		return
	}

	logger.Infof("[feishu] refresh access token by reason: %s", reason)

	c.accessToken, err = access_token.GetTenantAccessToken(c.cfg.AppID, c.cfg.AppSecret)
	return
}
