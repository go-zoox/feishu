package chat

import (
	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/fetch"
)

const CreateAPI = "https://open.feishu.cn/open-apis/im/v1/chats"

type CreateRequest struct {
	// 群名称
	Name string

	// 群描述
	Description string

	// 群头像对应的 Image Key，可通过上传图片获取（注意：上传图片的 image_type 需要指定为 avatar）
	// 示例值："default-avatar_44ae0ca3-e140-494b-956f-78091e348435"
	// 上传图片：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/create
	Avatar string

	// 创建群时指定的群主，不填时指定建群的机器人为群主。
	// 群主 ID，ID值与查询参数中的 user_id_type 对应。
	OwnerID string

	// 创建群时邀请的群成员，id 类型为 user_id_type
	// 最大长度：50
	UserIDList []string

	// 创建群时邀请的群机器人
	// 最大长度：5
	BotIDList []string

	// 群模式，可选值：
	// group
	ChatMode string

	// 群类型，可选值：
	// private - 私有群
	// public - 公开群
	ChatType string
}

type CreateResponse struct {
	ChatID      string `json:"chat_id"`
	Avatar      string `json:"avatar"`
	Name        string `json:"name"`
	Description string `json:"description"`
	I18nNames   struct {
		ZhCN string `json:"zh_cn"`
		EnUS string `json:"en_us"`
		JaJP string `json:"ja_jp"`
	}
	OwnerID                string `json:"owner_id"`
	AddMemberPermission    string `json:"add_member_permission"`
	ShareCardPermission    string `json:"share_card_permission"`
	AtAllPermission        string `json:"at_all_permission"`
	ChatMode               string `json:"chat_mode"`
	ChatType               string `json:"chat_type"`
	ChatTag                string `json:"chat_tag"`
	External               bool   `json:"external"`
	TenantKey              string `json:"tenant_key"`
	JoinMessageVisibility  string `json:"join_message_visibility"`
	LeaveMessageVisibility string `json:"leave_message_visibility"`
	MembershipApproval     string `json:"membership_approval"`
	ModerationPermission   string `json:"moderation_permission"`
}

func Create(client client.Client, req *CreateRequest) (resp *CreateResponse, err error) {
	err = client.Request(CreateAPI, &fetch.Config{
		Method: "POST",
		Body: map[string]any{
			"name":         req.Name,
			"description":  req.Description,
			"avatar":       req.Avatar,
			"owner_id":     req.OwnerID,
			"user_id_list": req.UserIDList,
			"bot_id_list":  req.BotIDList,
			"chat_mode":    req.ChatMode,
			"chat_type":    req.ChatType,
		},
	}, &resp)
	return

	// response, err := fetch.Post(CreateAPI, &fetch.Config{
	// 	Headers: map[string]string{
	// 		"Content-Type":  "application/json; charset=utf-8",
	// 		"Authorization": fmt.Sprintf("Bearer %s", accessToken),
	// 	},
	// 	Body: map[string]any{
	// 		"name":         req.Name,
	// 		"description":  req.Description,
	// 		"avatar":       req.Avatar,
	// 		"owner_id":     req.OwnerID,
	// 		"user_id_list": req.UserIDList,
	// 		"bot_id_list":  req.BotIDList,
	// 		"chat_mode":    req.ChatMode,
	// 		"chat_type":    req.ChatType,
	// 	},
	// })

	// if err != nil {
	// 	return nil, err
	// }

	// // if response.Status != 200 {
	// // 	return nil, fmt.Errorf("[%d] %s", response.Status, response.String())
	// // }

	// code := response.Get("code").Int()
	// if code != 0 {
	// 	msg := response.Get("msg").String()
	// 	return nil, fmt.Errorf("[%d] %s", code, msg)
	// }

	// var res CreateResponse
	// if err := json.Unmarshal([]byte(response.Get("data").String()), &res); err != nil {
	// 	return nil, fmt.Errorf("json unmarshal: %s", err)
	// }

	// return &res, nil
}
