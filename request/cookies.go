package request

import (
	"net/http"

	"github.com/monaco-io/request/context"
)

// Cookies http cookies
type Cookies struct {
	Data []*http.Cookie
	Map  map[string]string
}

// Apply http cookies
func (c Cookies) Apply(ctx *context.Context) {
	for _, cookie := range c.Data {
		ctx.Request.AddCookie(cookie)
	}

	for k, v := range c.Map {
		cookie := &http.Cookie{Name: k, Value: v}
		ctx.Request.AddCookie(cookie)
	}
}

// Valid http cookies valid?
func (c Cookies) Valid() bool {
	if c.Data == nil && c.Map == nil {
		return false
	}
	return true
}
