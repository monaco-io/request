package xcontext

// TraceDo ...
func (c *Context) TraceDo() (err error) {
	c.Request = c.GetRequest().WithContext(c.GetRequest().Context())
	c.Response, err = c.Client.Do(c.GetRequest())
	if err != nil {
		c.SetError(err)
	}
	return
}
