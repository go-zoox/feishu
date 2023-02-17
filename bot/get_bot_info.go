package bot

import (
	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/fetch"
)

const GetBotInfoAPI = "/open-apis/bot/v3/info"

type GetBotInfoResponse struct {
	ActivateStatus int      `json:"activate_status"`
	AppName        string   `json:"app_name"`
	AvatarURL      string   `json:"avatar_url"`
	IPWhiteList    []string `json:"ip_white_list"`
	OpenID         string   `json:"open_id"`
}

func GetBotInfo(client client.Client) (resp *GetBotInfoResponse, err error) {
	err = client.Request(GetBotInfoAPI, &fetch.Config{
		Method: "GET",
	}, &resp, "bot")
	return
}
