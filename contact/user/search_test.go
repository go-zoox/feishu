package user

import (
	"os"
	"testing"

	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/dotenv"
	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/feishu/config"
)

func TestSearch(t *testing.T) {
	var cfg config.Config
	if err := dotenv.Load(&cfg); err != nil {
		t.Fatal(err)
	}

	c := client.New(&client.Config{
		BaseURI:   cfg.BaseURI,
		AppID:     cfg.AppID,
		AppSecret: cfg.ApptSecret,
	})

	user, err := Search(c, &SearchRequest{
		// UserIDType: "open_id",
		UserIDType: "union_id",
		// UserIDType: "user_id",
		// Mobiles: []string{os.Getenv("TEST_MOBILE")},
		Emails: []string{os.Getenv("TEST_EMAIL")},
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON(user)
}
