package chat

import (
	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/fetch"
)

const GetAnnouncementAPI = "https://open.feishu.cn/open-apis/im/v1/chats/:chat_id/announcement"

type GetAnnouncementRequest struct {
	// 待获取公告的群 ID，详情参见群ID 说明
	// 注意：不支持P2P单聊
	// 示例值："oc_5ad11d72b830411d72b836c20"
	ChatID string `json:"chat_id"`

	// 用户 ID 类型
	// 示例值："open_id"
	// 可选值有：
	// open_id：用户的 open id
	// union_id：用户的 union id
	// user_id：用户的 user id
	// 默认值：open_id
	UserIDType string `json:"user_id_type"`
}

type GetAnnouncementResponse struct {
	// 云文档序列化信息
	Content string `json:"content"`

	// 文档当前版本号 纯数字
	Revision string `json:"revision"`

	// 文档生成的时间戳（秒）
	CreateTime string `json:"create_time"`

	// 文档更新时间戳（秒）
	UpdateTime string `json:"update_time"`

	// 文档所有者的 ID 类型
	// 如果所有者是用户，则与查询参数中的user_id_type 相同；取值为open_id user_id union_id 其中之一，不同 ID 的说明参见 用户相关的 ID 概念
	// 如果所有者是机器人，为机器人应用的 app_id，详情参见 获取应用身份访问凭证
	// 可选值有：
	// user_id：以 user_id 来识别用户
	// union_id：以 union_id 来识别用户
	// open_id：以 open_id 来识别用户
	// app_id：以 app_id 来识别机器人应用
	OwnerIDType string `json:"owner_id_type"`

	// 文档所有者 ID，ID 值与owner_id_type 中的ID类型对应
	OwnerID string `json:"owner_id"`

	// 文档最新修改者 id 类型
	// 如果修改者是用户，则与查询参数中的user_id_type 相同；取值为open_id user_id union_id 其中之一，不同 ID 的说明参见 用户相关的 ID 概念
	// 如果修改者是机器人，为机器人应用的 app_id，详情参见 获取应用身份访问凭证
	// 可选值有：
	// user_id：以 user_id 来识别用户
	// union_id：以 union_id 来识别用户
	// open_id：以 open_id 来识别用户
	// app_id：以 app_id 来识别应用
	ModifierIDType string `json:"modifier_id_type"`

	// 文档最新修改者 ID，ID 值与modifier_id_type 中的ID类型对应
	ModifierID string `json:"modifier_id"`
}

func GetAnnouncement(client client.Client, req *GetAnnouncementRequest) (resp *GetAnnouncementResponse, err error) {
	err = client.Request(GetAnnouncementAPI, &fetch.Config{
		Method: "GET",
		Headers: map[string]string{
			"Content-Type": "application/json; charset=utf-8",
		},
		Params: map[string]string{
			"chat_id": req.ChatID,
		},
		Query: map[string]string{
			"user_id_type": req.UserIDType,
		},
	}, &resp)
	return
}
