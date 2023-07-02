package request

import (
	"github.com/monaco-io/request/xcontext"
)

// Plugin parse user defined request
type Plugin interface {

	// Apply the plugin to http request
	Apply(*xcontext.Context)

	// Valid is the plugin need run?
	Valid() bool
}
