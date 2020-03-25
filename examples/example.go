package main

import (
	"log"

	"github.com/Monaco-io/request"
)

func main() {
	client := request.Client{
		URL:    "https://baidu.com",
		Method: "POST",
		Params: map[string]string{"hello": "world"},
		Body:   []byte(`{"hello": "world"}`),
	}
	resp, err := client.Do()

	log.Println(string(resp), err)
}
