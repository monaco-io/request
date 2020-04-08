package request

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

// Do send http request
func (c *Client) Do() (resp []byte, err error) {
	client := &http.Client{
		Timeout: c.Timeout * time.Second,
	}
	if err := c.buildRequest(); err != nil {
		return
	}
	// send request and close on func call end
	response, err := client.Do(c.req)
	if err != nil {
		return []byte{}, err
	}
	defer func() { _ = response.Body.Close() }()

	// read response data form resp
	return ioutil.ReadAll(response.Body)
}

func (c *Client) buildRequest() (err error) {

	// encode like https://google.com?hello=world&package=request
	if c.URL, err = EncodeURL(c.URL, c.Params); err != nil {
		return err
	}

	// build request
	c.req, err = http.NewRequest(c.Method, c.URL, bytes.NewReader(c.Body))
	if err != nil {
		return err
	}

	// add Header to request
	if c.Method == "POST" {
		if c.ContentType == "" {
			c.ContentType = ApplicationJSON
		}
		c.req.Header.Set("Content-Type", string(c.ContentType))
	}
	for k, v := range c.Header {
		c.req.Header.Add(k, v)
	}

	// set basic Auth of request
	if c.BasicAuth.Username != "" && c.BasicAuth.Password != "" {
		c.req.SetBasicAuth(c.BasicAuth.Username, c.BasicAuth.Password)
	}
	return nil
}
