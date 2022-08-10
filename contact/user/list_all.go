package user

import (
	"fmt"

	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/feishu/contact/department"
)

type ListAllResponse struct {
	// 部门列表
	Items []UserEntity `json:"items"`
}

func ListAll(client client.Client) (users []UserEntity, err error) {
	items := []UserEntity{}
	departments, err := department.ListAll(client)
	if err != nil {
		return nil, fmt.Errorf("list department failed: %s", err)
	}

	for _, department := range departments.Items {
		childUsers, err := listAllChild(client, department.DeparntmentID)
		if err != nil {
			return nil, fmt.Errorf("list user failed: %s", err)
		}

		items = append(items, childUsers...)
	}

	return items, nil
}

func listAllChild(client client.Client, departmentID string) ([]UserEntity, error) {
	var items []UserEntity
	var pageToken string
	var hasMore bool = true

	for hasMore {
		l, err := List(client, &ListRequest{
			DepartmentIDType: "department_id",
			DepartmentID:     departmentID,
			PageToken:        pageToken,
		})
		if err != nil {
			return nil, err
		}

		items = append(items, l.Items...)

		pageToken = l.PageToken
		hasMore = l.HasMore
	}

	return items, nil
}
