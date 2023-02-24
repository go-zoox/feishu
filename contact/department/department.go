package department

import "github.com/go-zoox/feishu/client"

type Department interface {
	List(cfg *ListRequest) (resp *ListResponse, err error)
	Retrieve(cfg *RetrieveRequest) (resp *RetrieveResponse, err error)

	//
	ListAll() (resp *ListAllResponse, err error)
}

type department struct {
	client client.Client
}

func New(c client.Client) Department {
	return &department{
		client: c,
	}
}

func (d *department) List(cfg *ListRequest) (resp *ListResponse, err error) {
	return List(d.client, cfg)
}

func (d *department) Retrieve(cfg *RetrieveRequest) (resp *RetrieveResponse, err error) {
	return Retrieve(d.client, cfg)
}

func (d *department) ListAll() (resp *ListAllResponse, err error) {
	return ListAll(d.client)
}
