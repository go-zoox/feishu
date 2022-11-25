package chat

import (
	"os"
	"testing"

	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/dotenv"
	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/feishu/config"
)

func TestUpdateAnnouncement(t *testing.T) {
	var cfg config.Config
	if err := dotenv.Load(&cfg); err != nil {
		t.Fatal(err)
	}

	c := client.New(&client.Config{
		BaseURI:   cfg.BaseURI,
		AppID:     cfg.AppID,
		AppSecret: cfg.ApptSecret,
	})

	fmt.Println("chat_id:", os.Getenv("CHAT_ID"))

	announcement, err := UpdateAnnouncement(c, &UpdateAnnouncementRequest{
		ChatID:   os.Getenv("CHAT_ID"),
		Revision: "0",
		Content: []string{
			"测试1: line1",
			"测试2: line 2",
			"测试3: line 3",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON("announcement:", announcement)
}
