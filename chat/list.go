package chat

import (
	"fmt"

	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/fetch"
)

// doc: https://open.feishu.cn/document/server-docs/group/chat/list
const ListAPI = "/open-apis/im/v1/chats"

type ListRequest struct {
	// UserIDType 用户 ID 类型
	// 可选值：
	// 	open_id（标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同）
	//	union_id（标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的，在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID，应用开发商可以把同个用户在多个应用中的身份关联起来）
	//  user_id（标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内，一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据）
	// 示例值：open_id
	// 默认值: open_id
	// 必填：否
	UserIDType string `json:"user_id_type"`

	// SortType 群组排序方式
	// 可选值:
	//	ByCreateTimeAsc：按群组创建时间升序排列
	//	ByActiveTimeDesc：按群组活跃时间降序排列
	// 示例值：ByCreateTimeAsc
	// 默认值: ByCreateTimeAsc
	// 必填：否
	SortType string `json:"sort_type"`

	// PageToken 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果
	// 示例值：dmJCRHhpd3JRbGV1VEVNRFFyTitRWDY5ZFkybmYrMEUwMUFYT0VMMWdENEtuYUhsNUxGMDIwemtvdE5ORjBNQQ
	// 必填：否
	PageToken string `json:"page_token"`

	// PageSize 分页大小，最大为 100
	// 示例值：10
	// 默认值: 20
	// 必填：否
	PageSize uint `json:"page_size"`
}

type ListResponse struct {
	Items []ChatEntity `json:"items"`

	// HasMore 是否还有更多项
	HasMore bool `json:"has_more"`

	// PageToken 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
	PageToken string `json:"page_token"`
}

func List(client client.Client, cfg *ListRequest) (resp *ListResponse, err error) {
	err = client.Request(ListAPI, &fetch.Config{
		Method: "GET",
		Query: fetch.Query{
			"user_id_type": cfg.UserIDType,
			"sort_type":    cfg.SortType,
			"page_token":   cfg.PageToken,
			"page_size":    fmt.Sprintf("%v", cfg.PageSize),
		},
	}, &resp)
	return
}
