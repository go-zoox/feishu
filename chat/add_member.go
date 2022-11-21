package chat

import (
	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/fetch"
)

// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/create

const AddMemberAPI = "https://open.feishu.cn/open-apis/im/v1/chats/:chat_id/members"

type AddMemberRequest struct {
	// 群 ID
	ChatID string

	// 进群成员 ID 类型 open_id/user_id/union_id/app_id
	// 注意：拉机器人入群请使用 app_id
	// 示例值："open_id"
	MemberIDType string

	// 出现不可用ID后的处理方式 0/1/2
	// 示例值：0
	SucceedType string

	// 成员ID列表，获取ID请参见如何获得 User ID、Open ID 和 Union ID？
	IDList []string
}

type AddMemberResponse struct {
	InvalidIDList    []string `json:"invalid_id_list"`
	NotExistedIDList []string `json:"not_existed_id_list"`
}

func AddMember(client client.Client, req *AddMemberRequest) (resp *AddMemberResponse, err error) {
	err = client.Request(AddMemberAPI, &fetch.Config{
		Method: "POST",
		Params: map[string]string{
			"chat_id": req.ChatID,
		},
		Query: map[string]string{
			"member_id_type": req.MemberIDType,
			"succeed_type":   req.SucceedType,
		},
		Body: map[string]any{
			"id_list": req.IDList,
		},
	}, &resp)
	return
}
