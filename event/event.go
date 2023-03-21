package event

import (
	"github.com/go-zoox/core-utils/regexp"
	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/feishu/message"
	ct "github.com/go-zoox/feishu/message/content"
	"github.com/go-zoox/logger"
)

// docs: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/events/receive

var chatRe, _ = regexp.New("^im.message")

type Event interface {
	OnChatReceiveMessage(handler func(content string, request *EventRequest, reply func(content string, msgType ...string) error) error) error
}

type event struct {
	client  client.Client
	request *EventRequest
}

func New(client client.Client, request *EventRequest) Event {
	return &event{
		client:  client,
		request: request,
	}
}

func (e *event) OnChatReceiveMessage(handler func(content string, request *EventRequest, reply func(content string, msgType ...string) error) error) error {
	return e.request.OnChatReceiveMessage(e.client, handler)
}

type EventRequest struct {
	// Scheme is api version, e.g. 2.0
	Schema string             `json:"schema"`
	Header EventRequestHeader `json:"header"`
	Event  EventRequestBody   `json:"event"`

	// @TODO chanllenge (feishu bad design in the same message)
	// 未加密
	Challenge string `json:"challenge"`
	Token     string `json:"token"`
	Type      string `json:"type"`
	// 已加密
	Encrypt string `json:"encrypt"`
}

type EventResponse struct {
	Message string `json:"msg"`

	// @TODO chanllenge (feishu bad design in the same message)
	Challenge string `json:"challenge"`
}

type EventRequestHeader struct {
	// AppID is feishu app id, e.g. cli_123ab27db8200c
	AppID string `json:"app_id"`
	// CreateTime is event creation time, e.g. 1676566565810
	CreateTime string `json:"create_time"`
	// EventID is event id, e.g. 0e622f9d40f1752d282425b2b370b501
	EventID string `json:"event_id"`
	// EventType is event type, e.g. im.message.receive_v1
	EventType string `json:"event_type"`
	// TenantKey is tenant key, e.g. 70e62d17588cfc8f
	TenantKey string `json:"tenant_key"`
	// Token is token, e.g. KvcNNo641J123123Aehh
	Token string `json:"token"`
}

type EventRequestBody struct {
	// Sender is the message sender
	Sender  EventRequestSender `json:"sender"`
	Message struct {
		// MessageID is the message id, e.g. om_5ce6d572455d361153b7cb51da133945
		MessageID string `json:"message_id"`
		// MessageType is the message type, e.g. text
		MessageType string `json:"message_type"`
		// ChatID is the chat room id, e.g. oc_7a9aa4739f81bd2e61108fecbe12bf93
		ChatID string `json:"chat_id"`
		// ChatType is the chat type, options: group | p2p, e.g. group
		ChatType string `json:"chat_type"`
		// Content is message content, e.g. "{\"text\":\"啊实打实的 @_user_1 @_user_2\"}",
		Content string `json:"content"`
		// CreateTime is the creation time, e.g. 1676566565604
		CreateTime string `json:"create_time"`
		// Metions is the metions
		Mentions []EventRequestChatMention `json:"mentions"`
		// RootID is the root message id, e.g. om_5ce6d572455d361153b7cb5xxfsdfsdfdsf
		RootID string `json:"root_id"`
		// ParentID is the parent message id, e.g. om_5ce6d572455d361153b7cb5xxfsdfsdfdsf
		ParentID string `json:"parent_id"`
	} `json:"message"`
}

type EventRequestSender struct {
	SenderID struct {
		OpenID  string `json:"open_id"`
		UnionID string `json:"union_id"`
		UserID  string `json:"user_id"`
	} `json:"sender_id"`
	SenderType string `json:"sender_type"`
	TenantKey  string `json:"tenant_key"`
}

type EventRequestChatMention struct {
	// Key is the mention key, e.g. @_user_1
	Key string `json:"key"`
	// Name is the mention name, e.g. Zero
	Name string `json:"name"`
	//
	ID struct {
		OpenID  string `json:"open_id"`
		UnionID string `json:"union_id"`
		UserID  string `json:"user_id"`
	} `json:"id"`
	//
	TenantKey string `json:"tenant_key"`
}

func (e *EventRequest) IsChallenge() bool {
	return e.Challenge != ""
}

func (e *EventRequest) IsChat() bool {
	return chatRe.Match(e.EventType())
}

func (e *EventRequest) IsGroupChat() bool {
	return e.ChatType() == "group"
}

func (e *EventRequest) IsP2pChat() bool {
	return e.ChatType() == "p2p"
}

func (e *EventRequest) EventType() string {
	return e.Header.EventType
}

func (e *EventRequest) ChatType() string {
	return e.Event.Message.ChatType
}

func (e *EventRequest) ChatID() string {
	return e.Event.Message.ChatID
}

