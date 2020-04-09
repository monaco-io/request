package main

import (
	"log"

	"github.com/monaco-io/request"
)

func main() {
	client := request.Client{
		URL:    "https://google.com",
		Method: "POST",
		Params: map[string]string{"hello": "world"},
		Body:   []byte(`{"hello": "world"}`),
	}
	resp, err := client.Do()

	log.Println(resp.Code, string(resp.Data), err)
}
