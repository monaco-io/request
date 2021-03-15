package request

import (
	"testing"
)

func TestRequest_URLEncodedForm(t *testing.T) {
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

func TestRequest_Form(t *testing.T) {
	resp := New().
		POST("http://httpbin.org/post").
		AddHeader(map[string]string{"Google": "google"}).
		AddBasicAuth("google", "google").
		AddMultipartForm(map[string]string{"field": "value"}, []string{"no_exist.txt"}).
		Send()

	if resp.Error().Error() != "read local file failed: open no_exist.txt: no such file or directory" {
		t.Error(resp.Error())
	}
}
