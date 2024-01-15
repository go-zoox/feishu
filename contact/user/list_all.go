package user

import (
	"fmt"

	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/feishu/contact/department"
	"github.com/go-zoox/logger"
)

type ListALLRequest struct {
	DepartmentIDs []string `json:"department_ids"`
}

type ListALLResponse struct {
	// 部门列表
	Items []UserEntity `json:"items"`
}

func ListALL(client client.Client, cfg *ListALLRequest) (users []UserEntity, err error) {
	items := []UserEntity{}

	if cfg.DepartmentIDs == nil {
		logger.Infof("[contact.user.list_all] list all departments ...")
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
		logger.Infof("[contact.user.list_all] list all departments (%d) done", len(departments.Items))
	}

	for _, departmentID := range cfg.DepartmentIDs {
		logger.Infof("[contact.user.list_all] list users of department %s ...", departmentID)
		childUsers, err := listAllChild(client, departmentID)
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

		fmt.Println("pageToken:", pageToken, "hasMore:", hasMore)
	}

	return items, nil
}
