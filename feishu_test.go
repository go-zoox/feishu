package feishu

import (
	"os"
	"testing"

	"github.com/go-zoox/core-utils/fmt"
	_ "github.com/go-zoox/dotenv"
	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/feishu/contact/user"
)

func TestContactUserWithOpenID(t *testing.T) {

	c := New(&client.Config{
		AppID:     os.Getenv("APP_ID"),
		AppSecret: os.Getenv("APP_SECRET"),
	})

	u, err := c.Contact().User().Retrieve(&user.RetrieveRequest{
		UserIDType: "open_id",
		UserID:     os.Getenv("OPEN_ID"),
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON(u)
}

func TestContactUserWithUnionID(t *testing.T) {

	c := New(&client.Config{
		AppID:     os.Getenv("APP_ID"),
		AppSecret: os.Getenv("APP_SECRET"),
	})

	u, err := c.Contact().User().Retrieve(&user.RetrieveRequest{
		UserIDType: "union_id",
		UserID:     os.Getenv("UNION_ID"),
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON(u)
}

func TestContactUserWithUserID(t *testing.T) {

	c := New(&client.Config{
		AppID:     os.Getenv("APP_ID"),
		AppSecret: os.Getenv("APP_SECRET"),
	})

	u, err := c.Contact().User().Retrieve(&user.RetrieveRequest{
		UserIDType: "user_id",
		UserID:     os.Getenv("USER_ID"),
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON(u)
}
