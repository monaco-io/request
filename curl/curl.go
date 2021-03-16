package curl

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
)

func bashEscape(str string) string {
	return `'` + strings.Replace(str, `'`, `'\''`, -1) + `'`
}

// GetCommand build curl command by http request
func GetCommand(req *http.Request) (cmd string, err error) {
	var (
		command []string
		keys    []string
		body    []byte
		space   = " "
	)

	command = append(command, "curl", "-X", bashEscape(req.Method))

	if req.Body != nil {
		body, err = io.ReadAll(req.Body)
		if err != nil {
			return
		}
		req.Body = io.NopCloser(bytes.NewBuffer(body))
		if len(string(body)) > 0 {
			bodyEscaped := bashEscape(string(body))
			command = append(command, "-d", bodyEscaped)
		}
	}

	for k := range req.Header {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		command = append(command, "-H", bashEscape(fmt.Sprintf("%s: %s", k, strings.Join(req.Header[k], space))))
	}

	command = append(command, bashEscape(req.URL.String()))
	cmd = strings.Join(command, space)
	return
}
