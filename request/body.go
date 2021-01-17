package request

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/url"
	"strconv"
	"strings"

	"github.com/monaco-io/request/context"
	"gopkg.in/yaml.v2"
)

// BodyString body of type string
type BodyString struct {
	Data string
}

// Apply string body
func (b BodyString) Apply(ctx *context.Context) {
	bBytes := bytes.NewReader([]byte(b.Data))
	rc, ok := io.Reader(bBytes).(io.ReadCloser)
	if !ok && bBytes != nil {
		rc = ioutil.NopCloser(bBytes)
	}

	ctx.Request.Body = rc
	ctx.Request.ContentLength = int64(bytes.NewBufferString(b.Data).Len())
}

// Valid string body valid?
func (b BodyString) Valid() bool {
	if b.Data == "" {
		return false
	}
	return true
}

// BodyJSON body of type json
type BodyJSON struct {
	Data interface{}
}

// Apply json body
func (b BodyJSON) Apply(ctx *context.Context) {
	buf := &bytes.Buffer{}

	switch b.Data.(type) {
	case string:
		buf.WriteString(b.Data.(string))
	case []byte:
		buf.Write(b.Data.([]byte))
	default:
		if err := json.NewEncoder(buf).Encode(b.Data); err != nil {
			fmt.Println(err)
			return
		}
	}

	ctx.Request.Body = ioutil.NopCloser(buf)
	ctx.Request.ContentLength = int64(buf.Len())
	setContentType(ctx, JSON)
}

// Valid json body valid?
func (b BodyJSON) Valid() bool {
	if b.Data == nil {
		return false
	}
	return true
}

// BodyXML body of type xml
type BodyXML struct {
	Data interface{}
}

// Apply xml body
func (b BodyXML) Apply(ctx *context.Context) {
	buf := &bytes.Buffer{}

	switch b.Data.(type) {
	case string:
		buf.WriteString(b.Data.(string))
	case []byte:
		buf.Write(b.Data.([]byte))
	default:
		if err := xml.NewEncoder(buf).Encode(b.Data); err != nil {
			ctx.SetError(fmt.Errorf("unknown json encoded type: %T", b.Data))
			return
		}
	}

	ctx.Request.Body = ioutil.NopCloser(buf)
	ctx.Request.ContentLength = int64(buf.Len())
	setContentType(ctx, XML)
}

// Valid xml body valid?
func (b BodyXML) Valid() bool {
	if b.Data == nil {
		return false
	}
	return true
}

// BodyYAML body of type yaml
type BodyYAML struct {
	Data interface{}
}

// Apply yaml body
func (b BodyYAML) Apply(ctx *context.Context) {
	buf := &bytes.Buffer{}

	switch b.Data.(type) {
	case string:
		buf.WriteString(b.Data.(string))
	case []byte:
		buf.Write(b.Data.([]byte))
	default:
		if err := yaml.NewEncoder(buf).Encode(b.Data); err != nil {
			ctx.SetError(fmt.Errorf("unknown yaml encoded type: %T", b.Data))
			return
		}
	}

	ctx.Request.Body = ioutil.NopCloser(buf)
	ctx.Request.ContentLength = int64(buf.Len())
}

// Valid json body valid?
func (b BodyYAML) Valid() bool {
	if b.Data == nil {
		return false
	}
	return true
}

// BodyURLEncodedForm application/x-www-form-urlencoded
type BodyURLEncodedForm struct {
	Data interface{}
}

// Apply application/x-www-form-urlencoded
func (b BodyURLEncodedForm) Apply(ctx *context.Context) {
	buf := &bytes.Buffer{}

	switch b.Data.(type) {
	case string:
		buf.WriteString(b.Data.(string))
	case []byte:
		buf.Write(b.Data.([]byte))
	case map[string]string:
		data := make(url.Values)
		for k, v := range b.Data.(map[string]string) {
			data.Set(k, v)
		}
		buf.WriteString(data.Encode())
	case map[string][]string:
		buf.WriteString(url.Values(b.Data.(map[string][]string)).Encode())
	case url.Values:
		buf.WriteString(b.Data.(url.Values).Encode())
	default:
		ctx.SetError(fmt.Errorf("unknown urlencoded type: %T", b.Data))
		return
	}

	ctx.Request.Body = ioutil.NopCloser(buf)
	ctx.Request.ContentLength = int64(buf.Len())
	setContentType(ctx, URLEncodedForm)
}

// Valid application/x-www-form-urlencoded valid?
func (b BodyURLEncodedForm) Valid() bool {
	if b.Data == nil {
		return false
	}
	return true
}

// FormFile represents the file form field data.
type FormFile struct {
	Name   string
	Reader io.Reader
}

// BodyForm represents the supported form fields by file and string data.
type BodyForm struct {
	Fields map[string]string
	Files  []FormFile
}

// Apply TODO
func (fd BodyForm) Apply(ctx *context.Context) {
	buf := &bytes.Buffer{}
	multipartWriter := multipart.NewWriter(buf)

	for index, file := range fd.Files {
		if err := writeFile(multipartWriter, fd, file, index); err != nil {
			return
		}
	}

	// Populate the other parts of the form (if there are any)
	for k, v := range fd.Fields {
		multipartWriter.WriteField(k, v)
	}
	if err := multipartWriter.Close(); err != nil {
		return
	}
	if buf.Len() == 0 {
		return
	}

	ctx.Request.Body = ioutil.NopCloser(buf)
	ctx.Request.Header.Add("Content-Type", multipartWriter.FormDataContentType())
	return
}

func writeFile(multipartWriter *multipart.Writer, fd BodyForm, ff FormFile, index int) error {
	if ff.Reader == nil {
		return errors.New("github/monaco-io/request: file reader cannot be nil")
	}

	rc, ok := ff.Reader.(io.ReadCloser)
	if !ok && ff.Reader != nil {
		rc = ioutil.NopCloser(ff.Reader)
	}

	fileName := "file"
	if len(fd.Files) > 1 {
		fileName = strings.Join([]string{fileName, strconv.Itoa(index + 1)}, "")
	}
	if ff.Name != "" {
		fileName = ff.Name
	}

	writer, err := multipartWriter.CreateFormFile(fileName, ff.Name)
	if err != nil {
		return err
	}
	if _, err = io.Copy(writer, rc); err != nil && err != io.EOF {
		return err
	}
	rc.Close()

	return nil
}

// Valid form body valid?
func (fd BodyForm) Valid() bool {
	if fd.Fields == nil && fd.Files == nil {
		return false
	}
	return true
}
