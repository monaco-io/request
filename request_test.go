package request

import (
	originContext "context"
	"crypto/tls"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/monaco-io/request/request"
)

func TestClient_Send(t *testing.T) {
	type fields struct {
		URL          string
		Method       string
		Header       map[string]string
		SortedHeader [][2]string
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
				Method:      POST,
				Header:      map[string]string{"google": "google"},
				Query:       map[string]string{"google": "google"},
				JSON:        map[string]string{"google": "google"},
				BasicAuth:   BasicAuth{Username: "google", Password: "google"},
				Timeout:     time.Second * 10,
				TLSTimeout:  time.Second * 10,
				DialTimeout: time.Second * 10,
				CookiesMap:  map[string]string{"google": "google"},
				TLSConfig:   &tls.Config{},
				Transport:   &http.Transport{},
			},
		},
		{
			name: "SortedHeader",
			fields: fields{
				URL:          "http://httpbin.org/post",
				Method:       POST,
				SortedHeader: [][2]string{{"A", "A"}, {"B", "B"}, {"C", "C"}},
				Query:        map[string]string{"google": "google"},
				JSON:         map[string]string{"google": "google"},
				BasicAuth:    BasicAuth{Username: "google", Password: "google"},
				Timeout:      time.Second * 10,
				TLSTimeout:   time.Second * 10,
				DialTimeout:  time.Second * 10,
				CookiesMap:   map[string]string{"google": "google"},
				TLSConfig:    &tls.Config{},
				Transport:    &http.Transport{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				URL:          tt.fields.URL,
				Method:       tt.fields.Method,
				Header:       tt.fields.Header,
				SortedHeader: tt.fields.SortedHeader,
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

func TestClient_Send_With_Context(t *testing.T) {
	type fields struct {
		URL          string
		Method       string
		Header       map[string]string
		SortedHeader [][2]string
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
		Context      originContext.Context
	}
	tests := []struct {
		name      string
		fields    fields
		wantError bool
	}{
		{
			fields: fields{
				Context:     originContext.TODO(),
				URL:         "http://httpbin.org/post",
				Method:      POST,
				Header:      map[string]string{"google": "google"},
				Query:       map[string]string{"google": "google"},
				JSON:        map[string]string{"google": "google"},
				BasicAuth:   BasicAuth{Username: "google", Password: "google"},
				Timeout:     time.Second * 10,
				TLSTimeout:  time.Second * 10,
				DialTimeout: time.Second * 10,
				CookiesMap:  map[string]string{"google": "google"},
			},
		},
		{
			name: "SortedHeader",
			fields: fields{
				Context:      originContext.TODO(),
				URL:          "http://httpbin.org/post",
				Method:       POST,
				SortedHeader: [][2]string{{"A", "A"}, {"B", "B"}, {"C", "C"}},
				Query:        map[string]string{"google": "google"},
				JSON:         map[string]string{"google": "google"},
				BasicAuth:    BasicAuth{Username: "google", Password: "google"},
				Timeout:      time.Second * 10,
				TLSTimeout:   time.Second * 10,
				DialTimeout:  time.Second * 10,
				CookiesMap:   map[string]string{"google": "google"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				URL:          tt.fields.URL,
				Method:       tt.fields.Method,
				Header:       tt.fields.Header,
				SortedHeader: tt.fields.SortedHeader,
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
	var result struct {
		Data        string `json:"google"`
		ContentType string `json:"contentType"`
	}

	resp := (&Client{
		URL:            "http://httpbin.org/post",
		Method:         POST,
		Header:         map[string]string{"google": "google"},
		Query:          map[string]string{"google": "google"},
		URLEncodedForm: map[string]string{"google": "google"},
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
}

func TestRequest_Send(t *testing.T) {
	var data map[string]interface{}
	resp := New().
		POST("http://httpbin.org/post").
		AddHeader(map[string]string{"Google": "google"}).
		AddBasicAuth("google", "google").
		AddURLEncodedForm(map[string]string{"data": "google"}).
		Send().
		Scan(&data)

	if !resp.OK() {
		t.Error(resp.Error())
	}

	if data["headers"].(map[string]interface{})["Authorization"].(string) != "Basic Z29vZ2xlOmdvb2dsZQ==" {
		t.Error("Authorization")
	}

	if data["form"].(map[string]interface{})["data"] != "google" {
		t.Error("form")
	}
}

func TestContext_TraceDo(t *testing.T) {
	var data map[string]interface{}
	body := make(map[string]string)
	for i := 0; i < 2000; i++ {
		body[fmt.Sprint(i)] = fmt.Sprint(i)
	}
	resp := New().
		POST("http://httpbin.org/post").
		AddHeader(map[string]string{"Google": "google"}).
		AddBasicAuth("google", "google").
		AddJSON(body).
		Send().
		Scan(&data)

	if !resp.OK() {
		t.Error(resp.Error())
	}

	if resp.TimeTrace().Duration == 0 {
		t.Fail()
	}
}

func TestJSONBody(t *testing.T) {
	type fields struct {
		URL          string
		Method       string
		Header       map[string]string
		SortedHeader [][2]string
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
				URL:       "http://httpbin.org/post",
				Method:    POST,
				JSON:      map[string]interface{}{"A": "A", "B": 001},
				BasicAuth: BasicAuth{Username: "google", Password: "google"},
			},
		},
		{
			name: "SortedHeader",
			fields: fields{
				URL:    "http://httpbin.org/post",
				Method: POST,
				JSON: struct {
					A string
					B int
				}{A: "A", B: 001},
				BasicAuth: BasicAuth{Username: "google", Password: "google"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				URL:          tt.fields.URL,
				Method:       tt.fields.Method,
				Header:       tt.fields.Header,
				SortedHeader: tt.fields.SortedHeader,
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

func TestFileBody(t *testing.T) {
	type fields struct {
		URL          string
		Method       string
		Header       map[string]string
		SortedHeader [][2]string
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
		FormFields   map[string]string
		FormFiles    []string
	}
	tests := []struct {
		name      string
		fields    fields
		wantError bool
	}{
		{
			name: "form field",
			fields: fields{
				URL:        "http://httpbin.org/post",
				Method:     POST,
				FormFields: map[string]string{"field": "value"},
				BasicAuth:  BasicAuth{Username: "google", Password: "google"},
			},
			wantError: false,
		},
		{
			name: "form file",
			fields: fields{
				URL:        "http://httpbin.org/post",
				Method:     POST,
				FormFields: map[string]string{"field": "value"},
				FormFiles:  []string{"not_exist.txt"},
			},
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				URL:           tt.fields.URL,
				Method:        tt.fields.Method,
				Header:        tt.fields.Header,
				SortedHeader:  tt.fields.SortedHeader,
				Query:         tt.fields.Query,
				JSON:          tt.fields.JSON,
				XML:           tt.fields.XML,
				String:        tt.fields.String,
				BasicAuth:     tt.fields.BasicAuth,
				CustomerAuth:  tt.fields.CustomerAuth,
				Bearer:        tt.fields.Bearer,
				Timeout:       tt.fields.Timeout,
				TLSTimeout:    tt.fields.TLSTimeout,
				DialTimeout:   tt.fields.DialTimeout,
				ProxyURL:      tt.fields.ProxyURL,
				ProxyServers:  tt.fields.ProxyServers,
				Cookies:       tt.fields.Cookies,
				CookiesMap:    tt.fields.CookiesMap,
				TLSConfig:     tt.fields.TLSConfig,
				Transport:     tt.fields.Transport,
				MultipartForm: MultipartForm{tt.fields.FormFields, tt.fields.FormFiles},
			}
			resp := c.Send()
			if got := resp.Error(); got == nil == tt.wantError {
				t.Errorf("Client.Send() = %v, want %v", got, tt.wantError)
			}
		})
	}
}

func TestNewWithContext(t *testing.T) {
	type args struct {
		ctx originContext.Context
	}
	tests := []struct {
		name string
		args args
		want *request.Request
	}{
		{
			args: args{
				ctx: originContext.Background(),
			},
			want: NewWithContext(originContext.Background()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWithContext(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWithContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
