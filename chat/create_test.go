package chat

import (
	"testing"

	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/dotenv"
	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/feishu/config"
)

func TestCreate(t *testing.T) {
	// var cfg config.Config
	// if err := dotenv.Load(&cfg); err != nil {
	// 	t.Fatal(err)
	// }

	// fmt.PrintJSON("cfg", cfg)

	// token, err := access_token.GetTenantAccessToken(cfg.AppID, cfg.ApptSecret)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// fmt.Println("token:", token)

	var cfg config.Config
	if err := dotenv.Load(&cfg); err != nil {
		t.Fatal(err)
	}

	c := client.New(&client.Config{
		BaseURI:   cfg.BaseURI,
		AppID:     cfg.AppID,
		AppSecret: cfg.ApptSecret,
	})

	chat, err := Create(c, &CreateRequest{
		Name: "test_from_go_zoox",
		// OwnerId: "3gf84g45",
		UserIDList: []string{
			// "3gf84g45",
			"ou_5bda2db8af1fb71fb75a5c1de92d680d",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON("chat:", chat)
}
