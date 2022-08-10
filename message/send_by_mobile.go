package message

import (
	"fmt"

	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/feishu/contact/user"
)

type SendByMobileRequest struct {
	// User Mobile
	Mobile string `json:"mobile"`

	// 消息类型 包括：text、post、image、file、audio、media、sticker、interactive、share_chat、share_user等
	// 类型定义请参考发送消息content说明（https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json）
	// 示例值："text"
	MsgType string `json:"msg_type"`

	// 消息内容，json结构序列化后的字符串。不同msg_type对应不同内容。消息类型 包括：text、post、image、file、audio、media、sticker、interactive、share_chat、share_user等
	// 具体格式说明参考：发送消息content说明（https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json）
	Content string `json:"content"`
}

type SendByMobileResponse = SendResponse

func SendByMobile(client client.Client, cfg *SendByMobileRequest) (*SendByMobileResponse, error) {
	unionID, err := user.GetUnionIDByMobile(client, cfg.Mobile)
	if err != nil {
		return nil, fmt.Errorf("get union id by mobile(%s): %w", cfg.Mobile, err)
	}

	resp, err := Send(client, &SendRequest{
		ReceiveIDType: "union_id",
		ReceiveID:     unionID,
		Content:       cfg.Content,
		MsgType:       cfg.MsgType,
	})
	if err != nil {
		return nil, fmt.Errorf("send message by mobile(union_id: %s): %w", unionID, err)
	}

	return resp, nil
}
