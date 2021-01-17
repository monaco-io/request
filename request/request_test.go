package request

import (
	"testing"
)

func TestRequest_Send(t *testing.T) {
	var data map[string]interface{}
	resp := New().
		POST("http://httpbin.org/post").
		AddHeader(map[string]string{"Google": "google"}).
		AddBasicAuth("google", "google").
		AddURLEncodedForm(map[string]string{"data": "google"}).
		Send().
		Scan(&data)

	if !resp.OK() {
		t.Error(resp.Error())
	}

	if data["headers"].(map[string]interface{})["Authorization"].(string) != "Basic Z29vZ2xlOmdvb2dsZQ==" {
		t.Error("Authorization")
	}

	if data["form"].(map[string]interface{})["data"] != "google" {
		t.Error("form")
	}
}
