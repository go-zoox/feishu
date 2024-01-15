package chat

import "github.com/go-zoox/feishu/client"

type Chat interface {
	Create(req *CreateRequest) (resp *CreateResponse, err error)
	List(cfg *ListRequest) (resp *ListResponse, err error)
	Update(cfg *UpdateRequest) (resp *UpdateResponse, err error)

	//
	ListALL() (chats []ChatEntity, err error)

	AddMember(cfg *AddMemberRequest) (resp *AddMemberResponse, err error)
	RemoveMember(cfg *RemoveMemberRequest) (resp *RemoveMemberResponse, err error)
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

func (i *chat) Update(cfg *UpdateRequest) (resp *UpdateResponse, err error) {
	return Update(i.client, cfg)
}

func (i *chat) ListALL() (chats []ChatEntity, err error) {
	return ListALL(i.client)
}

func (i *chat) AddMember(cfg *AddMemberRequest) (resp *AddMemberResponse, err error) {
	return AddMember(i.client, cfg)
}

func (i *chat) RemoveMember(cfg *RemoveMemberRequest) (resp *RemoveMemberResponse, err error) {
	return RemoveMember(i.client, cfg)
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
