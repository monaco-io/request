package request

import (
	"crypto/tls"
	"net/http"
	"testing"
	"time"
)

func TestClient_Send(t *testing.T) {
	type fields struct {
		URL          string
		Method       string
		Header       map[string]string
		Query        map[string]string
		JSON         interface{}
		XML          interface{}
		String       string
		BasicAuth    BasicAuth
		CustomerAuth string
		Bearer       string
		Timeout      time.Duration
		TLSTimeout   time.Duration
		DialTimeout  time.Duration
		ProxyURL     string
		ProxyServers map[string]string
		Cookies      []*http.Cookie
		CookiesMap   map[string]string
		TLSConfig    *tls.Config
		Transport    *http.Transport
	}
	tests := []struct {
		name      string
		fields    fields
		wantError bool
	}{
		{
			fields: fields{
				URL:         "http://httpbin.org/post",
				Method:      "POST",
				Header:      map[string]string{"google": "google"},
				Query:       map[string]string{"google": "google"},
				JSON:        map[string]string{"google": "google"},
				BasicAuth:   BasicAuth{Username: "google", Password: "google"},
				Timeout:     time.Second * 10,
				TLSTimeout:  time.Second * 10,
				DialTimeout: time.Second * 10,
				ProxyURL:    "http://www.google.com",
				CookiesMap:  map[string]string{"google": "google"},
				TLSConfig:   &tls.Config{},
				Transport:   &http.Transport{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				URL:          tt.fields.URL,
				Method:       tt.fields.Method,
				Header:       tt.fields.Header,
				Query:        tt.fields.Query,
				JSON:         tt.fields.JSON,
				XML:          tt.fields.XML,
				String:       tt.fields.String,
				BasicAuth:    tt.fields.BasicAuth,
				CustomerAuth: tt.fields.CustomerAuth,
				Bearer:       tt.fields.Bearer,
				Timeout:      tt.fields.Timeout,
				TLSTimeout:   tt.fields.TLSTimeout,
				DialTimeout:  tt.fields.DialTimeout,
				ProxyURL:     tt.fields.ProxyURL,
				ProxyServers: tt.fields.ProxyServers,
				Cookies:      tt.fields.Cookies,
				CookiesMap:   tt.fields.CookiesMap,
				TLSConfig:    tt.fields.TLSConfig,
				Transport:    tt.fields.Transport,
			}
			if got := c.Send().Error(); got == nil == tt.wantError {
				t.Errorf("Client.Send() = %v, want %v", got, tt.wantError)
			}
		})
	}
}

func TestClient_Send_Form(t *testing.T) {
	type S struct {
		Data        string `json:"data"`
		ContentType string `json:"contentType"`
	}

	var result S
	resp := (&Client{
		// URL:         "http://httpbin.org/post",
		URL:            "http://localhost:8080/ping",
		Method:         "POST",
		Header:         map[string]string{"google": "google"},
		Query:          map[string]string{"google": "google"},
		URLEncodedForm: map[string]string{"data": "ddd"},
		BasicAuth:      BasicAuth{Username: "google", Password: "google"},
		Timeout:        time.Second * 5,
		TLSTimeout:     time.Second * 5,
		DialTimeout:    time.Second * 5,
		CookiesMap:     map[string]string{"google": "google"},
	}).
		Send().
		Scan(&result)

	if !resp.OK() {
		t.Error(resp.Error())
	}
	// t.Log(resp.String())
	t.Log(result)
}
