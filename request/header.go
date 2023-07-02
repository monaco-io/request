package request

import "github.com/monaco-io/request/xcontext"

// Header http header
type Header struct {
	Data map[string]string
}

// Apply apply http headers
func (h Header) Apply(ctx *xcontext.Context) {
	for k, v := range h.Data {
		ctx.Request.Header.Set(k, v)
	}
}

// Valid user agent in header valid?
func (h Header) Valid() bool {
	return h.Data != nil
}

// SortedHeader header slice, example [][2]string{{k1,v1},{k2,v2}}
type SortedHeader struct {
	Data [][2]string
}

// Apply apply http headers
func (h SortedHeader) Apply(ctx *xcontext.Context) {
	for i := range h.Data {
		ctx.Request.Header.Set(h.Data[i][0], h.Data[i][1])
	}
}

// Valid user agent in header valid?
func (h SortedHeader) Valid() bool {
	return h.Data != nil
}

// UserAgent user agent in header
type UserAgent struct {
	Version string
}

// Apply user agent in header
func (ua UserAgent) Apply(ctx *xcontext.Context) {
	ctx.Request.Header.Set("User-Agent", "github.com/monaco-io/request/"+ua.Version)
}

// Valid user agent in header valid?
func (ua UserAgent) Valid() bool {
	return ua.Version != ""
}
