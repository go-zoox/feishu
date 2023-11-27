package user

import (
	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/fetch"
)

const SearchAPI = "/open-apis/contact/v3/users/batch_get_id"

type SearchRequest struct {
	// 用户 ID 类型
	// 示例值："user_id"
	// 可选值："open_id" "user_id" "union_id"
	UserIDType string `json:"user_id_type"`

	// 要查询的用户邮箱，最多 50 条。
	Emails []string `json:"emails"`

	// 要查询的用户手机号，最多 50 条。 非中国大陆地区的手机号需要添加以 “+” 开头的国家 / 地区代码
	Mobiles []string `json:"mobiles"`
}

type SearchResponse struct {
	// 手机号或者邮箱对应的用户id信息
	UserList []struct {
		// 用户id，值为user_id_type所指定的类型。如果查询的手机号、邮箱不存在，或者无权限查看对应的用户，则此项为空
		UserId string `json:"user_id"`

		// 手机号，通过手机号查询时返回
		Mobile string `json:"mobile"`

		// 邮箱，通过邮箱查询时返回
		Email string `json:"email"`
	} `json:"user_list"`
}

func Search(client client.Client, cfg *SearchRequest) (resp *SearchResponse, err error) {
	body := make(map[string]interface{})
	if len(cfg.Emails) > 0 {
		body["emails"] = cfg.Emails
	}
	if len(cfg.Mobiles) > 0 {
		body["mobiles"] = cfg.Mobiles
	}

	err = client.Request(SearchAPI, &fetch.Config{
		Method: fetch.POST,
		Headers: map[string]string{
			"Content-Type": "application/json; charset=utf-8",
		},
		Query: map[string]string{
			"user_id_type": cfg.UserIDType,
		},
		Body: body,
	}, &resp)
	return
}
