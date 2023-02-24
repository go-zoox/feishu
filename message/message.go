package message

import "github.com/go-zoox/feishu/client"

type Message interface {
	Send(cfg *SendRequest) (resp *SendResponse, err error)
	SendByEmail(cfg *SendByEmailRequest) (resp *SendByEmailResponse, err error)
	SendByMobile(cfg *SendByMobileRequest) (resp *SendByMobileResponse, err error)
}

type message struct {
	client client.Client
}

func New(c client.Client) Message {
	return &message{
		client: c,
	}
}

func (m *message) Send(cfg *SendRequest) (resp *SendResponse, err error) {
	return Send(m.client, cfg)
}

func (m *message) SendByEmail(cfg *SendByEmailRequest) (resp *SendByEmailResponse, err error) {
	return SendByEmail(m.client, cfg)
}

func (m *message) SendByMobile(cfg *SendByMobileRequest) (resp *SendByMobileResponse, err error) {
	return SendByMobile(m.client, cfg)
}
