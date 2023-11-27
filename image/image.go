package image

import "github.com/go-zoox/feishu/client"

type Image interface {
	Upload(cfg *UploadRequest) (resp *UploadResponse, err error)
}

type image struct {
	client client.Client
}

func New(c client.Client) Image {
	return &image{
		client: c,
	}
}

func (i *image) Upload(cfg *UploadRequest) (resp *UploadResponse, err error) {
	return Upload(i.client, cfg)
}
