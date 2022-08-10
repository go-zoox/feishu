package department

import (
	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/fetch"
)

const RetrieveAPI = "/open-apis/contact/v3/departments/:department_id"

type RetrieveRequest struct {
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
}

type RetrieveResponse struct {
	// 部门信息
	Department DepartmentEntity `json:"department"`
}

func Retrieve(client client.Client, cfg *RetrieveRequest) (resp *RetrieveResponse, err error) {
	err = client.Request(RetrieveAPI, &fetch.Config{
		Method: "GET",
		Params: map[string]string{
			"department_id": cfg.DepartmentID,
		},
		Query: map[string]string{
			"user_id_type":       cfg.UserIDType,
			"department_id_type": cfg.DepartmentIDType,
		},
	}, &resp)
	return
}
