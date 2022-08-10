package user

import (
	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/fetch"
)

const RetrieveAPI = "/open-apis/contact/v3/users/:user_id"

type RetrieveRequest struct {
	// 用户ID，需要与查询参数中的user_id_type类型保持一致。
	UserID string `json:"user_id"`

	// 用户 ID 类型
	// 示例值："user_id"
	// 可选值："open_id" "user_id" "union_id"
	UserIDType string `json:"user_id_type"`

	// 部门 ID 类型
	// 示例值："open_department_id"
	// 可选值："open_department_id" "department_id"
	DepartmentIDType string `json:"department_id_type"`
}

type RetrieveResponse struct {
	// 用户信息
	User UserEntity `json:"user"`
}

func Retrieve(client client.Client, cfg *RetrieveRequest) (resp *RetrieveResponse, err error) {
	err = client.Request(RetrieveAPI, &fetch.Config{
		Method: "GET",
		Params: map[string]string{
			"user_id": cfg.UserID,
		},
		Query: map[string]string{
			"user_id_type":       cfg.UserIDType,
			"department_id_type": cfg.DepartmentIDType,
		},
	}, &resp)
	return
}
