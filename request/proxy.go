package request

import (
	"net/http"
	"net/url"

	"github.com/monaco-io/request/context"
)

// Proxy http proxy url or multi services
type Proxy struct {
	Servers map[string]string
	URL     string
}

// Apply http proxy url or multi services
func (p Proxy) Apply(ctx *context.Context) {
	// Assert http.Transport to work with the instance
	transport, ok := ctx.Client.Transport.(*http.Transport)
	if !ok {
		// If using a custom transport, just ignore it
		return
	}

	if p.URL != "" {
		proxy, _ := url.Parse(p.URL)
		transport.Proxy = http.ProxyURL(proxy)
	}

	if p.Servers != nil {
		// Define the proxy function to be used during the transport
		transport.Proxy = func(req *http.Request) (*url.URL, error) {
			if value, ok := p.Servers[req.URL.Scheme]; ok {
				return url.Parse(value)
			}
			return http.ProxyFromEnvironment(req)
		}
	}

	// Override the transport
	ctx.Client.Transport = transport
}

// Valid http proxy url or multi services valid?
func (p Proxy) Valid() bool {
	if p.Servers == nil && p.URL == "" {
		return false
	}
	return true
}
