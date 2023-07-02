package request

import (
	"time"

	"github.com/monaco-io/request/xcontext"
)

var (
	// defaultTLSHandshakeTimeout represents the maximum amount of time that
	// TLS handshake can take defined in the default http.Transport.
	// defaultTLSHandshakeTimeout = 10 * time.Second

	// defaultRequestTimeout represents the maximum about of time that
	// a request can take, including dial / request / redirect processes.
	defaultRequestTimeout = 60 * time.Second
)

// Timeout represents the supported timeouts
type Timeout struct {
	// Data represents the total timeout including dial / request / redirect steps
	Data time.Duration
}

// Apply http timeouts of tls, dial, keepalive or all
func (to Timeout) Apply(ctx *xcontext.Context) {
	if to.Data == 0 {
		to.Data = defaultRequestTimeout
	}
	ctx.Client.Timeout = to.Data
}

// Valid http timeouts of tls, dial, keepalive or all valid?
func (to Timeout) Valid() bool {
	return to.Data != 0
}
