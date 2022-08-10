package feishu

import (
	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/feishu/message"
)

type Client interface {
	// GetAppAccessToken() (string, error)
	// GetTenanttAccessToken() (string, error)

	Message() message.Message
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
