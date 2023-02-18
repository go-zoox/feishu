package content

import "testing"

func TestInferContentMsgType(t *testing.T) {
	content := `{"text":"testcontent"}`
	expected := "text"
	msgType, err := InferContentMsgType(content)
	if err != nil {
		t.Fatal(err)
	}
	if msgType != expected {
		t.Fatalf("expected %v, got %v", expected, msgType)
	}

	content = `{"zh_cn":{"content":[[{"tag":"text","text":"testcontent"},{"tag":"a","text":"zero.com","href":"https://zero.com"}]]}}`
	expected = "post"
	msgType, err = InferContentMsgType(content)
	if err != nil {
		t.Fatal(err)
	}
	if msgType != expected {
		t.Fatalf("expected %v, got %v", expected, msgType)
	}

	content = `{"image_key":"img_7ea74629-9191-4176-998c-2e603c9c5e8g"}`
	expected = "image"
	msgType, err = InferContentMsgType(content)
	if err != nil {
		t.Fatal(err)
	}
	if msgType != expected {
		t.Fatalf("expected %v, got %v", expected, msgType)
	}
}
