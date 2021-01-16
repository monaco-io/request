package request

import (
	"net/http"
	"reflect"

	"github.com/monaco-io/request/context"
)

// Transport http Transport
type Transport struct {
	http.RoundTripper
}

// Apply http Transport
func (t Transport) Apply(ctx *context.Context) {

	// Override the http.Client transport
	ctx.Client.Transport = t.RoundTripper
}

// Valid http timeouts of tls, dial, keepalive or all valid?
func (t Transport) Valid() bool {
	return !reflect.ValueOf(t.RoundTripper).IsNil()
}
