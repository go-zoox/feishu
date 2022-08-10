package user

import (
	"fmt"

	"github.com/go-zoox/feishu/client"
)

func GetUnionIDByEmail(client client.Client, email string) (string, error) {
	users, err := Search(client, &SearchRequest{
		UserIDType: "union_id",
		Emails:     []string{email},
	})
	if err != nil {
		return "", err
	}

	if len(users.UserList) == 0 || users.UserList[0].UserId == "" {
		return "", fmt.Errorf("user not found or no permission to view")
	}

	return users.UserList[0].UserId, nil
}
