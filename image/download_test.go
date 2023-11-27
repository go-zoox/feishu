package image

import (
	"testing"

	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/dotenv"
	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/feishu/config"
	"github.com/go-zoox/fs"
)

func TestDownload(t *testing.T) {
	var cfg config.Config
	if err := dotenv.Load(&cfg); err != nil {
		t.Fatal(err)
	}

	c := client.New(&client.Config{
		BaseURI:   cfg.BaseURI,
		AppID:     cfg.AppID,
		AppSecret: cfg.ApptSecret,
	})

	imagePath := fs.TmpFilePath()
	imageKey := "img_v3_025j_6e599774-ae13-4577-8ace-cffb62d7578g"

	err := Download(c, imageKey, imagePath)
	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON("image download response:", map[string]interface{}{
		"image_path": imagePath,
		"image_key":  imageKey,
	})
}
