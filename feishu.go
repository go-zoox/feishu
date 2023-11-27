package feishu

import (
	"github.com/go-zoox/feishu/bot"
	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/feishu/contact"
	"github.com/go-zoox/feishu/event"
	"github.com/go-zoox/feishu/image"
	"github.com/go-zoox/feishu/message"
)

type Client interface {
	// GetAppAccessToken() (string, error)
	// GetTenanttAccessToken() (string, error)

	//
	Bot() bot.Bot

	//
	Contact() contact.Contact
	Message() message.Message
	Event(request *event.EventRequest) event.Event
	Image() image.Image
}

type Config = client.Config

type feishu struct {
	sdk client.Client
}

func New(cfg *client.Config) Client {
	sdk := client.New(cfg)
	return &feishu{sdk: sdk}
}

func (c *feishu) Message() message.Message {
	return message.New(c.sdk)
}

func (c *feishu) Contact() contact.Contact {
	return contact.New(c.sdk)
}

func (c *feishu) Event(request *event.EventRequest) event.Event {
	return event.New(c.sdk, request)
}

func (c *feishu) Bot() bot.Bot {
	return bot.New(c.sdk)
}

func (c *feishu) Image() image.Image {
	return image.New(c.sdk)
}
