package department

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

	department, err := Retrieve(c, &RetrieveRequest{
		DepartmentID:     "ba5cb33332bf6c4a",
		DepartmentIDType: "department_id",
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON("department:", department)
}
