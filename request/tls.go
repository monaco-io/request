package request

import (
	"crypto/tls"
	"net/http"

	"github.com/monaco-io/request/context"
)

// TLSConfig http tls config of transport
type TLSConfig struct {
	*tls.Config
}

// Apply http tls config of transport
func (tc TLSConfig) Apply(ctx *context.Context) {

	// Assert http.Transport to work with the instance
	transport, ok := ctx.Client.Transport.(*http.Transport)
	if !ok {
		// If using a custom transport, just ignore it
		// TODO:
		return
	}

	// Override the http.Client transport
	transport.TLSClientConfig = tc.Config
	ctx.Client.Transport = transport
}

// Valid http timeouts of tls, dial, keepalive or all valid?
func (tc TLSConfig) Valid() bool {
	return true
}
