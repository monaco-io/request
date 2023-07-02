package request

import "github.com/monaco-io/request/xcontext"

// Method http method: GET, POST, DELETE ...
type Method struct {
	Data string
}

// Apply http method: GET, POST, DELETE ...
func (m Method) Apply(ctx *xcontext.Context) {
	ctx.Request.Method = m.Data
}

// Valid  http method: GET, POST, DELETE ... valid?
func (m Method) Valid() bool {
	return m.Data != ""
}
