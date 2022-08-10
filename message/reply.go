package message

import (
	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/fetch"
)

const ReplyAPI = "/open-apis/im/v1/messages/:message_id/reply"

type ReplyRequest struct {
	// 待回复的消息的ID
	// 示例值："om_dc13264520392913993dd051dba21dcf"
	MessageID string `json:"message_id"`

	// 消息内容，json结构序列化后的字符串。不同msg_type对应不同内容。消息类型 包括：text、post、image、file、audio、media、sticker、interactive、share_chat、share_user等
	// 具体格式说明参考：发送消息content说明（https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json）
	Content string `json:"content"`

	// 消息类型 包括：text、post、image、file、audio、media、sticker、interactive、share_chat、share_user等
	// 类型定义请参考发送消息content说明（https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json）
	// 示例值："text"
	MsgType string `json:"msg_type"`
}

type ReplyResponse struct {
	MessageEntity
}

func Reply(client client.Client, cfg *ReplyRequest) (resp *ReplyResponse, err error) {
	err = client.Request(ReplyAPI, &fetch.Config{
		Method: "POST",
		Params: map[string]string{
			"message_id": cfg.MessageID,
		},
		Body: map[string]interface{}{
			"content":  cfg.Content,
			"msg_type": cfg.MsgType,
		},
	}, &resp)
	return
}
