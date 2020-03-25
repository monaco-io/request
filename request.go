package request

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func (c *Client) Do() ([]byte, error) {
	client := &http.Client{}

	// encode to https://google.com?hello=world&package=request
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
	for k, v := range c.Header {
		_request.Header.Add(k, v)
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

func (c *Client) GET() {

}

func (c *Client) POST() {

}

func (c *Client) DELETE() {

}

func (c *Client) PUT() {

}

func (c *Client) PATCH() {

}

func (c *Client) HEAD() {

}
