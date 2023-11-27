package image

import (
	"os"
	"testing"

	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/dotenv"
	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/feishu/config"
)

func TestUpload(t *testing.T) {
	var cfg config.Config
	if err := dotenv.Load(&cfg); err != nil {
		t.Fatal(err)
	}

	c := client.New(&client.Config{
		BaseURI:   cfg.BaseURI,
		AppID:     cfg.AppID,
		AppSecret: cfg.ApptSecret,
	})

	image, _ := os.Open("/tmp/generated_00.png")

	response, err := Upload(c, &UploadRequest{
		ImageType: "message",
		Image:     image,
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON("image upload response:", response)
}
