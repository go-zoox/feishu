package content

import (
	"encoding/json"
	"strings"
)

type ContentBuilder interface {
}

type contentBuilder struct {
	msgType string
	content any // ContentType
}

func NewContent() ContentBuilder {
	return &contentBuilder{}
}

func (c *contentBuilder) Build() (msgType, content string, err error) {
	contentByte, err := json.Marshal(content)
	if err != nil {
		return "", "", err
	}

	return c.msgType, string(contentByte), nil
}

func (c *contentBuilder) Text(v *ContentTypeText) ContentBuilder {
	c.msgType = "text"
	c.content = v
	return c
}

func (c *contentBuilder) Post(v *ContentTypePost) ContentBuilder {
	c.msgType = "post"
	c.content = v
	return c
}

func (c *contentBuilder) Image(v *ContentTypeImage) ContentBuilder {
	c.msgType = "image"
	c.content = v
	return c
}

func (c *contentBuilder) Interactive(v *ContentTypeInteractive) ContentBuilder {
	c.msgType = strings.ToLower("Interactive")
	c.content = v
	return c
}

func (c *contentBuilder) ShareChat(v *ContentTypeShareChat) ContentBuilder {
	c.msgType = "share_chat"
	c.content = v
	return c
}

func (c *contentBuilder) ShareUser(v *ContentTypeShareUser) ContentBuilder {
	c.msgType = "share_user"
	c.content = v
	return c
}

func (c *contentBuilder) Audio(v *ContentTypeAudio) ContentBuilder {
	c.msgType = "audio"
	c.content = v
	return c
}

func (c *contentBuilder) Media(v *ContentTypeMedia) ContentBuilder {
	c.msgType = "media"
	c.content = v
	return c
}

func (c *contentBuilder) File(v *ContentTypeFile) ContentBuilder {
	c.msgType = "file"
	c.content = v
	return c
}

func (c *contentBuilder) Sticker(v *ContentTypeSticker) ContentBuilder {
	c.msgType = "sticker"
	c.content = v
	return c
}
