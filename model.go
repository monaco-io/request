package request

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"
)

// Client Method
/*
     Method         = "OPTIONS"                ; Section 9.2
                    | "GET"                    ; Section 9.3
                    | "HEAD"                   ; Section 9.4
                    | "POST"                   ; Section 9.5
                    | "PUT"                    ; Section 9.6
                    | "DELETE"                 ; Section 9.7
                    | "TRACE"                  ; Section 9.8
                    | "CONNECT"                ; Section 9.9
                    | extension-method
   extension-method = token
     token          = 1*<any CHAR except CTLs or separators>
*/
type Client struct {
	// Context go context
	Context context.Context

	// URL http request url like: https://www.google.com
	URL string

	// Method http method GET/POST/POST/DELETE ...
	Method string

	// Header http header
	Header map[string]string

	// SortedHeader http sorted header, example: [][2]string{{"h1", "v1"}, {"h2", "v2"}}
	SortedHeader [][2]string

	// Query params on http url
	Query map[string]string

	// JSON body as json string/bytes/struct
	JSON interface{}

	// XML body as xml string/bytes/struct
	XML interface{}

	// YAML body as yaml string/bytes/struct
	YAML interface{}

	// XML body as string
	String string

	// URLEncodedForm string/bytes/map[string][]string
	URLEncodedForm interface{}

	// MultipartForm key value pairs
	MultipartForm MultipartForm

	// BasicAuth http basic auth with username and password
	BasicAuth BasicAuth

	// CustomerAuth add Authorization xxx to header
	CustomerAuth string

	// CustomerAuth add Authorization bearer xxx to header
	Bearer string

	// Timeout http request timeout
	Timeout time.Duration

	// TLSTimeout tls timeout
	TLSTimeout time.Duration

	// DialTimeout dial timeout
	DialTimeout time.Duration

	// ProxyURL proxy url
	ProxyURL string

	// Define the proxy function to be used during the transport
	ProxyServers map[string]string

	// Cookies original http cookies
	Cookies []*http.Cookie

	// CookiesMap add cookies as map
	CookiesMap map[string]string

	// TLSConfig tls config on transport
	TLSConfig *tls.Config

	// Transport http transport
	Transport *http.Transport
}

// BasicAuth Add Username:Password as Basic Auth
type BasicAuth struct {
	Username string
	Password string
}

const (
	// OPTIONS http options
	OPTIONS = "OPTIONS"

	// GET http get
	GET = "GET"

	// HEAD http head
	HEAD = "HEAD"

	// POST http post
	POST = "POST"

	// PUT http put
	PUT = "PUT"

	// DELETE http delete
	DELETE = "DELETE"

	// TRACE http trace
	TRACE = "TRACE"

	// CONNECT http connect
	CONNECT = "CONNECT"

	// PATCH http patch
	PATCH = "PATCH"
)

// A alias of []interface{}
type A []interface{}

// H alias of map[string]interface{}
type H map[string]interface{}

// MultipartForm Fields is key value pairs, Files is a list of local files
type MultipartForm struct {
	Fields map[string]string
	Files  []string
}
