package chat

import (
	"time"

	"github.com/go-zoox/feishu/client"
)

func ListALL(client client.Client) (chats []ChatEntity, err error) {
	pageToken := ""

	for {
		response, err := List(client, &ListRequest{
			PageToken: pageToken,
			PageSize:  100,
		})
		if err != nil {
			return nil, err
		}

		chats = append(chats, response.Items...)
		if !response.HasMore {
			break
		}

		pageToken = response.PageToken

		time.Sleep(1 * time.Second)
	}

	return
}
