package content

import (
	"encoding/json"
	"fmt"
)

func InferContentMsgType(content string) (msgType string, err error) {
	jd := map[string]any{}
	if err := json.Unmarshal([]byte(content), &jd); err != nil {
		return "", fmt.Errorf("content must be a valid json, but got %s (err: %v)", content, err)
	}

	// text
	if _, ok := jd["text"]; ok {
		return "text", nil
	}

	// post
	if _, ok := jd["zh_cn"]; ok {
		return "post", nil
	}
	if _, ok := jd["en_us"]; ok {
		return "post", nil
	}

	// image
	if _, ok := jd["image_key"]; ok {
		return "image", nil
	}

	// interactive
	if _, ok := jd["elements"]; ok {
		return "interactive", nil
	}

	// share_chat
	if _, ok := jd["chat_id"]; ok {
		return "share_chat", nil
	}

	// share_user
	if _, ok := jd["user_id"]; ok {
		return "share_user", nil
	}

	// audio | media | file | sticker
	if _, ok := jd["FileKey"]; ok {
		if _, ok := jd["image_key"]; ok {
			return "media", nil
		}

		// cannot infer audio/file/sticker
	}

	return "", fmt.Errorf("cannot infer content msg_type(maybe audio/file/sticker), maybe you need custom set msgType at content(%s)", content)
}
