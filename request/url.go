package request

import (
	"net/url"

	"github.com/monaco-io/request/context"
)

// Host http host name like: https://www.google.com
type Host struct {
	Data string
}

// Apply http host name like: https://www.google.com
func (h Host) Apply(ctx *context.Context) {
	ctx.Request.URL.Host = h.Data
}

// Valid http host name like: https://www.google.com valid?
func (h Host) Valid() bool {
	if h.Data == "" {
		return false
	}
	return true
}

// Path http url path like: /api/v1/xx
type Path struct {
	Data string
}

// Apply http url path like: /api/v1/xx
func (p Path) Apply(ctx *context.Context) {
	if p.Data == "" {
		return
	}
	ctx.Request.URL.Path = p.Data
}

// Valid http url path like: /api/v1/xx valid?
func (p Path) Valid() bool {
	if p.Data == "" {
		return false
	}
	return true
}

// Query http query params like: ?a=1&b=2
type Query struct {
	Data map[string]string
}

// Apply http query params like: ?a=1&b=2
func (q Query) Apply(ctx *context.Context) {
	if q.Data == nil {
		return
	}
	query := ctx.Request.URL.Query()
	for k, v := range q.Data {
		query.Set(k, v)
	}
	ctx.Request.URL.RawQuery = query.Encode()
}

// Valid http url path like: /api/v1/xx valid?
func (q Query) Valid() bool {
	if q.Data == nil {
		return false
	}
	return true
}

// URL http url (host+path+params)
type URL struct {
	Data string
}

// Apply http url (host+path+params)
func (_u URL) Apply(ctx *context.Context) {
	if _u.Data == "" {
		return
	}
	u, err := url.Parse(_u.Data)
	if err != nil {
		return
	}
	ctx.Request.URL = u
}

// Valid http url path like: /api/v1/xx valid?
func (_u URL) Valid() bool {
	if _u.Data == "" {
		return false
	}
	return true
}
