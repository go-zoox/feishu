package image

import (
	"io"
	"os"

	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/fetch"
)

// doc: https://open.feishu.cn/document/server-docs/im-v1/image/get
const DownloadAPI = "/open-apis/im/v1/images/:image_key"

type DownloadRequest struct {
	// ImageKey 图片的 Key
	ImageKey string `json:"image_key"`
}

func Download(client client.Client, imageKey string, imageFilePath string) (err error) {
	fetch, err := client.CreateFetch(DownloadAPI, &fetch.Config{
		Method: "GET",
		Params: map[string]string{
			"image_key": imageKey,
		},
		IsStream: true,
	})
	if err != nil {
		return err
	}

	response, err := fetch.Send()
	if err != nil {
		return err
	}
	defer response.Stream.Close()

	f, err := os.OpenFile(imageFilePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	if _, err := io.Copy(f, response.Stream); err != nil {
		return err
	}

	return nil
}
