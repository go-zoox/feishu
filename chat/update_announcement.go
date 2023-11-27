package chat

import (
	"encoding/json"

	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/fetch"
)

const UpdateAnnouncementAPI = "https://open.feishu.cn/open-apis/im/v1/chats/:chat_id/announcement"

type UpdateAnnouncementRequest struct {
	// 待获取公告的群 ID，详情参见群ID 说明
	// 注意：不支持P2P单聊
	// 示例值："oc_5ad11d72b830411d72b836c20"
	ChatID string `json:"chat_id"`

	// 文档当前版本号 int64 类型，获取群公告信息接口会返回
	// 示例值："12"
	Revision string `json:"revision"`

	// 修改文档请求的序列化字段
	// 更新公告信息的格式和更新云文档格式相同
	Requests []string `json:"requests"`

	Content []string `json:"content"`
}

type UpdateAnnouncementResponse struct {
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

func UpdateAnnouncement(client client.Client, req *UpdateAnnouncementRequest) (resp *UpdateAnnouncementResponse, err error) {
	if req.Content != nil {
		req.Requests, err = generateRequestsFromContent(req.Content)
		if err != nil {
			return nil, err
		}

		fmt.Println("requests: ", req.Requests)
	}

	err = client.Request(UpdateAnnouncementAPI, &fetch.Config{
		Method: "PATCH",
		Headers: map[string]string{
			"Content-Type": "application/json; charset=utf-8",
		},
		Params: map[string]string{
			"chat_id": req.ChatID,
		},
		Body: map[string]any{
			"revision": req.Revision,
			"requests": req.Requests,
		},
	}, &resp)
	return
}

func generateRequestsFromContent(content []string) ([]string, error) {
	lines := []map[string]any{}
	for _, line := range content {
		lines = append(lines, map[string]any{
			"type": "paragraph",
			"paragraph": map[string]any{
				"elements": []map[string]any{
					{
						"type": "textRun",
						"textRun": map[string]any{
							"style": map[string]any{},
							"text":  line,
						},
					},
				},
				"style": map[string]any{},
			},
		})
	}

	payload := map[string]any{
		"blocks": lines,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	newContent := map[string]any{
		"requestType": "InsertBlocksRequestType",
		"insertBlocksRequest": map[string]any{
			"location": map[string]any{
				"zoneId":      "0",
				"index":       0,
				"startOfZone": true,
				// "endOfZone":   true,
			},
			"payload": string(payloadBytes),
		},
	}
	newContentBytes, err := json.Marshal(newContent)
	if err != nil {
		return nil, err
	}

	deleteRequest := map[string]any{
		"requestType": "DeleteContentRangeRequestType",
		"deleteContentRangeRequest": map[string]any{
			"deleteRange": map[string]any{
				"zoneId":     "0",
				"startIndex": 0,
				"endIndex":   111,
				// "index":     0,
				// "endOfZone": true,
			},
		},
	}
	deleteRequestBytes, err := json.Marshal(deleteRequest)
	if err != nil {
		return nil, err
	}

	return []string{string(deleteRequestBytes), string(newContentBytes)}, nil
	// fmt.Println(string(deleteRequestBytes))
	// return []string{string(newContentBytes)}, nil
}
