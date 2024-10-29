package request

import (
	"github.com/monaco-io/request/xcontext"
	"strings"
)

// BasicAuth http basic auth with username and password
type BasicAuth struct {
	Username string
	Password string
}

// Apply http basic auth with username and password
func (b BasicAuth) Apply(ctx *xcontext.Context) {
	ctx.Request.SetBasicAuth(b.Username, b.Password)
}

// Valid http basic auth with username and password valid?
// https://www.ietf.org/rfc/rfc2617.txt (page 5)
//
//	user-pass   = userid ":" password
//	userid      = *<TEXT excluding ":">
//	password    = *TEXT
//
//	The `*TEXT` rule allows an empty string as a value.
//	See https://datatracker.ietf.org/doc/html/rfc5234#section-3.6
func (b BasicAuth) Valid() bool {
	if b.Username != "" {
		return !strings.Contains(b.Username, ":")
	}
	return true
}

// BearerAuth token
type BearerAuth struct {
	Data string
}

// Apply bearer token
func (b BearerAuth) Apply(ctx *xcontext.Context) {
	ctx.Request.Header.Set("Authorization", "Bearer "+b.Data)
}

// Valid bearer token valid?
func (b BearerAuth) Valid() bool {
	return b.Data != ""
}

// CustomerAuth customer Authorization on header
type CustomerAuth struct {
	Data string
}

// Apply customer Authorization on header
func (c CustomerAuth) Apply(ctx *xcontext.Context) {
	ctx.Request.Header.Set("Authorization", c.Data)
}

// Valid customer Authorization on header valid?
func (c CustomerAuth) Valid() bool {
	return c.Data != ""
}
