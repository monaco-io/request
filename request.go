package request

import (
	originContext "context"

	"github.com/monaco-io/request/request"
	"github.com/monaco-io/request/response"
	"github.com/monaco-io/request/xcontext"
)

// Send http request
func (c *Client) Send() *response.Sugar {

	ctx := c.initContext()
	resp := response.New(ctx)

	return resp.Do()
}

func (c *Client) initContext() *xcontext.Context {
	var ctx *xcontext.Context

	if c.Context != nil {
		ctx = xcontext.NewWithContext(c.Context)
	} else {
		ctx = xcontext.New()
	}

	plugins := []request.Plugin{
		request.URL{Data: c.URL},
		request.UserAgent{Version: Version},
		request.Query{Data: c.Query},
		request.Method{Data: c.Method},
		request.Header{Data: c.Header},
		request.SortedHeader{Data: c.SortedHeader},
		request.Cookies{Data: c.Cookies, Map: c.CookiesMap},
		request.BasicAuth{Username: c.BasicAuth.Username, Password: c.BasicAuth.Password},
		request.BearerAuth{Data: c.Bearer},
		request.CustomerAuth{Data: c.CustomerAuth},
		request.Timeout{Data: c.Timeout},
		request.Proxy{Servers: c.ProxyServers, URL: c.ProxyURL},
		request.BodyJSON{Data: c.JSON},
		request.BodyString{Data: c.String},
		request.BodyXML{Data: c.XML},
		request.BodyYAML{Data: c.YAML},
		request.BodyForm{Fields: c.MultipartForm.Fields, Files: c.MultipartForm.Files},
		request.BodyURLEncodedForm{Data: c.URLEncodedForm},
		request.Transport{Transport: c.Transport},
		request.TLSConfig{Config: c.TLSConfig},
	}

	for _, plugin := range plugins {
		if plugin.Valid() {
			plugin.Apply(ctx)
		}
	}

	return ctx
}

// New a empty request
func New() *request.Request {
	return request.New()
}

// NewWithContext a empty request
func NewWithContext(ctx originContext.Context) *request.Request {
	return request.NewWithContext(ctx)
}
