package webhook

import (
	"fmt"

	"github.com/go-zoox/feishu/message/content"
	"github.com/go-zoox/fetch"
)

const BaseURI = "https://open.feishu.cn/open-apis/bot/v2/hook"

// WebHook is the feishu WebHook SDK.
type WebHook interface {
	Send(text string) error
	SendText(text string) error
	SendPost(post content.ContentTypePost) error
	SendShareChat(chatID string) error
	SendImage(imageKey string) error
	SendCard(card content.ContentTypeInteractive) error
}

type webhook struct {
	token string
}

func New(token string) WebHook {
	return &webhook{token: token}
}

func (w *webhook) URL() string {
	return fmt.Sprintf("%s/%s", BaseURI, w.token)
}

func (w *webhook) Send(text string) error {
	return w.SendText(text)
}

func (w *webhook) SendText(text string) error {
	return w.trigger(map[string]any{
		"msg_type": "text",
		"content": map[string]any{
			"text": text,
		},
	})
}

func (w *webhook) SendPost(post content.ContentTypePost) error {
	return w.trigger(map[string]any{
		"msg_type": "post",
		"content": map[string]any{
			"post": post,
		},
	})
}

func (w *webhook) SendShareChat(chatID string) error {
	return w.trigger(map[string]any{
		"msg_type": "share_chat",
		"content": map[string]any{
			"share_chat_id": chatID,
		},
	})
}

func (w *webhook) SendImage(imageKey string) error {
	return w.trigger(map[string]any{
		"msg_type": "image",
		"content": map[string]any{
			"image_key": imageKey,
		},
	})
}

func (w *webhook) SendCard(card content.ContentTypeInteractive) error {
	return w.trigger(map[string]any{
		"msg_type": "image",
		"card":     card,
	})
}

func (w *webhook) trigger(data map[string]any) error {
	resp, err := fetch.Post(w.URL(), &fetch.Config{
		Headers: fetch.Headers{
			"Content-Type": "application/json",
		},
		Body: data,
	})
	if err != nil {
		return err
	}

	if !resp.Ok() {
		return fmt.Errorf("failed to trigger webhook: [%d] %s", resp.Get("code").Int(), resp.Get("msg").String())
	}

	return nil
}
