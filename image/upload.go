package image

import (
	"io"

	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/fetch"
)

// doc: https://open.feishu.cn/document/server-docs/im-v1/image/create
const UploadAPI = "/open-apis/im/v1/images"

type UploadRequest struct {
	// ImageType 图片类型
	// 可选值: message（用于发送消息)，avatar（用于设置头像）
	// 示例值: message
	// 必填: 是
	ImageType string `json:"image_type"`

	// Image 图片内容
	// 注意：上传的图片大小不能超过 10MB
	// 示例值: 二进制文件
	Image io.Reader `json:"image"`
}

type UploadResponse struct {
	// ImageKey 图片的 Key
	ImageKey string `json:"image_key"`
}

func Upload(client client.Client, cfg *UploadRequest) (resp *UploadResponse, err error) {
	err = client.Request(UploadAPI, &fetch.Config{
		Method: "POST",
		Headers: map[string]string{
			"Content-Type": "multipart/form-data",
		},
		Body: map[string]interface{}{
			"image_type": cfg.ImageType,
			"image":      cfg.Image,
		},
	}, &resp)
	return
}
