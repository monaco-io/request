package response

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/monaco-io/request/context"
)

// Sugar response with status code and body data
type Sugar struct {
	ctx *context.Context

	// Internal buffer store
	buffer *bytes.Buffer

	done bool
	err  error
}

// New new sugared response
func New(ctx *context.Context) *Sugar {
	return &Sugar{ctx: ctx, buffer: bytes.NewBuffer([]byte{})}
}

func (s *Sugar) setError(err error) {
	s.err = err
}

func (s *Sugar) hasError() bool {
	if s.err != nil {
		return true
	}
	return false
}

func (s *Sugar) resetError() {
	s.err = nil
}

// Error get error
func (s *Sugar) Error() error {
	return s.err
}

// ErrorString error as string
func (s *Sugar) ErrorString() string {
	if s.err != nil {
		return s.err.Error()
	}
	return ""
}

// OK is ok?
func (s *Sugar) OK() bool {
	if s.Error() != nil {
		return false
	}
	return true
}

// Close close http response body
func (s *Sugar) Close() *Sugar {
	if s.hasError() {
		return s
	}
	if _, s.err = io.Copy(ioutil.Discard, s.ctx.Response.Body); s.hasError() {
		return s
	}
	s.setError(s.ctx.Response.Body.Close())
	return s
}

// Do do http request with client
func (s *Sugar) Do() *Sugar {
	if s.done {
		goto OUT
	}
	// send request and close on func call end
	if s.ctx.Response, s.err = s.ctx.Client.Do(s.ctx.Request); s.hasError() {
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

	_, s.err = io.Copy(s.buffer, s.ctx.Response.Body)
	if s.hasError() && s.Error() != io.EOF {
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

// Scan response body to struct(tag: json/xml)
func (s *Sugar) Scan(data interface{}) *Sugar {
	if s.hasError() {
		return s
	}
	firstError := s.ScanJSON(data).Error()
	if firstError != nil && s.ScanXML(data).hasError() {
		s.setError(firstError)
	}
	return s
}

// SaveToFile allows you to download the contents
// of the response to a file.
// TODO test
func (s *Sugar) SaveToFile(fileName string) *Sugar {
	if s.hasError() {
		return s
	}

	fd, err := os.Create(fileName)
	if err != nil {
		s.setError(err)
		goto OUT
	}

	defer s.Close()
	defer fd.Close()

	if _, err = io.Copy(fd, s.buffer); err != nil && err != io.EOF {
		s.setError(err)
		goto OUT
	}

OUT:
	return s
}

// ScanJSON is a method that will populate a struct that is provided `userStruct`
// with the JSON returned within the response body.
func (s *Sugar) ScanJSON(userStruct interface{}) *Sugar {
	if s.hasError() {
		return s
	}

	jsonDecoder := json.NewDecoder(s.buffer)

	defer s.Close()

	if err := jsonDecoder.Decode(&userStruct); err != nil && err != io.EOF {
		s.setError(err)
	}

	return s
}

// ScanXML is a method that will populate a struct that is provided
// `userStruct` with the XML returned within the response body.
func (s *Sugar) ScanXML(userStruct interface{}) *Sugar {
	if s.hasError() {
		return s
	}

	xmlDecoder := xml.NewDecoder(s.buffer)

	defer s.Close()

	if err := xmlDecoder.Decode(&userStruct); err != nil && err != io.EOF {
		s.setError(err)
	}

	return s
}
