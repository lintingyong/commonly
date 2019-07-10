package http

import (
	"testing"
)

func TestHttpGet(t *testing.T) {
	bytes, err := HttpGet("https://www.baidu.com/")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log("\n", string(bytes))
}

func TestHttpPost(t *testing.T) {
	bytes, err := HttpPost("https://www.baidu.com", map[string]string{
		"ie": "UTF-8",
		"wd": "http",
	})
	if err != nil {
		t.Log(err)
		return
	}
	t.Log("\n", string(bytes))
}

func TestHttpClient_Get(t *testing.T) {
	var client HttpClient
	bytes, err := client.Get("https://www.baidu.com/", []byte{})
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(bytes))
}

func TestHttpClient_Post(t *testing.T) {
	var client HttpClient
	bytes, err := client.Post("https://www.baidu.com/", []byte{})
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(bytes))
}
