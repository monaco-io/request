package request

import (
	"testing"
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
	t.Log(string(resp))
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
	t.Log(string(resp))
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
	t.Log(string(resp))
}

func TestClient_Do_Encode_Error(t *testing.T) {
	c := Client{
		URL:    " " + serverURL + "/POST",
		Method: "POST",
	}
	_, err := c.Do()
	if err == nil {
		t.Error(err)
	}
}

func TestClient_Do_Build_Request_Error(t *testing.T) {
	c := Client{
		URL:    serverURL + "/POST",
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
