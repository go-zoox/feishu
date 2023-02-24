package contact

import (
	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/feishu/contact/department"
	"github.com/go-zoox/feishu/contact/user"
)

type Contact interface {
	User() user.User
	Department() department.Department
}

type contact struct {
	client client.Client
}

func New(c client.Client) Contact {
	return &contact{
		client: c,
	}
}

func (c *contact) User() user.User {
	return user.New(c.client)
}

func (c *contact) Department() department.Department {
	return department.New(c.client)
}
