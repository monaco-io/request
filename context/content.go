package context

import "strings"

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

// SetContentType set content type on header of http request
func (ctx *Context) SetContentType(ct contentType) {
	ctx.Request.Header.Set("Content-Type", ContentTypes[ct])
}

// ContentTypeValid ...
func ContentTypeValid(current string, ct contentType) bool {
	return strings.HasPrefix(current, ContentTypes[ct])
}
