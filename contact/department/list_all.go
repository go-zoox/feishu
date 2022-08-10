package department

import (
	"github.com/go-zoox/feishu/client"
)

type ListAllResponse struct {
	// 部门列表
	Items []DepartmentEntity `json:"items"`
}

func ListAll(client client.Client) (resp *ListAllResponse, err error) {
	items, err := listChild(client, "0")
	if err != nil {
		return nil, err
	}

	return &ListAllResponse{
		Items: items,
	}, nil
}

func listChild(client client.Client, departmentID string) ([]DepartmentEntity, error) {
	var items []DepartmentEntity
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
		for _, item := range l.Items {
			childItems, err := listChild(client, item.DeparntmentID)
			if err != nil {
				return nil, err
			}

			items = append(items, childItems...)
		}

		pageToken = l.PageToken
		hasMore = l.HasMore
	}

	return items, nil
}
