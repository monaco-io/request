package request

import (
	"github.com/monaco-io/request/xcontext"
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
func (b BasicAuth) Valid() bool {
	return b.Username != ""
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
