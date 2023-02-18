package content

import "testing"

func TestBuilder(t *testing.T) {
	// text
	msgType, content, err := NewContent().Text(&ContentTypeText{
		Text: "testcontent",
	}).Build()
	if err != nil {
		t.Fatal(err)
	}
	expectedMsgType := "text"
	expectedContent := `{"text":"testcontent"}`
	if msgType != expectedMsgType || content != expectedContent {
		t.Errorf("msgType expected %s, got %s", expectedMsgType, msgType)
		t.Errorf("content expected %s, got %s", expectedContent, content)
	}

	// post
	msgType, content, err = NewContent().Post(&ContentTypePost{
		ZhCN: &ContentTypePostBody{
			Content: [][]ContentTypePostBodyItem{
				{
					{
						Tag:  "text",
						Text: "testcontent",
					},
					{
						Tag:  "a",
						Href: "https://zero.com",
						Text: "zero.com",
					},
				},
			},
		},
	}).Build()
	if err != nil {
		t.Fatal(err)
	}
	expectedMsgType = "post"
	expectedContent = `{"zh_cn":{"content":[[{"tag":"text","text":"testcontent"},{"tag":"a","text":"zero.com","href":"https://zero.com"}]]}}`
	if msgType != expectedMsgType || content != expectedContent {
		t.Errorf("msgType expected %s, got %s", expectedMsgType, msgType)
		t.Errorf("content expected %s, got %s", expectedContent, content)
	}

	// image
	msgType, content, err = NewContent().Image(&ContentTypeImage{
		ImageKey: "img_7ea74629-9191-4176-998c-2e603c9c5e8g",
	}).Build()
	if err != nil {
		t.Fatal(err)
	}
	expectedMsgType = "image"
	expectedContent = `{"image_key":"img_7ea74629-9191-4176-998c-2e603c9c5e8g"}`
	if msgType != expectedMsgType || content != expectedContent {
		t.Errorf("msgType expected %s, got %s", expectedMsgType, msgType)
		t.Errorf("content expected %s, got %s", expectedContent, content)
	}
}
