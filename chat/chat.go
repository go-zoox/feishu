package chat

import "github.com/go-zoox/feishu/client"

type Chat interface {
	Create(req *CreateRequest) (resp *CreateResponse, err error)
	List(cfg *ListRequest) (resp *ListResponse, err error)

	//
	ListALL() (chats []ChatEntity, err error)
}

type chat struct {
	client client.Client
}

func New(c client.Client) Chat {
	return &chat{
		client: c,
	}
}

func (i *chat) Create(cfg *CreateRequest) (resp *CreateResponse, err error) {
	return Create(i.client, cfg)
}

func (i *chat) List(cfg *ListRequest) (resp *ListResponse, err error) {
	return List(i.client, cfg)
}

func (i *chat) ListALL() (chats []ChatEntity, err error) {
	return ListALL(i.client)
}

type ChatEntity struct {
	ChatID      string `json:"chat_id"`
	Avatar      string `json:"avatar"`
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerID     string `json:"owner_id"`
	OwnerIDType string `json:"owner_id_type"`
	External    bool   `json:"external"`
	TenantKey   string `json:"tenant_key"`
}
