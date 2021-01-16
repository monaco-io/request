package request

import (
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
	URL          string
	Method       string
	Header       map[string]string
	Query        map[string]string
	BodyJSON     interface{}
	BodyXML      interface{}
	BodyString   string
	BasicAuth    BasicAuth
	CustomerAuth string
	Bearer       string
	Timeout      time.Duration
	TLSTimeout   time.Duration
	DialTimeout  time.Duration
	ProxyURL     string
	ProxyServers map[string]string
	Cookies      []*http.Cookie
	CookiesMap   map[string]string
	TLSConfig    *tls.Config
	Transport    *http.Transport
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
