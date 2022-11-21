package chat

import (
	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/feishu/message"
)

type SendMessageRequest struct {
	// 群 ID
	ChatID string `json:"chat_id"`

	// 消息类型 包括：text、post、image、file、audio、media、sticker、interactive、share_chat、share_user等
	// 类型定义请参考发送消息content说明（https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json）
	// 示例值："text"
	MsgType string `json:"msg_type"`

	// 消息内容，json结构序列化后的字符串。不同msg_type对应不同内容。消息类型 包括：text、post、image、file、audio、media、sticker、interactive、share_chat、share_user等
	// 具体格式说明参考：发送消息content说明（https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json）
	Content string `json:"content"`
}

type SendMessageResponse = message.SendResponse

func SendMessage(client client.Client, req *SendMessageRequest) (resp *SendMessageResponse, err error) {
	return message.Send(client, &message.SendRequest{
		ReceiveIDType: "chat_id",
		ReceiveID:     req.ChatID,
		MsgType:       req.MsgType,
		Content:       req.Content,
	})
}
