package chat

import (
	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/fetch"
)

// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/delete

const RemoveMemberAPI = "https://open.feishu.cn/open-apis/im/v1/chats/:chat_id/members"

type RemoveMemberRequest struct {
	// 群 ID
	ChatID string

	// 进群成员 ID 类型 open_id/user_id/union_id/app_id
	// 注意：拉机器人入群请使用 app_id
	// 示例值："open_id"
	MemberIDType string

	// 成员ID列表，获取ID请参见如何获得 User ID、Open ID 和 Union ID？
	IDList []string
}

type RemoveMemberResponse struct {
	InvalidIDList []string `json:"invalid_id_list"`
}

func RemoveMember(client client.Client, req *RemoveMemberRequest) (resp *RemoveMemberResponse, err error) {
	err = client.Request(RemoveMemberAPI, &fetch.Config{
		Method: "DELETE",
		Headers: map[string]string{
			"Content-Type": "application/json; charset=utf-8",
		},
		Params: map[string]string{
			"chat_id": req.ChatID,
		},
		Query: map[string]string{
			"member_id_type": req.MemberIDType,
		},
		Body: map[string]any{
			"id_list": req.IDList,
		},
	}, &resp)
	return
}
