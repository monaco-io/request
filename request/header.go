package request

import (
	"github.com/monaco-io/request/context"
)

type contentType int

const (
	// HTML text/html
	HTML contentType = iota

	// JSON application/json
	JSON

	// XML application/xml
	XML

	// Text text/plain
	Text

	// URLEncodedForm application/x-www-form-urlencoded
	URLEncodedForm
)

// ContentTypes http content type map
var ContentTypes = map[contentType]string{
	HTML:           "text/html",
	JSON:           "application/json",
	XML:            "application/xml",
	Text:           "text/plain",
	URLEncodedForm: "application/x-www-form-urlencoded",
}

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

// setContentType set content type on header of http request
func setContentType(ctx *context.Context, ct contentType) {
	ctx.Request.Header.Set("Content-Type", ContentTypes[ct])
}
