package user

import (
	"os"
	"testing"

	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/dotenv"
	"github.com/go-zoox/feishu/client"
	"github.com/go-zoox/feishu/config"
)

func TestList(t *testing.T) {
	var cfg config.Config
	if err := dotenv.Load(&cfg); err != nil {
		t.Fatal(err)
	}

	c := client.New(&client.Config{
		BaseURI:   cfg.BaseURI,
		AppID:     cfg.AppID,
		AppSecret: cfg.ApptSecret,
	})

	users, err := List(c, &ListRequest{
		DepartmentID: os.Getenv("DEPARTMENT_ID"),
		// UserIDType:   "user_id",
		DepartmentIDType: "department_id",
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON("user list:", users)
}
