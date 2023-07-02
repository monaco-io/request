package xcontext

import (
	originContext "context"
	"net/http"
)

// Context  HTTP transactions
type Context struct {
	Client   *http.Client   `json:"client,omitempty"`
	Request  *http.Request  `json:"request,omitempty"`
	Response *http.Response `json:"response,omitempty"`
	err      error
}

// New creates an empty Context
func New() *Context {
	return &Context{
		Request: newRequest(),
		Client:  &http.Client{Transport: http.DefaultTransport},
	}
}

// NewWithContext creates an empty Context
func NewWithContext(ctx originContext.Context) *Context {
	return &Context{
		Request: newRequestWithContext(ctx),
		Client:  &http.Client{Transport: http.DefaultTransport},
	}
}

// GetClient get original http client
func (c *Context) GetClient() *http.Client {
	return c.Client
}

// GetRequest get original http request
func (c *Context) GetRequest() *http.Request {
	return c.Request
}

// GetResponse get original http response
func (c *Context) GetResponse() *http.Response {
	return c.Response
}

func newRequest() *http.Request {
	r, _ := http.NewRequest("", "", nil)
	return r
}

func newRequestWithContext(ctx originContext.Context) *http.Request {
	r, _ := http.NewRequestWithContext(ctx, "", "", nil)
	return r
}
