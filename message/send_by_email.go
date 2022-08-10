package message

import (
	"fmt"

	"github.com/go-zoox/feishu/client"
)

type SendByEmailRequest struct {
	// User Email
	Email string `json:"email"`

	// 消息类型 包括：text、post、image、file、audio、media、sticker、interactive、share_chat、share_user等
	// 类型定义请参考发送消息content说明（https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json）
	// 示例值："text"
	MsgType string `json:"msg_type"`

	// 消息内容，json结构序列化后的字符串。不同msg_type对应不同内容。消息类型 包括：text、post、image、file、audio、media、sticker、interactive、share_chat、share_user等
	// 具体格式说明参考：发送消息content说明（https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json）
	Content string `json:"content"`
}

type SendByEmailResponse = SendResponse

func SendByEmail(client client.Client, cfg *SendByEmailRequest) (*SendByEmailResponse, error) {
	// unionID, err := user.GetUnionIDByEmail(client, os.Getenv("SEND_BY_EMAIL"))
	// if err != nil {
	// 	return nil, fmt.Errorf("get union id by email: %w", err)
	// }

	resp, err := Send(client, &SendRequest{
		ReceiveIDType: "email",
		ReceiveID:     cfg.Email,
		// ReceiveIDType: "union_id",
		// ReceiveID:     unionID,
		Content: cfg.Content,
		MsgType: cfg.MsgType,
	})
	if err != nil {
		return nil, fmt.Errorf("send message by email: %w", err)
	}

	return resp, nil
}
