package request

import (
	"github.com/monaco-io/request/context"
)

// Plugin parse user defined request
type Plugin interface {

	// Apply the plugin to http request
	Apply(*context.Context)

	// Valid is the plugin need run?
	Valid() bool
}
