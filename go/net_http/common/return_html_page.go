package common

import (
	"encoding/json"
	"fmt"
)

func ReturnHtmlPage01(text string) (string, error) {
	return "<html><script>alert(1)</script></html>", nil
}

func ReturnHtmlPage02(text string) (string, error) {
	text = "<script>"
	tmpStruct := struct {
		Title string `json:"title"`
	}{text}
	// safe for xss: htmlEscape
	str, err := json.Marshal(&tmpStruct)
	if err != nil {
		return "", err
	}
	html := fmt.Sprintf("<html><body><p>%s</p></body></html>", str)
	return html, nil
}

func ReturnHtmlPage03(text string) (string, error) {
	return "<html><script>alert(1)</script></html>", nil
}

func ReturnHtmlPage04(text string) (string, error) {
	return "<html><script>alert(1)</script></html>", nil
}

func ReturnHtmlPage05(text string) (string, error) {
	return "<html><script>alert(1)</script></html>", nil
}
