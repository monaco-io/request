package request

import (
	"net/http"
	"testing"
	"time"
)

var serverURL = "http://httpbin.org"

func TestClient_Do_Error(t *testing.T) {
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
			name: "TestClient_Do_GET",
			fields: fields{
				URL:    serverURL + "/get",
				Method: "GET",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "TestClient_Do_GET_1params",
			fields: fields{
				URL:    serverURL + "/get",
				Method: "GET",
				Params: map[string]string{"package": "request"},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "TestClient_Do_GET_2params",
			fields: fields{
				URL:    serverURL + "/get",
				Method: "GET",
				Params: map[string]string{"package": "request", "lang": "golang"},
			},
			want:    nil,
			wantErr: false,
		},

		{
			name: "TestClient_Do_POST",
			fields: fields{
				URL:    serverURL + "/post",
				Method: "POST",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "TestClient_Do_POST_1params",
			fields: fields{
				URL:    serverURL + "/post",
				Method: "POST",
				Params: map[string]string{"package": "request"},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "TestClient_Do_POST_1params_body",
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
		//{
		//	name: "TestClient_Resp",
		//	fields: fields{
		//		URL: serverURL,
		//	},
		//	wantResp: nil,
		//	wantErr:  false,
		//},
		//{
		//	name: "TestClient_Resp_error",
		//	fields: fields{
		//		URL:    "http://localhost:1",
		//		Method: "POST",
		//	},
		//	wantResp: nil,
		//	wantErr:  true,
		//},
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

func TestClient_StatusCode(t *testing.T) {
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
		wantCode int
	}{
		{
			name: "TestClient_StatusCode_200",
			fields: fields{
				URL:    serverURL + "/post",
				Method: "POST",
			},
			wantCode: 200,
		},
		{
			name: "TestClient_StatusCode_500",
			fields: fields{
				URL:    serverURL + "/get",
				Method: "POST",
			},
			wantCode: 405,
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
			if gotCode := resp.StatusCode(); gotCode != tt.wantCode {
				t.Errorf("StatusCode() = %v, want %v", gotCode, tt.wantCode)
			}
		})
	}
}

func TestSugaredResp_Status(t *testing.T) {
	type fields struct {
		Data []byte
		Code int
		resp *http.Response
	}
	tests := []struct {
		name       string
		fields     fields
		wantStatus string
	}{
		{
			name: "TestSugaredResp_Status_200_ok",
			fields: fields{
				Data: nil,
				Code: 0,
				resp: nil,
			},
			wantStatus: "200 OK",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				URL: serverURL + "/get",
			}
			s, _ := c.Do()
			if gotStatus := s.Status(); gotStatus != tt.wantStatus {
				t.Errorf("Status() = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}