func (e *EventRequest) Sender() EventRequestSender {
	return e.Event.Sender
}

func (e *EventRequest) Mentions() []EventRequestChatMention {
	return e.Event.Message.Mentions
}

type MessageReply = func(content string, msgType ...string) error

type MessageHandler = func(content string, request *EventRequest, reply MessageReply) error

// 接收消息
//
// - 机器人接收到用户发送的消息后触发此事件。
//
// - 注意事项:;- 需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)，并订阅 ==消息与群组== 分类下的 ==接收消息v2.0== 事件才可接收推送;- 同时，将根据应用具备的权限，判断可推送的信息：;	- 当具备==获取用户发给机器人的单聊消息==权限或者==读取用户发给机器人的单聊消息（历史权限）==，可接收与机器人单聊会话中用户发送的所有消息;	- 当具备==获取群组中所有消息== 权限时，可接收与机器人所在群聊会话中用户发送的所有消息;	- 当具备==获取用户在群组中@机器人的消息== 权限或者==获取用户在群聊中@机器人的消息（历史权限）==，可接收机器人所在群聊中用户 @ 机器人的消息
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/events/receive
func (e *EventRequest) OnChatReceiveMessage(client client.Client, handler MessageHandler) error {
	if !e.IsChat() {
		return nil
	}

	if e.EventType() != "im.message.receive_v1" {
		return nil
	}

	return handler(e.Event.Message.Content, e, func(content string, msgType ...string) error {
		if content == "" {
			return nil
		}

		msgTypeX, err := inferMessageTypeFromContent(content, msgType...)
		if err != nil {
			return err
		}

		// _, err = message.Send(client, &message.SendRequest{
		// 	ReceiveIDType: "chat_id",
		// 	ReceiveID:     e.ChatID(),
		// 	MsgType:       msgTypeX,
		// 	Content:       content,
		// })
		// if err != nil && len(msgType) == 0 {
		// 	logger.Warnf("failed to send message: %v, maybe not infer msgType correctly, you need set msgType", err)
		// }

		if e.IsP2pChat() {
			_, err = message.Send(client, &message.SendRequest{
				ReceiveIDType: "chat_id",
				ReceiveID:     e.ChatID(),
				MsgType:       msgTypeX,
				Content:       content,
			})
		} else {
			_, err = message.Reply(client, &message.ReplyRequest{
				MessageID: e.Event.Message.MessageID,
				MsgType:   msgTypeX,
				Content:   content,
			})
		}
		if err != nil && len(msgType) == 0 {
			logger.Warnf("failed to reply message: %v, maybe not infer msgType correctly, you need set msgType", err)
		}

		return err
	})
}

// 机器人进群
//
// - 机器人被用户添加至群聊时触发此事件。
//
// - 注意事项：;- 需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability);- 需要订阅 ==消息与群组== 分类下的 ==机器人进群== 事件;- 事件会向进群的机器人进行推送;- 机器人邀请机器人不会触发事件
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-member-bot/events/added
func (e *EventRequest) OnChatBotAddedToGroup(client client.Client, handler MessageHandler) error {
	if !e.IsChat() {
		return nil
	}

	if e.EventType() != "im.chat.member.bot.added_v1" {
		return nil
	}

	return handler(e.Event.Message.Content, e, func(content string, msgType ...string) error {
		if content == "" {
			return nil
		}

		msgTypeX, err := inferMessageTypeFromContent(content, msgType...)
		if err != nil {
			return err
		}

		_, err = message.Reply(client, &message.ReplyRequest{
			MessageID: e.ChatID(),
			MsgType:   msgTypeX,
			Content:   content,
		})
		return err
	})
}

// 机器人被移出群
//
// - 机器人被移出群聊后触发此事件。
//
// - 注意事项：;- 需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability);- 需要订阅 ==消息与群组== 分类下的 ==机器人被移出群== 事件;- 事件会向被移出群的机器人进行推送
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-member-bot/events/deleted
func (e *EventRequest) OnChatBotDeletedFromGroup(client client.Client, handler MessageHandler) error {
	if !e.IsChat() {
		return nil
	}

	if e.EventType() != "im.chat.member.bot.deleted_v1" {
		return nil
	}

	return handler(e.Event.Message.Content, e, func(content string, msgType ...string) error {
		if content == "" {
			return nil
		}

		msgTypeX, err := inferMessageTypeFromContent(content, msgType...)
		if err != nil {
			return err
		}

		_, err = message.Send(client, &message.SendRequest{
			ReceiveID:     e.ChatID(),
			Content:       content,
			MsgType:       msgTypeX,
			ReceiveIDType: "chat_id",
		})
		return err
	})
}

func inferMessageTypeFromContent(content string, msgType ...string) (string, error) {
	if len(msgType) != 0 && msgType[0] != "" {
		return msgType[0], nil
	}

	return ct.InferContentMsgType(content)
}
