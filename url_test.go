package request

import (
	"net/url"
	"testing"
)

func Test_requestURL_EncodeURL(t *testing.T) {
	type fields struct {
		httpURL    *url.URL
		urlString  string
		parameters map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				urlString:  "https://google.com",
				parameters: map[string]string{"a": "1", "b": "2"},
			},
			wantErr: false,
		},
		{
			name: "success",
			fields: fields{
				urlString:  "https://google.com/search",
				parameters: map[string]string{"a": "1", "b": "2"},
			},
			wantErr: false,
		},
		{
			name: "failed",
			fields: fields{
				urlString:  " https://google.com",
				parameters: map[string]string{"a": "1", "b": "2"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ru := &requestURL{
				httpURL:    tt.fields.httpURL,
				urlString:  tt.fields.urlString,
				parameters: tt.fields.parameters,
			}
			if err := ru.EncodeURL(); (err != nil) != tt.wantErr {
				t.Errorf("requestURL.EncodeURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_requestURL_string(t *testing.T) {
	type fields struct {
		httpURL    *url.URL
		urlString  string
		parameters map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				urlString:  "https://google.com",
				parameters: map[string]string{"a": "1", "b": "2"},
			},
			want: "https://google.com?a=1&b=2",
		},
		{
			name: "success",
			fields: fields{
				urlString:  "https://google.com/search",
				parameters: map[string]string{"a": "1", "b": "2"},
			},
			want: "https://google.com/search?a=1&b=2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ru := requestURL{
				httpURL:    tt.fields.httpURL,
				urlString:  tt.fields.urlString,
				parameters: tt.fields.parameters,
			}
			ru.EncodeURL()
			if got := ru.string(); got != tt.want {
				t.Errorf("requestURL.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_requestURL_schema(t *testing.T) {
	type fields struct {
		httpURL    *url.URL
		urlString  string
		parameters map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "https1",
			fields: fields{
				urlString:  "https://google.com",
				parameters: map[string]string{"a": "1", "b": "2"},
			},
			want: "https",
		},
		{
			name: "https2",
			fields: fields{
				urlString:  "https://google.com/search",
				parameters: map[string]string{"a": "1", "b": "2"},
			},
			want: "https",
		},
		{
			name: "http1",
			fields: fields{
				urlString:  "http://google.com",
				parameters: map[string]string{"a": "1", "b": "2"},
			},
			want: "http",
		},
		{
			name: "http2",
			fields: fields{
				urlString:  "http://google.com/search",
				parameters: map[string]string{"a": "1", "b": "2"},
			},
			want: "http",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ru := requestURL{
				httpURL:    tt.fields.httpURL,
				urlString:  tt.fields.urlString,
				parameters: tt.fields.parameters,
			}
			ru.EncodeURL()
			if got := ru.scheme(); got != tt.want {
				t.Errorf("requestURL.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_requestURL_host(t *testing.T) {
	type fields struct {
		httpURL    *url.URL
		urlString  string
		parameters map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "localhost",
			fields: fields{
				urlString:  "https://localhost",
				parameters: map[string]string{"a": "1", "b": "2"},
			},
			want: "localhost",
		},
		{
			name: "google.com",
			fields: fields{
				urlString:  "https://google.com/search",
				parameters: map[string]string{"a": "1", "b": "2"},
			},
			want: "google.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ru := requestURL{
				httpURL:    tt.fields.httpURL,
				urlString:  tt.fields.urlString,
				parameters: tt.fields.parameters,
			}
			ru.EncodeURL()
			if got := ru.host(); got != tt.want {
				t.Errorf("requestURL.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
