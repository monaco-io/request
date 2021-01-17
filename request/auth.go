package request

import (
	"github.com/monaco-io/request/context"
)

// BasicAuth http basic auth with username and password
type BasicAuth struct {
	Username string
	Password string
}

// Apply http basic auth with username and password
func (b BasicAuth) Apply(ctx *context.Context) {
	ctx.Request.SetBasicAuth(b.Username, b.Password)
}

// Valid http basic auth with username and password valid?
func (b BasicAuth) Valid() bool {
	if b.Username == "" {
		return false
	}
	return true
}

// BearerAuth token
type BearerAuth struct {
	Data string
}

// Apply bearer token
func (b BearerAuth) Apply(ctx *context.Context) {
	ctx.Request.Header.Set("Authorization", "Bearer "+b.Data)
}

// Valid bearer token valid?
func (b BearerAuth) Valid() bool {
	if b.Data == "" {
		return false
	}
	return true
}

// CustomerAuth customer Authorization on header
type CustomerAuth struct {
	Data string
}

// Apply customer Authorization on header
func (c CustomerAuth) Apply(ctx *context.Context) {
	ctx.Request.Header.Set("Authorization", c.Data)
}

// Valid customer Authorization on header valid?
func (c CustomerAuth) Valid() bool {
	if c.Data == "" {
		return false
	}
	return true
}
