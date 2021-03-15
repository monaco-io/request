package curl

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestGetCommand(t *testing.T) {
	form := url.Values{}
	form.Add("key1", "val1")
	form.Add("key2", "val2")
	body := form.Encode()

	req, _ := http.NewRequest(http.MethodPost, "https://www.google.com", ioutil.NopCloser(bytes.NewBufferString(body)))
	req.Header.Set("HEADER1", "HEADER_VAL1")

	command, _ := GetCommand(req)

	if command != "curl -X 'POST' -d 'key1=val1&key2=val2' -H 'Header1: HEADER_VAL1' 'https://www.google.com'" {
		t.Fatal(command)
	}
}

func TestGetCommand_json(t *testing.T) {
	req, _ := http.NewRequest(http.MethodPut, "https://www.google.com?a=1&b=2", bytes.NewBufferString(`{"hello":"world"}`))
	req.Header.Set("Content-Type", "application/json")

	command, _ := GetCommand(req)

	if command != `curl -X 'PUT' -d '{"hello":"world"}' -H 'Content-Type: application/json' 'https://www.google.com?a=1&b=2'` {
		t.Fatal(command)
	}
}
