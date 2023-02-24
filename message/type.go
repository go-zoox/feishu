package message

type MessageEntity struct {
	// 消息 ID
	MessageID string `json:"message_id"`

	// 根消息 ID
	RootID string `json:"root_id"`

	// 父消息 ID
	ParentID string `json:"parent_id"`

	// 消息类型 包括：text、post、image、file、audio、media、sticker、interactive、share_chat、share_user等
	MsgType string `json:"msg_type"`

	// 消息生成的时间戳（毫秒）
	CreateTime string `json:"create_time"`

	// 消息更新的时间戳（毫秒）
	UpdateTime string `json:"update_time"`

	// 消息是否被撤回
	Deleted bool `json:"deleted"`

	// 消息是否被更新
	Updated bool `json:"updated"`

	// 所属的群
	ChatID string `json:"chat_id"`

	// 消息发送者
	Sender MessageEntitySender `json:"sender"`

	// 消息内容
	Body MessageEntityBody `json:"body"`

	// 被@的用户或机器人的id列表
	Mentions []MessageEntityMention `json:"mentions"`

	// 合并转发消息中，上一层级的消息id message_id
	UpperMessageID string `json:"upper_message_id"`
}

type MessageEntitySender struct {
	// 该字段标识发送者的id
	ID string `json:"id"`

	// 该字段标识发送者的id类型
	IDType string `json:"id_type"`

	// 该字段标识发送者的类型
	SenderType string `json:"sender_type"`

	// 为租户在飞书上的唯一标识，用来换取对应的tenant_access_token，也可以用作租户在应用里面的唯一标识
	TenantID string `json:"tenant_id"`
}

type MessageEntityBody struct {
	// 消息内容，json结构序列化后的字符串。不同msg_type对应不同内容。消息类型 包括：text、post、image、file、audio、media、sticker、interactive、share_chat、share_user等
	Content interface{} `json:"content"`
}

type MessageEntityMention struct {
	// 被@的用户或机器人的序号。例如，第3个被@到的成员，值为“@_user_3”
	Key string `json:"key"`

	// 被@的用户或者机器人的open_id
	ID string `json:"id"`

	// 该字段标识发送者的id类型
	IDType string `json:"id_type"`

	// 被@的用户或机器人的姓名
	Name string `json:"name"`

	// 为租户在飞书上的唯一标识，用来换取对应的tenant_access_token，也可以用作租户在应用里面的唯一标识
	TenantKey string `json:"tenant_key"`
}
