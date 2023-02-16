package event

type EventRequest struct {
	// Scheme is api version, e.g. 2.0
	Schema string             `json:"schema"`
	Header EventRequestHeader `json:"header"`
	Event  EventRequestBody   `json:"event"`
}

type EventResponse struct {
	Message string `json:"msg"`
}

type EventRequestHeader struct {
	// AppID is feishu app id, e.g. cli_123ab27db8200c
	AppID string `json:"app_id"`
	// CreateTime is event creation time, e.g. 1676566565810
	CreateTime string `json:"create_time"`
	// EventID is event id, e.g. 0e622f9d40f1752d282425b2b370b501
	EventID string `json:"event_id"`
	// EventType is event type, e.g. im.message.receive_v1
	EventType string `json:"event_type"`
	// TenantKey is tenant key, e.g. 70e62d17588cfc8f
	TenantKey string `json:"tenant_key"`
	// Token is token, e.g. KvcNNo641J123123Aehh
	Token string `json:"token"`
}

type EventRequestBody struct {
	// Sender is the message sender
	Sender struct {
		SenderID struct {
			OpenID  string `json:"open_id"`
			UnionID string `json:"union_id"`
			UserID  string `json:"user_id"`
		}
		SenderType string `json:"sender_type"`
		TenantKey  string `json:"tenant_key"`
	} `json:"sender"`
	Message struct {
		// ChatID is the chat room id, e.g. oc_7a9aa4739f81bd2e61108fecbe12bf93
		ChatID string `json:"chat_id"`
		// ChatType is the chat type, options: group | p2p, e.g. group
		ChatType string `json:"chat_type"`
		// Content is message content, e.g. "{\"text\":\"啊实打实的 @_user_1 @_user_2\"}",
		Content string `json:"content"`
		// CreateTime is the creation time, e.g. 1676566565604
		CreateTime string `json:"create_time"`
		// Metions is the metions
		Mentions []struct {
			// Key is the mention key, e.g. @_user_1
			Key string `json:"key"`
			// Name is the mention name, e.g. Zero
			Name string `json:"name"`
			//
			ID struct {
				OpenID  string `json:"open_id"`
				UnionID string `json:"union_id"`
				UserID  string `json:"user_id"`
			} `json:"id"`
			//
			TenantKey string `json:"tenant_key"`
		} `json:"mentions"`
	} `json:"message"`
}
