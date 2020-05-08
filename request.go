package request

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Do send http request
func (c *Client) Do() (resp SugaredResp, err error) {
	defer resp.Close()

	if err := c.buildRequest(); err != nil {
		return resp, err
	}

	// send request and close on func call end
	if resp.resp, err = c.client.Do(c.req); err != nil {
		return resp, err
	}

	// read response data form resp
	resp.Data, err = ioutil.ReadAll(resp.resp.Body)
	resp.Code = resp.resp.StatusCode
	return resp, err
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

	c.client = &http.Client{}

	if c.Timeout > 0 {
		c.client.Timeout = c.Timeout * time.Second
	}

	if c.ProxyURL != "" {
		if proxy, err := url.Parse(c.ProxyURL); err == nil && proxy != nil {
			c.client.Transport = &http.Transport{
				Proxy:           http.ProxyURL(proxy),
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
		}
	}

	return err
}

// Resp do request and get original http response struct
func (c *Client) Resp() (resp *http.Response, err error) {
	if err = c.buildRequest(); err != nil {
		return resp, err
	}
	return c.client.Do(c.req)
}
