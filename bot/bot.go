package bot

import "github.com/go-zoox/feishu/client"

type Bot interface {
	GetBotInfo() (resp *GetBotInfoResponse, err error)
}

type bot struct {
	client client.Client
}

func New(c client.Client) Bot {
	return &bot{
		client: c,
	}
}

func (m *bot) GetBotInfo() (resp *GetBotInfoResponse, err error) {
	return GetBotInfo(m.client)
}
