package context

import (
	"net/http/httptrace"
	"time"
)

var startTime = time.Now()

func now() time.Duration { return time.Since(startTime) }

// Time ...
type Time struct {
	Offset        time.Duration `json:"offset"`
	Duration      time.Duration `json:"duration"`
	ConnDuration  time.Duration `json:"connDuration"`  // connection setup(DNS lookup + Dial up) duration
	DNSDuration   time.Duration `json:"dnsDuration"`   // dns lookup duration
	ReqDuration   time.Duration `json:"reqDuration"`   // request "write" duration
	ResDuration   time.Duration `json:"resDuration"`   // response "read" duration
	DelayDuration time.Duration `json:"delayDuration"` // delay between response and request
}

// TraceDo ...
func (c *Context) TraceDo() (err error) {
	c.TimeTrace.Offset = now()
	var dnsStart, connStart, resStart, reqStart, delayStart time.Duration

	trace := &httptrace.ClientTrace{
		DNSStart: func(info httptrace.DNSStartInfo) {
			dnsStart = now()
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			c.TimeTrace.DNSDuration = now() - dnsStart
		},
		GetConn: func(h string) {
			connStart = now()
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			if !connInfo.Reused {
				c.TimeTrace.ConnDuration = now() - connStart
			}
			reqStart = now()
		},
		WroteRequest: func(w httptrace.WroteRequestInfo) {
			c.TimeTrace.ReqDuration = now() - reqStart
			delayStart = now()
		},
		GotFirstResponseByte: func() {
			c.TimeTrace.DelayDuration = now() - delayStart
			resStart = now()
		},
	}
	c.Request = c.GetRequest().WithContext(httptrace.WithClientTrace(c.GetRequest().Context(), trace))
	c.Response, err = c.Client.Do(c.GetRequest())
	if err != nil {
		c.SetError(err)
	}
	end := now()
	c.TimeTrace.ResDuration = end - resStart
	c.TimeTrace.Duration = end - c.TimeTrace.Offset
	return
}
