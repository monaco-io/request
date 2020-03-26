package request

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

// Do send http request
func (c *Client) Do() ([]byte, error) {
	client := &http.Client{
		Timeout: c.Timeout * time.Second,
	}

	// encode like https://google.com?hello=world&package=request
	var encodeError error
	if c.URL, encodeError = EncodeURL(c.URL, c.Params); encodeError != nil {
		return []byte{}, encodeError
	}

	// build request
	_request, buildRequestError := http.NewRequest(c.Method, c.URL, bytes.NewReader(c.Body))
	if buildRequestError != nil {
		return []byte{}, buildRequestError
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
	resp, respError := client.Do(_request)
	if respError != nil {
		return []byte{}, respError
	}
	defer func() { _ = resp.Body.Close() }()

	// read response data form resp
	data, readBodyErr := ioutil.ReadAll(resp.Body)
	if readBodyErr != nil {
		return []byte{}, readBodyErr
	}

	return data, nil
}
