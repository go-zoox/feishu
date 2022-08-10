package user

import (
	"testing"

	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/dotenv"
	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/feishu/config"
)

func TestRetrieve(t *testing.T) {
	var cfg config.Config
	if err := dotenv.Load(&cfg); err != nil {
		t.Fatal(err)
	}

	c := client.New(&client.Config{
		BaseURI:   cfg.BaseURI,
		AppID:     cfg.AppID,
		AppSecret: cfg.ApptSecret,
	})

	user, err := Retrieve(c, &RetrieveRequest{
		// UserIDType: "user_id",
		// UserID:     "3gf84g45",
		// //
		UserIDType: "union_id",
		UserID:     "on_37417f8ee5f36529088034fe274172be",
		//
		// UserIDType: "open_id",
		// UserID:     "ou_5bda2db8af1fb71fb75a5c1de92d680d",
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON(user)
}
