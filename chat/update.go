package chat

import (
	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/fetch"
)

const UpdateAPI = "https://open.feishu.cn/open-apis/im/v1/chats/:chat_id"

type UpdateRequest struct {
	// 群 ID
	ChatID string `json:"chat_id"`

	// 用户 ID 类型
	// 示例值："open_id"
	// 可选值有：
	// open_id：用户的 open id
	// union_id：用户的 union id
	// user_id：用户的 user id
	UserIDType string `json:"user_id_type"`

	// 群名称
	Name string `json:"name"`

	// 群描述
	Description string `json:"description"`

	// 群头像对应的 Image Key，可通过上传图片获取（注意：上传图片的 image_type 需要指定为 avatar）
	// 示例值："default-avatar_44ae0ca3-e140-494b-956f-78091e348435"
	// 上传图片：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/create
	Avatar string `json:"avatar"`

	// 邀请用户或机器人入群权限
	// 可选值有：
	// only_owner：仅群主和管理员
	// all_members：所有成员
	AddMemberPermission string `json:"add_member_permission"`

	// 群分享权限
	// 可选值有：
	// allowed：允许
	// not_allowed：不允许
	ShareCardPermission string `json:"share_card_permission"`

	// at 所有人权限
	// 可选值有：
	// only_owner：仅群主和管理员
	// all_members：所有成员
	// 示例值："all_members"
	AtAllPermission string `json:"at_all_permission"`

	// 群编辑权限
	// 可选值有：
	// only_owner：仅群主和管理员
	// all_members：所有成员
	// 示例值："all_members"
	EditPermission string `json:"edit_permission"`

	// 创建群时指定的群主，不填时指定建群的机器人为群主。
	// 群主 ID，ID值与查询参数中的 user_id_type 对应。
	OwnerID string

	// 入群消息可见性
	// 可选值有：
	// only_owner：仅群主和管理员可见
	// all_members：所有成员可见
	// not_anyone：任何人均不可见
	// 示例值："only_owner"
	JoinMessageVisibility string `json:"join_message_visibility"`

	// 出群消息可见性
	// 可选值有：
	// only_owner：仅群主和管理员可见
	// all_members：所有成员可见
	// not_anyone：任何人均不可见
	// 示例值："only_owner"
	LeaveMessageVisibility string `json:"leave_message_visibility"`

	// 加群审批
	// 可选值有：
	// no_approval_required：无需审批
	// approval_required：需要审批
	// 示例值："no_approval_required"
	MembershipApproval string `json:"membership_approval"`

	// 群类型，可选值：
	// private - 私有群
	// public - 公开群
	ChatType string
}

type UpdateResponse struct {
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
	ChatType               string `json:"chat_type"`
	JoinMessageVisibility  string `json:"join_message_visibility"`
	LeaveMessageVisibility string `json:"leave_message_visibility"`
	MembershipApproval     string `json:"membership_approval"`
}

func Update(client client.Client, req *UpdateRequest) (resp *UpdateResponse, err error) {
	err = client.Request(UpdateAPI, &fetch.Config{
		Method: "PUT",
		Headers: map[string]string{
			"Content-Type": "application/json; charset=utf-8",
		},
		Params: map[string]string{
			"chat_id": req.ChatID,
		},
		Query: map[string]string{
			"user_id_type": req.UserIDType,
		},
		Body: map[string]any{
			"name":                  req.Name,
			"description":           req.Description,
			"avatar":                req.Avatar,
			"add_member_permission": req.AddMemberPermission,
			"share_card_permission": req.ShareCardPermission,
			"at_all_permission":     req.AtAllPermission,
			"edit_permission":       req.EditPermission,
			// "owner_id":                 req.OwnerID,
			"join_message_visibility":  req.JoinMessageVisibility,
			"leave_message_visibility": req.LeaveMessageVisibility,
			"membership_approval":      req.MembershipApproval,
			"chat_type":                req.ChatType,
		},
	}, &resp)
	return

	// response, err := fetch.Post(UpdateAPI, &fetch.Config{
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

	// var res UpdateResponse
	// if err := json.Unmarshal([]byte(response.Get("data").String()), &res); err != nil {
	// 	return nil, fmt.Errorf("json unmarshal: %s", err)
	// }

	// return &res, nil
}
