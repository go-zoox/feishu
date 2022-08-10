package access_token

import (
	"fmt"

	"github.com/go-zoox/fetch"
)

const tenantAccessTokenURI = "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal"

func GetTenantAccessToken(appID, appSecret string) (string, error) {
	response, err := fetch.Post(tenantAccessTokenURI, &fetch.Config{
		Headers: map[string]string{
			"Content-Type": "application/json; charset=utf-8",
		},
		Body: map[string]string{
			"app_id":     appID,
			"app_secret": appSecret,
		},
	})

	if err != nil {
		return "", err
	}

	// if response.Status != 200 {
	// 	return "", fmt.Errorf("[%d] %s", response.Status, response.String())
	// }

	code := response.Get("code").Int()
	if code != 0 {
		msg := response.Get("msg").String()
		return "", fmt.Errorf("[%d] %s", code, msg)
	}

	return response.Get("tenant_access_token").String(), nil
}
