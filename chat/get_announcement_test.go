package chat

import (
	"os"
	"testing"

	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/dotenv"
	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/feishu/config"
)

func TestGetAnnouncement(t *testing.T) {
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

	announcement, err := GetAnnouncement(c, &GetAnnouncementRequest{
		ChatID: os.Getenv("CHAT_ID"),
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON("announcement:", announcement)
}
