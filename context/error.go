package context

// Error get context error
func (c *Context) Error() error {
	return c.err
}

// ErrorString get context error as string
func (c *Context) ErrorString() string {
	if c.Error() == nil {
		return ""
	}
	return c.err.Error()
}

// HasError get context error as string
func (c *Context) HasError() bool {
	return c.Error() != nil
}

// SetError set context error
func (c *Context) SetError(err error) {
	if c.HasError() {
		return
	}
	c.err = err
}

// ResetError set context error to nil
func (c *Context) ResetError() {
	c.err = nil
}
