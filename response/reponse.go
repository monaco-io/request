package response

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/monaco-io/request/context"
	"gopkg.in/yaml.v2"
)

// Sugar response with status code and body data
type Sugar struct {
	ctx *context.Context

	// Internal buffer store
	buffer *bytes.Buffer

	done bool
}

// New new sugared response
func New(ctx *context.Context) *Sugar {
	return &Sugar{ctx: ctx, buffer: bytes.NewBuffer([]byte{})}
}

// OK is ok?
func (s *Sugar) OK() bool {
	return !s.ctx.HasError()
}

func (s *Sugar) Error() error {
	return s.ctx.Error()
}

// Close close http response body
func (s *Sugar) Close() *Sugar {
	if s.ctx.HasError() {
		return s
	}
	if _, err := io.Copy(ioutil.Discard, s.ctx.Response.Body); err != nil {
		s.ctx.SetError(err)
		return s
	}
	s.ctx.SetError(s.ctx.Response.Body.Close())
	return s
}

// Do do http request with client
func (s *Sugar) Do() *Sugar {
	if s.done || s.ctx.HasError() {
		goto OUT
	}
	// send request and close on func call end
	if err := s.ctx.TraceDo(); err != nil {
		s.ctx.SetError(err)
		goto OUT
	}
	// Have I done this already?
	if s.buffer.Len() != 0 {
		goto OUT

	}
	// Is there any content?
	if s.ctx.Response.ContentLength == 0 {
		goto OUT
	}

	// Did the server tell us how big the response is going to be?
	if s.ctx.Response.ContentLength > 0 {
		s.buffer.Grow(int(s.ctx.Response.ContentLength))
	}

	if _, err := io.Copy(s.buffer, s.ctx.Response.Body); err != nil && err != io.EOF {
		s.ctx.SetError(err)
		s.ctx.Response.Body.Close()
	}
	s.done = true
OUT:
	return s
}

// Response original http response
func (s *Sugar) Response() *http.Response {
	return s.ctx.Response
}

// Code http status code
func (s *Sugar) Code() int {
	return s.ctx.Response.StatusCode
}

// Bytes read response body  as bytes
func (s *Sugar) Bytes() []byte {
	// Is still empty?
	if s.buffer.Len() == 0 {
		return nil
	}
	return s.buffer.Bytes()
}

// String read response body  as string
func (s *Sugar) String() string {
	return s.buffer.String()
}

// ContentType read response content type
func (s *Sugar) ContentType() string {
	return s.Response().Header.Get("Content-Type")
}

// Scan response body to struct(tag: json/xml)
func (s *Sugar) Scan(data interface{}) *Sugar {
	if s.ctx.HasError() {
		return s
	}
	switch ct := s.ContentType(); {
	case context.ContentTypeValid(ct, context.JSON):
		s.ScanJSON(data)
	case context.ContentTypeValid(ct, context.XML):
		s.ScanXML(data)
	default:
		s.ctx.SetError(fmt.Errorf("content type unsupported: %s", ct))
	}
	return s
}

// SaveToFile allows you to download the contents
// of the response to a file.
// TODO test
func (s *Sugar) SaveToFile(fileName string) *Sugar {
	if s.ctx.HasError() {
		return s
	}

	fd, err := os.Create(fileName)
	if err != nil {
		s.ctx.SetError(err)
		goto OUT
	}

	defer s.Close()
	defer fd.Close()

	if _, err = io.Copy(fd, s.buffer); err != nil && err != io.EOF {
		s.ctx.SetError(err)
		goto OUT
	}

OUT:
	return s
}

// ScanJSON is a method that will populate a struct that is provided `userStruct`
// with the JSON returned within the response body.
func (s *Sugar) ScanJSON(userStruct interface{}) *Sugar {
	if s.ctx.HasError() {
		return s
	}

	jsonDecoder := json.NewDecoder(s.buffer)

	defer s.Close()

	if err := jsonDecoder.Decode(userStruct); err != nil && err != io.EOF {
		s.ctx.SetError(err)
	}

	return s
}

// ScanXML is a method that will populate a struct that is provided
// `userStruct` with the XML returned within the response body.
func (s *Sugar) ScanXML(userStruct interface{}) *Sugar {
	if s.ctx.HasError() {
		return s
	}

	xmlDecoder := xml.NewDecoder(s.buffer)

	defer s.Close()

	if err := xmlDecoder.Decode(userStruct); err != nil && err != io.EOF {
		s.ctx.SetError(err)
	}

	return s
}

// ScanYAML is a method that will populate a struct that is provided
// `userStruct` with the yaml returned within the response body.
func (s *Sugar) ScanYAML(userStruct interface{}) *Sugar {
	if s.ctx.HasError() {
		return s
	}

	yamlDecoder := yaml.NewDecoder(s.buffer)

	defer s.Close()

	if err := yamlDecoder.Decode(userStruct); err != nil && err != io.EOF {
		s.ctx.SetError(err)
	}

	return s
}

// TimeTrace ...
func (s *Sugar) TimeTrace() context.Time {
	return s.ctx.TimeTrace
}
