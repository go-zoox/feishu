package message

import (
	"encoding/json"
	"testing"

	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/dotenv"
	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/feishu/config"
)

func TestSend(t *testing.T) {
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

	department, err := Send(c, &SendRequest{
		Content:       string(dx),
		MsgType:       "text",
		ReceiveIDType: "open_id",
		ReceiveID:     "ou_5bda2db8af1fb71fb75a5c1de92d680d",
		// ReceiveIDType: "email",
		// ReceiveID: "",
		// ReceiveIDType: "user_id",
		// ReceiveID:     "3gf84g45",
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON("send message response:", department)
}
