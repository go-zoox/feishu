package client

import (
	"encoding/json"
	"fmt"

	"github.com/go-zoox/feishu/access_token"
	"github.com/go-zoox/fetch"
)

type Client interface {
	Request(resource string, request *fetch.Config, response interface{}) error
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

func (c *client) Request(resource string, request *fetch.Config, response interface{}) error {
	if err := c.refreshAccessToken(); err != nil {
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
	if code == 99991663 {
		err = c.refreshAccessToken()
		if err != nil {
			return fmt.Errorf("refresh access token failed(2): %s", err)
		}

		resp, err = client.Retry(func(f *fetch.Fetch) {
			f.SetBearToken(c.accessToken)
		})
		if err != nil {
			return fmt.Errorf("retry failed: %s", err)
		}
	}

	code = resp.Get("code").Int()
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
	} else if resp.Status != 200 {
		return fmt.Errorf("[status: %d] %s", resp.Status, resp.String())
	}

	if err := json.Unmarshal([]byte(resp.Get("data").String()), response); err != nil {
		return fmt.Errorf("parse response failed: %s", err)
	}

	return nil
}

func (c *client) refreshAccessToken() (err error) {
	if c.accessToken != "" {
		return
	}

	c.accessToken, err = access_token.GetTenantAccessToken(c.cfg.AppID, c.cfg.AppSecret)
	return
}
