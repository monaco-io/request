package request

import (
	"net"
	"net/http"
	"testing"
	"time"
)

var serverURL = "http://httpbin.org"

func TestClient_Do(t *testing.T) {
	type fields struct {
		URL    string
		Method string
		Params map[string]string
		Header map[string]string
		Body   []byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "get",
			fields: fields{
				URL:    serverURL + "/get",
				Method: "GET",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "get with params",
			fields: fields{
				URL:    serverURL + "/get",
				Method: "GET",
				Params: map[string]string{"package": "request"},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "get with 2 params",
			fields: fields{
				URL:    serverURL + "/get",
				Method: "GET",
				Params: map[string]string{"package": "request", "lang": "golang"},
			},
			want:    nil,
			wantErr: false,
		},

		{
			name: "post",
			fields: fields{
				URL:    serverURL + "/post",
				Method: "POST",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "post with params",
			fields: fields{
				URL:    serverURL + "/post",
				Method: "POST",
				Params: map[string]string{"package": "request"},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "post with 2 params",
			fields: fields{
				URL:    serverURL + "/post",
				Method: "POST",
				Params: map[string]string{"package": "request"},
				Header: map[string]string{"Content-Type": "application/json"},
				Body:   []byte(`{"data1":1,"data2":2}`),
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				URL:    tt.fields.URL,
				Method: tt.fields.Method,
				Params: tt.fields.Params,
				Header: tt.fields.Header,
				Body:   tt.fields.Body,
			}
			_, err := c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClient_Do_Get(t *testing.T) {
	c := Client{
		URL:       serverURL + "/get",
		Method:    "GET",
		Header:    map[string]string{"Customer-Header": "c-h-value"},
		Params:    map[string]string{"v1": "1", "v2": "2"},
		BasicAuth: BasicAuth{Username: "username", Password: "password"},
	}
	resp, err := c.Do()
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.Code, string(resp.Data))
}

// TODO check response json
func TestClient_Do_Post_Json(t *testing.T) {
	c := Client{
		URL:         serverURL + "/post",
		Method:      "POST",
		Header:      map[string]string{"Customer-Header": "c-h-value"},
		Params:      map[string]string{"v1": "1", "v2": "2"},
		ContentType: ApplicationXWwwFormURLEncoded,
		Body:        []byte(`{"v3":3, "v4"="4"`),
		BasicAuth:   BasicAuth{Username: "username", Password: "password"},
	}
	resp, err := c.Do()
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.Code, string(resp.Data))
}

// TODO check response form
func TestClient_Do_Post(t *testing.T) {
	c := Client{
		URL:       serverURL + "/post",
		Method:    "POST",
		Header:    map[string]string{"Customer-Header": "c-h-value"},
		Params:    map[string]string{"v1": "1", "v2": "2"},
		Body:      []byte(`{"v3":3, "v4"="4"`),
		BasicAuth: BasicAuth{Username: "username", Password: "password"},
	}
	resp, err := c.Do()
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.Code, string(resp.Data))
}

func TestClient_Do_Encode_Error(t *testing.T) {
	c := Client{
		URL:    " " + serverURL + "/post",
		Method: "POST",
	}
	_, err := c.Do()
	if err == nil {
		t.Error(err)
	}
}

func TestClient_Do_Build_Request_Error(t *testing.T) {
	c := Client{
		URL:    serverURL + "/post",
		Method: " ",
	}
	_, err := c.Do()
	if err == nil {
		t.Error(err)
	}
}

func TestClient_Do_Resp_Error(t *testing.T) {
	c := Client{
		URL:    "http://localhost:1",
		Method: "POST",
	}
	_, err := c.Do()
	if err == nil {
		t.Error(err)
	}
}

func TestClient_buildRequest(t *testing.T) {
	type fields struct {
		URL         string
		Method      string
		Header      map[string]string
		Params      map[string]string
		Body        []byte
		BasicAuth   BasicAuth
		Timeout     time.Duration
		ContentType ContentType
		client      *http.Client
		req         *http.Request
		resp        *http.Response
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "TestClient_buildRequest_encode_error",
			fields: fields{
				URL:    " " + serverURL + "/post",
				Method: "POST",
			},
			wantErr: true,
		},
		{
			name: "TestClient_buildRequest_encode_error",
			fields: fields{
				URL:    serverURL + "/post",
				Method: " ",
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
				ContentType: tt.fields.ContentType,
				client:      tt.fields.client,
				req:         tt.fields.req,
			}
			if err := c.buildRequest(); (err != nil) != tt.wantErr {
				t.Errorf("buildRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_Resp(t *testing.T) {
	type fields struct {
		URL         string
		Method      string
		Header      map[string]string
		Params      map[string]string
		Body        []byte
		BasicAuth   BasicAuth
		Timeout     time.Duration
		ContentType ContentType
		client      *http.Client
		req         *http.Request
		resp        *http.Response
	}
	tests := []struct {
		name     string
		fields   fields
		wantResp *http.Response
		wantErr  bool
	}{
		{
			name: "TestClient_Resp",
			fields: fields{
				URL: serverURL,
			},
			wantResp: nil,
			wantErr:  false,
		},
		{
			name: "TestClient_Resp_error",
			fields: fields{
				URL:    "http://localhost:1",
				Method: "POST",
			},
			wantResp: nil,
			wantErr:  true,
		},
		{
			name: "TestClient_Resp_build_error",
			fields: fields{
				URL:    " " + serverURL + "/post",
				Method: "POST",
			},
			wantResp: nil,
			wantErr:  true,
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
				ContentType: tt.fields.ContentType,
				client:      tt.fields.client,
				req:         tt.fields.req,
			}
			_, err := c.Resp()
			if (err != nil) != tt.wantErr {
				t.Errorf("Resp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_Close(t *testing.T) {
	type fields struct {
		URL         string
		Method      string
		Header      map[string]string
		Params      map[string]string
		Body        []byte
		BasicAuth   BasicAuth
		Timeout     time.Duration
		ContentType ContentType
		client      *http.Client
		req         *http.Request
		resp        *http.Response
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "TestClient_Close",
			fields: fields{
				URL:    serverURL + "/post",
				Method: "POST",
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
				ContentType: tt.fields.ContentType,
				client:      tt.fields.client,
				req:         tt.fields.req,
			}
			resp, _ := c.Do()
			resp.Close()
		})
	}
}

func TestClient_Do_Timeout(t *testing.T) {
	type fields struct {
		URL     string
		Method  string
		Params  map[string]string
		Header  map[string]string
		Body    []byte
		Timeout time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "TestClient_Do_Proxy_succeed",
			fields: fields{
				URL:     serverURL + "/get",
				Method:  "GET",
				Timeout: time.Second,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "TestClient_Do_Proxy_err",
			fields: fields{
				URL:     "http://1.com",
				Method:  "GET",
				Timeout: 1,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				URL:     tt.fields.URL,
				Method:  tt.fields.Method,
				Params:  tt.fields.Params,
				Header:  tt.fields.Header,
				Body:    tt.fields.Body,
				Timeout: tt.fields.Timeout,
			}
			_, err := c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClient_Do_Proxy(t *testing.T) {
	type fields struct {
		URL      string
		Method   string
		Params   map[string]string
		Header   map[string]string
		Body     []byte
		ProxyURL string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "err",
			fields: fields{
				URL:      serverURL + "/get",
				Method:   "GET",
				ProxyURL: "http://www.1.com",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				URL:      serverURL + "/get",
				Method:   "GET",
				ProxyURL: "http://www.baidu.com",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				URL:      tt.fields.URL,
				Method:   tt.fields.Method,
				Params:   tt.fields.Params,
				Header:   tt.fields.Header,
				Body:     tt.fields.Body,
				ProxyURL: tt.fields.ProxyURL,
			}
			_, err := c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClient_Do_Cookies(t *testing.T) {
	type fields struct {
		URL     string
		Method  string
		Cookies []*http.Cookie
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "cookie",
			fields: fields{
				URL:    serverURL + "/get",
				Method: "GET",
				Cookies: []*http.Cookie{
					{
						Name:  "cookie_name",
						Value: "cookie_value",
					},
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				URL:     tt.fields.URL,
				Method:  tt.fields.Method,
				Cookies: tt.fields.Cookies,
			}
			_, err := c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClient_Do_Transport(t *testing.T) {

	type fields struct {
		URL       string
		Method    string
		Params    map[string]string
		Header    map[string]string
		Body      []byte
		Transport *http.Transport
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "TestClient_Do_Transport_succeed",
			fields: fields{
				URL:    serverURL + "/get",
				Method: "GET",
				Transport: &http.Transport{
					Proxy: http.ProxyFromEnvironment,
					DialContext: (&net.Dialer{
						Timeout:   30 * time.Second,
						KeepAlive: 30 * time.Second,
						DualStack: true,
					}).DialContext,
					ForceAttemptHTTP2:     true,
					MaxIdleConns:          200,
					IdleConnTimeout:       90 * time.Second,
					TLSHandshakeTimeout:   10 * time.Second,
					ExpectContinueTimeout: 1 * time.Second,
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				URL:       tt.fields.URL,
				Method:    tt.fields.Method,
				Params:    tt.fields.Params,
				Header:    tt.fields.Header,
				Body:      tt.fields.Body,
				Transport: tt.fields.Transport,
			}
			_, err := c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
