package request

import (
	"reflect"
	"testing"
)

// TODO: Mock Http Server
var serverURL = "http://localhost:3000"

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
			name: "TestClient_Do_GET",
			fields: fields{
				URL:    serverURL,
				Method: "GET",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "TestClient_Do_GET_1params",
			fields: fields{
				URL:    serverURL,
				Method: "GET",
				Params: map[string]string{"package": "request"},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "TestClient_Do_GET_2params",
			fields: fields{
				URL:    serverURL,
				Method: "GET",
				Params: map[string]string{"package": "request", "lang": "golang"},
			},
			want:    nil,
			wantErr: false,
		},

		{
			name: "TestClient_Do_POST",
			fields: fields{
				URL:    serverURL,
				Method: "POST",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "TestClient_Do_POST_1params",
			fields: fields{
				URL:    serverURL,
				Method: "POST",
				Params: map[string]string{"package": "request"},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "TestClient_Do_POST_1params_body",
			fields: fields{
				URL:    serverURL,
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
			got, err := c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Do() got = %v, want %v", string(got), tt.want)
			}
		})
	}
}
