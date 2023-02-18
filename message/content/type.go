package content

// ContentType is the message content type interface.
type ContentType interface {
	ContentTypeText | ContentTypePost | ContentTypeImage | ContentTypeInteractive | ContentTypeShareChat | ContentTypeShareUser | ContentTypeAudio | ContentTypeMedia | ContentTypeFile | ContentTypeSticker
}

// docs: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json#c9e08671
type ContentTypeText struct {
	Text string `json:"text"`
}

// docs: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json#45e0953e
type ContentTypePost struct {
	ZhCN *ContentTypePostBody `json:"zh_cn,omitempty"`
	EnUS *ContentTypePostBody `json:"en_us,omitempty"`
}

// docs: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json#7111df05
type ContentTypeImage struct {
	ImageKey string `json:"image_key"`
}

// docs:
//   - 消息卡片 interactive: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json#11e75d0
//   - 配置卡片属性: https://open.feishu.cn/document/ukTMukTMukTM/uAjNwUjLwYDM14CM2ATN
//   - 消息卡片搭建工具: https://open.feishu.cn/tool/cardbuilder?templateId=ctp_AAfWvP5y4jp2
type ContentTypeInteractive struct {
	Config   ContentTypeInteractiveConfig    `json:"config"`
	Elements []ContentTypeInteractiveElement `json:"elements"`
}

// docs: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json#55c61488
type ContentTypeShareChat struct {
	ChatID string `json:"chat_id"`
}

// docs: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json#38563ae5
type ContentTypeShareUser struct {
	UserID string `json:"user_id"`
}

// docs: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json#5d353271
type ContentTypeAudio struct {
	FileKey string `json:"file_key"`
}

// docs: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json#54406d84
type ContentTypeMedia struct {
	FileKey  string `json:"file_key"`
	ImageKey string `json:"image_key"`
}

// docs: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json#c92e6d46
type ContentTypeFile struct {
	FileKey string `json:"file_key"`
}

// docs: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json#7215e4f6
type ContentTypeSticker struct {
	FileKey string `json:"file_key"`
}

type ContentTypePostBody struct {
	Title   string                      `json:"title,omitempty"`
	Content [][]ContentTypePostBodyItem `json:"content,omitempty"`
}

type ContentTypePostBodyItem struct {
	// text | a | at | img
	Tag string `json:"tag"`

	// tag = text and a
	Text string `json:"text,omitempty"`

	// tag = a
	Href string `json:"href,omitempty"`

	// tag = at
	UserID   string `json:"user_id,omitempty"`
	UserName string `json:"user_name,omitempty"`

	// tag = img | media
	ImageKey string `json:"image_key,omitempty"`

	// tag = media
	FileKey string `json:"file_key,omitempty"`

	// tag = emotion
	EmojiType string `json:"emoji_type,omitempty"`
}

type ContentTypeInteractiveConfig struct {
	EnableForward  bool `json:"enable_forward"`
	UpdateMulti    bool `json:"update_multi"`
	WideScreenMode bool `json:"wide_screen_mode"`
}

type ContentTypeInteractiveElement struct {
	Tag     string `json:"tag"`
	Content string `json:"content"`
	Href    struct {
		UrlVal struct {
			Url string `json:"url"`
		} `json:"urlVal"`
	} `json:"href"`
	At struct {
		Content string `json:"content"`
		Tag     string `json:"tag"`
	} `json:"at"`
	ImgKey string `json:"img_key"`
	Extra  struct {
		Alt struct {
			Content string `json:"content"`
			Tag     string `json:"tag"`
		} `json:"alt"`
		ImgKey string `json:"img_key"`
		Tag    string `json:"tag"`
	} `json:"extra"`
	Text struct {
		Content string `json:"content"`
		Tag     string `json:"tag"`
	} `json:"text"`
	Actions []struct {
		Tag  string `json:"tag"`
		Text struct {
			Content string `json:"content"`
			Tag     string `json:"tag"`
		} `json:"text"`
		Type string `json:"type"`
		Url  string `json:"url"`
	} `json:"actions"`
	Elements []struct {
		Content string `json:"content"`
		Tag     string `json:"tag"`
	} `json:"elements"`
}
