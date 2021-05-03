package request

import (
	originContext "context"
	"crypto/tls"
	"net/http"

	"github.com/monaco-io/request/context"
	"github.com/monaco-io/request/response"
)

// Request an alias of context
type Request struct {
	ctx *context.Context
}

func (r *Request) use(p Plugin) *Request {
	p.Apply(r.Ctx())
	return r
}

// New request struct
func New() *Request {
	return &Request{ctx: context.New()}
}

// NewWithContext request struct
func NewWithContext(ctx originContext.Context) *Request {
	return &Request{ctx: context.NewWithContext(ctx)}
}

// Error get request error
func (r *Request) Error() error {
	return r.ctx.Error()
}

// Ctx get request ctx
func (r *Request) Ctx() *context.Context {
	return r.ctx
}

// POST use POST method and http url
func (r *Request) POST(url string) *Request {
	return r.
		use(Method{Data: "POST"}).
		use(URL{Data: url})
}

// GET use GET method and http url
func (r *Request) GET(url string) *Request {
	return r.
		use(Method{Data: "GET"}).
		use(URL{Data: url})
}

// PUT use PUT method and http url
func (r *Request) PUT(url string) *Request {
	return r.
		use(Method{Data: "PUT"}).
		use(URL{Data: url})
}

// DELETE use DELETE method and http url
func (r *Request) DELETE(url string) *Request {
	return r.
		use(Method{Data: "DELETE"}).
		use(URL{Data: url})
}

// OPTIONS use OPTIONS method and http url
func (r *Request) OPTIONS(url string) *Request {
	return r.
		use(Method{Data: "OPTIONS"}).
		use(URL{Data: url})
}

// HEAD use HEAD method and http url
func (r *Request) HEAD(url string) *Request {
	return r.
		use(Method{Data: "HEAD"}).
		use(URL{Data: url})
}

// TRACE use TRACE method and http url
func (r *Request) TRACE(url string) *Request {
	return r.
		use(Method{Data: "TRACE"}).
		use(URL{Data: url})
}

// PATCH use PATCH method and http url
func (r *Request) PATCH(url string) *Request {
	return r.
		use(Method{Data: "PATCH"}).
		use(URL{Data: url})
}

// AddQuery ...
func (r *Request) AddQuery(data map[string]string) *Request {
	return r.
		use(Query{Data: data})
}

// AddHeader ...
func (r *Request) AddHeader(data map[string]string) *Request {
	return r.
		use(Header{Data: data})
}

// AddSortedHeader ...
func (r *Request) AddSortedHeader(data [][2]string) *Request {
	return r.
		use(SortedHeader{Data: data})
}

// AddBasicAuth ...
func (r *Request) AddBasicAuth(username, password string) *Request {
	return r.
		use(BasicAuth{Username: username, Password: password})
}

// AddBearerAuth ...
func (r *Request) AddBearerAuth(data string) *Request {
	return r.
		use(BearerAuth{Data: data})
}

// AddJSON ...
func (r *Request) AddJSON(data interface{}) *Request {
	return r.
		use(BodyJSON{Data: data})
}

// AddXML ...
func (r *Request) AddXML(data interface{}) *Request {
	return r.
		use(BodyXML{Data: data})
}

// AddYAML ...
func (r *Request) AddYAML(data interface{}) *Request {
	return r.
		use(BodyYAML{Data: data})
}

// AddCookiesMap ...
func (r *Request) AddCookiesMap(data map[string]string) *Request {
	return r.
		use(Cookies{Map: data})
}

// AddURLEncodedForm ...
func (r *Request) AddURLEncodedForm(data interface{}) *Request {
	return r.
		use(BodyURLEncodedForm{Data: data})
}

// AddTLSConfig ...
func (r *Request) AddTLSConfig(data *tls.Config) *Request {
	return r.
		use(TLSConfig{data})
}

// AddTransform ...
func (r *Request) AddTransform(data http.RoundTripper) *Request {
	return r.
		use(Transport{RoundTripper: data})
}

// AddMultipartForm ...
func (r *Request) AddMultipartForm(fields map[string]string, files []string) *Request {
	return r.
		use(BodyForm{Fields: fields, Files: files})
}

// Send ...
func (r *Request) Send() *response.Sugar {
	return response.New(r.ctx).Do()
}
