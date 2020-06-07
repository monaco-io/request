package request

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

// Do send http request
func (c *Client) Do() (resp SugaredResp, err error) {
	defer resp.Close()

	if err = c.buildRequest(); err != nil {
		return
	}

	// send request and close on func call end
	if resp.resp, err = c.client.Do(c.req); err != nil {
		return
	}

	// read response data form resp
	resp.Data, err = ioutil.ReadAll(resp.resp.Body)
	resp.Code = resp.resp.StatusCode
	return
}

func (c *Client) buildRequest() (err error) {
	if err = c.applyRequest(); err != nil {
		return
	}
	c.applyHTTPMethod()
	c.applyBasicAuth()
	c.applyClient()
	c.applyTimeout()
	c.applyCookies()
	c.applyProxy()
	return
}

func (c *Client) applyRequest() (err error) {
	// encode requestURL.httpURL like https://google.com?hello=world&package=request
	c.requestURL = requestURL{
		urlString:  c.URL,
		parameters: c.Params,
	}
	if err = c.requestURL.EncodeURL(); err != nil {
		return
	}
	c.req, err = http.NewRequest(c.Method, c.requestURL.string(), bytes.NewReader(c.Body))
	return
}

func (c *Client) applyHTTPMethod() {
	if c.Method == "POST" {
		if c.ContentType == "" {
			c.ContentType = ApplicationJSON
		}
		c.req.Header.Set("Content-Type", string(c.ContentType))
	}
	for k, v := range c.Header {
		c.req.Header.Add(k, v)
	}
}

func (c *Client) applyBasicAuth() {
	if c.BasicAuth.Username != "" && c.BasicAuth.Password != "" {
		c.req.SetBasicAuth(c.BasicAuth.Username, c.BasicAuth.Password)
	}
}

func (c *Client) applyClient() {
	c.client = &http.Client{}
}

func (c *Client) applyTimeout() {
	if c.Timeout > 0 {
		c.client.Timeout = c.Timeout * time.Second
	}
}

func (c *Client) applyCookies() {
	if c.Cookies != nil {
		jar, _ := cookiejar.New(nil)
		jar.SetCookies(&url.URL{Scheme: c.requestURL.scheme(), Host: c.requestURL.host()}, c.Cookies)
		c.client.Jar = jar
	}
}

// TODO: raise proxy error
func (c *Client) applyProxy() {
	if c.ProxyURL != "" {
		if proxy, err := url.Parse(c.ProxyURL); err == nil && proxy != nil {
			c.client.Transport = &http.Transport{
				Proxy:           http.ProxyURL(proxy),
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
		}
	}
}

// Resp do request and get original http response struct
func (c *Client) Resp() (resp *http.Response, err error) {
	if err = c.buildRequest(); err != nil {
		return
	}
	return c.client.Do(c.req)
}
