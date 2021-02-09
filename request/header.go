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

// SortedHeader header slice, example [][2]string{{k1,v1},{k2,v2}}
type SortedHeader struct {
	Data [][2]string
}

// Apply apply http headers
func (h SortedHeader) Apply(ctx *context.Context) {
	for i := range h.Data {
		ctx.Request.Header.Set(h.Data[i][0], h.Data[i][1])
	}
}

// Valid user agent in header valid?
func (h SortedHeader) Valid() bool {
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
