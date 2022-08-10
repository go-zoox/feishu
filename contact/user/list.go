package user

import (
	"fmt"

	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/fetch"
)

const ListAPI = "/open-apis/contact/v3/users/find_by_department"

type ListRequest struct {
	// 部门 ID
	DepartmentID string `json:"department_id"`

	// 用户 ID 类型
	// 示例值："user_id"
	// 可选值："open_id" "user_id" "union_id"
	UserIDType string `json:"user_id_type"`

	// 部门 ID 类型
	// 示例值："open_department_id"
	// 可选值："open_department_id" "department_id"
	DepartmentIDType string `json:"department_id_type"`

	// 分页大小
	PageSize int `json:"page_size"`

	// 分页标记
	// 第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果
	PageToken string `json:"page_token"`
}

type ListResponse struct {
	// 是否有更多选项
	HasMore bool `json:"has_more"`

	// 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
	PageToken string `json:"page_token"`

	// 部门列表
	Items []UserEntity `json:"items"`
}

func List(client client.Client, cfg *ListRequest) (resp *ListResponse, err error) {
	query := fetch.Query{
		"department_id":      cfg.DepartmentID,
		"user_id_type":       cfg.UserIDType,
		"department_id_type": cfg.DepartmentIDType,
		// "page_size":          fmt.Sprintf("%d", cfg.PageSize),
		"page_token": cfg.PageToken,
	}

	if cfg.PageSize > 0 {
		query.Set("page_size", fmt.Sprintf("%d", cfg.PageSize))
	}

	err = client.Request(ListAPI, &fetch.Config{
		Method: fetch.GET,
		Query:  query,
	}, &resp)
	return
}
