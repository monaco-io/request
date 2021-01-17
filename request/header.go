package request

import (
	"github.com/monaco-io/request/context"
)

// Header http header
type Header struct {
	Data map[string]string
}

// Apply apply http headers
func (h Header) Apply(ctx *context.Context) {
	for k, v := range h.Data {
		ctx.Request.Header.Set(k, v)
	}
}

// Valid user agent in header valid?
func (h Header) Valid() bool {
	if h.Data == nil {
		return false
	}
	return true
}

// UserAgent user agent in header
type UserAgent struct {
	Version string
}

// Apply user agent in header
func (ua UserAgent) Apply(ctx *context.Context) {
	ctx.Request.Header.Set("User-Agent", "github.com/monaco-io/request/"+ua.Version)
}

// Valid user agent in header valid?
func (ua UserAgent) Valid() bool {
	if ua.Version == "" {
		return false
	}
	return true
}
