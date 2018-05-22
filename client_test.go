package gosdk

import (
	"testing"
)

func TestReqeust(t *testing.T) {
	api := NewClient("YOUR_ACCESS_KEY", "YOUR_SECRET_KEY")
	data := map[string]string{
		"number":  "13800138000",
		"country": "CN",
	}
	url := "https://api.crawler.club/phone"
	response, _ := api.Request(url, data)
	t.Log(response)
}
