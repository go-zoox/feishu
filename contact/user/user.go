package user

import "github.com/go-zoox/feishu/client"

type User interface {
	List(cfg *ListRequest) (resp *ListResponse, err error)
	Retrieve(cfg *RetrieveRequest) (resp *RetrieveResponse, err error)

	//
	ListAll() (users []UserEntity, err error)
	Search(cfg *SearchRequest) (resp *SearchResponse, err error)
	GetUnionIDByEmail(email string) (string, error)
	GetUnionIDByMobile(mobile string) (string, error)
}

type user struct {
	client client.Client
}

func New(c client.Client) User {
	return &user{
		client: c,
	}
}

func (u *user) List(cfg *ListRequest) (resp *ListResponse, err error) {
	return List(u.client, cfg)
}

func (u *user) Retrieve(cfg *RetrieveRequest) (resp *RetrieveResponse, err error) {
	return Retrieve(u.client, cfg)
}

func (u *user) ListAll() (users []UserEntity, err error) {
	return ListAll(u.client)
}

func (u *user) Search(cfg *SearchRequest) (resp *SearchResponse, err error) {
	return Search(u.client, cfg)
}

func (u *user) GetUnionIDByEmail(email string) (string, error) {
	return GetUnionIDByEmail(u.client, email)
}

func (u *user) GetUnionIDByMobile(mobile string) (string, error) {
	return GetUnionIDByMobile(u.client, mobile)
}
