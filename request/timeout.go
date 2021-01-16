package request

import (
	"net"
	"net/http"
	"time"

	"github.com/monaco-io/request/context"
)

var (
	// DialTimeout represents the maximum amount of time the network dialer can take.
	DialTimeout = 30 * time.Second

	// DialKeepAlive represents the maximum amount of time too keep alive the socket.
	DialKeepAlive = 30 * time.Second

	// TLSHandshakeTimeout represents the maximum amount of time that
	// TLS handshake can take defined in the default http.Transport.
	TLSHandshakeTimeout = 10 * time.Second

	// RequestTimeout represents the maximum about of time that
	// a request can take, including dial / request / redirect processes.
	RequestTimeout = 60 * time.Second

	// DefaultDialer defines the default network dialer.
	DefaultDialer = &net.Dialer{
		Timeout:   DialTimeout,
		KeepAlive: DialKeepAlive,
	}
)

// Timeouts represents the supported timeouts
type Timeouts struct {
	// Request represents the total timeout including dial / request / redirect steps
	Request time.Duration

	// TLS represents the maximum amount of time for TLS handshake process
	TLS time.Duration

	// Dial represents the maximum amount of time for dialing process
	Dial time.Duration

	// KeepAlive represents the maximum amount of time to keep alive the socket
	KeepAlive time.Duration
}

// Apply http timeouts of tls, dial, keepalive or all
func (to Timeouts) Apply(ctx *context.Context) {
	if to.Request == 0 {
		to.Request = RequestTimeout
	}
	ctx.Client.Timeout = to.Request

	// Assert http.Transport to work with the instance
	transport, ok := ctx.Client.Transport.(*http.Transport)
	if !ok {
		// If using custom transport, just ignore it
		return
	}

	if to.TLS == 0 {
		to.TLS = TLSHandshakeTimeout
	}
	transport.TLSHandshakeTimeout = to.TLS

	if to.Dial == 0 {
		to.Dial = DialTimeout
	}
	if to.KeepAlive == 0 {
		to.KeepAlive = DialKeepAlive
	}

	transport.Dial = (&net.Dialer{
		Timeout:   to.Dial,
		KeepAlive: to.KeepAlive,
	}).Dial

	// Finally expose the transport to be used
	ctx.Client.Transport = transport
}

// Valid http timeouts of tls, dial, keepalive or all valid?
func (to Timeouts) Valid() bool {
	if to.Request == 0 && to.TLS == 0 &&
		to.Dial == 0 && to.KeepAlive == 0 {
		return false
	}
	return true
}
