package message

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/dotenv"
	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/feishu/config"
)

func TestReply(t *testing.T) {
	var cfg config.Config
	if err := dotenv.Load(&cfg); err != nil {
		t.Fatal(err)
	}

	c := client.New(&client.Config{
		BaseURI:   cfg.BaseURI,
		AppID:     cfg.AppID,
		AppSecret: cfg.ApptSecret,
	})

	d := struct {
		Text string `json:"text"`
	}{
		Text: "test from go-zoox",
	}

	dx, _ := json.Marshal(d)

	department, err := Reply(c, &ReplyRequest{
		MessageID: os.Getenv("TEST_MESSAGE_ID"),
		Content:   string(dx),
		MsgType:   "text",
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON("reply message response:", department)
}
