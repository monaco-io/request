package request

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

// Do send http request
func (c *Client) Do() ([]byte, error) {
	var err error
	client := &http.Client{
		Timeout: c.Timeout * time.Second,
	}

	// encode like https://google.com?hello=world&package=request
	if c.URL, err = EncodeURL(c.URL, c.Params); err != nil {
		return []byte{}, err
	}

	// build request
	_request, err := http.NewRequest(c.Method, c.URL, bytes.NewReader(c.Body))
	if err != nil {
		return []byte{}, err
	}

	// add Header to request
	if c.Method == "POST" {
		if c.ContentType == "" {
			c.ContentType = ApplicationJSON
		}
		_request.Header.Set("Content-Type", string(c.ContentType))
	}
	for k, v := range c.Header {
		_request.Header.Add(k, v)
	}

	// set basic Auth of request
	if c.BasicAuth.Username != "" && c.BasicAuth.Password != "" {
		_request.SetBasicAuth(c.BasicAuth.Username, c.BasicAuth.Password)
	}

	// send request and close on func call end
	resp, err := client.Do(_request)
	if err != nil {
		return []byte{}, err
	}
	defer func() { _ = resp.Body.Close() }()

	// read response data form resp
	return ioutil.ReadAll(resp.Body)
}
