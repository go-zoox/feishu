package message

import (
	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/fetch"
)

const SendAPI = "/open-apis/im/v1/messages"

type SendRequest struct {
	// 依据receive_id_type的值，填写对应的消息接收者id
	// 示例值："ou_7d8a6e6df7621556ce0d21922b676706ccs"
	ReceiveID string `json:"receive_id"`

	// 消息内容，json结构序列化后的字符串。不同msg_type对应不同内容。消息类型 包括：text、post、image、file、audio、media、sticker、interactive、share_chat、share_user等
	// 具体格式说明参考：发送消息content说明（https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json）
	Content string `json:"content"`

	// 消息类型 包括：text、post、image、file、audio、media、sticker、interactive、share_chat、share_user等
	// 类型定义请参考发送消息content说明（https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json）
	// 示例值："text"
	MsgType string `json:"msg_type"`

	// 消息接收者id类型 open_id/user_id/union_id/email/chat_id
	// 示例值："open_id"
	ReceiveIDType string `json:"receive_id_type"`
}

type SendResponse struct {
	MessageEntity
}

func Send(client client.Client, cfg *SendRequest) (resp *SendResponse, err error) {
	err = client.Request(SendAPI, &fetch.Config{
		Method: "POST",
		Headers: map[string]string{
			"Content-Type": "application/json; charset=utf-8",
		},
		Query: map[string]string{
			"receive_id_type": cfg.ReceiveIDType,
		},
		Body: map[string]interface{}{
			"receive_id": cfg.ReceiveID,
			"content":    cfg.Content,
			"msg_type":   cfg.MsgType,
		},
	}, &resp)
	return
}
