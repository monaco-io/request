package request

import (
	"crypto/tls"
	"net/http"
	"testing"
	"time"
)

func TestClient_applyProxy(t *testing.T) {
	type fields struct {
		URL         string
		Method      string
		Header      map[string]string
		Params      map[string]string
		Body        []byte
		BasicAuth   BasicAuth
		Timeout     time.Duration
		ProxyURL    string
		ContentType ContentType
		Cookies     []*http.Cookie
		client      *http.Client
		requestURL  requestURL
		req         *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "pass",
			fields: fields{
				ProxyURL: "http://proxy.com",
			},
			wantErr: false,
		},
		{
			name: "fail",
			fields: fields{
				ProxyURL: " http://proxy.com",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				URL:         tt.fields.URL,
				Method:      tt.fields.Method,
				Header:      tt.fields.Header,
				Params:      tt.fields.Params,
				Body:        tt.fields.Body,
				BasicAuth:   tt.fields.BasicAuth,
				Timeout:     tt.fields.Timeout,
				ProxyURL:    tt.fields.ProxyURL,
				ContentType: tt.fields.ContentType,
				Cookies:     tt.fields.Cookies,
				client:      tt.fields.client,
				requestURL:  tt.fields.requestURL,
				req:         tt.fields.req,
			}
			c.applyClient()
			c.buildRequest()
			if err := c.applyProxy(); (err != nil) != tt.wantErr {
				t.Errorf("Client.applyProxy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_applyTLSConfig(t *testing.T) {
	type fields struct {
		URL         string
		Method      string
		Header      map[string]string
		Params      map[string]string
		Body        []byte
		BasicAuth   BasicAuth
		Timeout     time.Duration
		ProxyURL    string
		ContentType ContentType
		Cookies     []*http.Cookie
		TLSConfig   *tls.Config
		Transport   *http.Transport
		client      *http.Client
		requestURL  requestURL
		req         *http.Request
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "pass",
			fields: fields{
				TLSConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
		{
			name: "fail",
			fields: fields{
				TLSConfig: &tls.Config{InsecureSkipVerify: false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				URL:         tt.fields.URL,
				Method:      tt.fields.Method,
				Header:      tt.fields.Header,
				Params:      tt.fields.Params,
				Body:        tt.fields.Body,
				BasicAuth:   tt.fields.BasicAuth,
				Timeout:     tt.fields.Timeout,
				ProxyURL:    tt.fields.ProxyURL,
				ContentType: tt.fields.ContentType,
				Cookies:     tt.fields.Cookies,
				TLSConfig:   tt.fields.TLSConfig,
				client:      tt.fields.client,
				requestURL:  tt.fields.requestURL,
				req:         tt.fields.req,
				Transport:   tt.fields.Transport,
			}
			c.buildRequest()
			c.applyTLSConfig()
		})
	}
}
