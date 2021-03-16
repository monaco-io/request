package request

import (
	"crypto/tls"
	"net/http"
	"testing"
	"time"
)

func TestClient_PrintCURL(t *testing.T) {
	type fields struct {
		URL            string
		Method         string
		Header         map[string]string
		SortedHeader   [][2]string
		Query          map[string]string
		JSON           interface{}
		XML            interface{}
		YAML           interface{}
		String         string
		URLEncodedForm interface{}
		MultipartForm  MultipartForm
		BasicAuth      BasicAuth
		CustomerAuth   string
		Bearer         string
		Timeout        time.Duration
		TLSTimeout     time.Duration
		DialTimeout    time.Duration
		ProxyURL       string
		ProxyServers   map[string]string
		Cookies        []*http.Cookie
		CookiesMap     map[string]string
		TLSConfig      *tls.Config
		Transport      *http.Transport
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "print curl",
			fields: fields{
				URL:    "https://www.google.com",
				Method: "POST",
				Header: map[string]string{
					"h1": "hv1",
				},
				Query: map[string]string{
					"q1": "qv1",
				},
				JSON:          `{"hello":"world"}`,
				MultipartForm: MultipartForm{},
				BasicAuth: BasicAuth{
					Username: "admin",
					Password: "admin123",
				},
				CookiesMap: map[string]string{
					"c1": "cv1",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				URL:            tt.fields.URL,
				Method:         tt.fields.Method,
				Header:         tt.fields.Header,
				SortedHeader:   tt.fields.SortedHeader,
				Query:          tt.fields.Query,
				JSON:           tt.fields.JSON,
				XML:            tt.fields.XML,
				YAML:           tt.fields.YAML,
				String:         tt.fields.String,
				URLEncodedForm: tt.fields.URLEncodedForm,
				MultipartForm:  tt.fields.MultipartForm,
				BasicAuth:      tt.fields.BasicAuth,
				CustomerAuth:   tt.fields.CustomerAuth,
				Bearer:         tt.fields.Bearer,
				Timeout:        tt.fields.Timeout,
				TLSTimeout:     tt.fields.TLSTimeout,
				DialTimeout:    tt.fields.DialTimeout,
				ProxyURL:       tt.fields.ProxyURL,
				ProxyServers:   tt.fields.ProxyServers,
				Cookies:        tt.fields.Cookies,
				CookiesMap:     tt.fields.CookiesMap,
				TLSConfig:      tt.fields.TLSConfig,
				Transport:      tt.fields.Transport,
			}
			c.PrintCURL()
		})
	}
}
