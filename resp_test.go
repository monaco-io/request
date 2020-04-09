package request

import (
	"net/http"
	"testing"
	"time"
)

func TestSugaredResp_StatusCode(t *testing.T) {
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
