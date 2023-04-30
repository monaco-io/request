package request

import (
	"net/http"

	"github.com/monaco-io/request/xcontext"
)

// Transport http Transport
type Transport struct {
	*http.Transport
}

// Apply http Transport
func (t Transport) Apply(ctx *xcontext.Context) {
	// Override the http.Client transport
	ctx.Client.Transport = t.Transport
}

// Valid http timeouts of tls, dial, keepalive or all valid?
func (t Transport) Valid() bool {
	return t.Transport != nil
}
