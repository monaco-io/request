package xcontext

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Context
	}{
		{
			want: &Context{
				Request: newRequest(),
				Client:  &http.Client{Transport: http.DefaultTransport},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); got == nil ||
				!reflect.DeepEqual(got.Client, tt.want.Client) ||
				!reflect.DeepEqual(got.Request, tt.want.Request) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWithContext(t *testing.T) {
	tests := []struct {
		name string
		ctx  context.Context
		want *Context
	}{
		{
			want: &Context{
				Request: newRequestWithContext(context.Background()),
				Client:  &http.Client{Transport: http.DefaultTransport},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWithContext(context.Background()); got == nil ||
				!reflect.DeepEqual(got.Client, tt.want.Client) ||
				!reflect.DeepEqual(got.Request, tt.want.Request) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContext_GetClient(t *testing.T) {
	type fields struct {
		Client   *http.Client
		Request  *http.Request
		Response *http.Response
	}
	tests := []struct {
		name   string
		fields fields
		want   *http.Client
	}{
		{
			fields: fields{
				Client: &http.Client{Transport: http.DefaultTransport},
			},
			want: &http.Client{Transport: http.DefaultTransport},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Context{
				Client:   tt.fields.Client,
				Request:  tt.fields.Request,
				Response: tt.fields.Response,
			}
			if got := c.GetClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Context.GetClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContext_GetRequest(t *testing.T) {
	type fields struct {
		Client   *http.Client
		Request  *http.Request
		Response *http.Response
	}
	tests := []struct {
		name   string
		fields fields
		want   *http.Request
	}{
		{
			fields: fields{
				Request: newRequest(),
			},
			want: newRequest(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Context{
				Client:   tt.fields.Client,
				Request:  tt.fields.Request,
				Response: tt.fields.Response,
			}
			if got := c.GetRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Context.GetRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContext_GetResponse(t *testing.T) {
	type fields struct {
		Client   *http.Client
		Request  *http.Request
		Response *http.Response
	}
	tests := []struct {
		name   string
		fields fields
		want   *http.Response
	}{
		{
			fields: fields{
				Response: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Context{
				Client:   tt.fields.Client,
				Request:  tt.fields.Request,
				Response: tt.fields.Response,
			}
			if got := c.GetResponse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Context.GetResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
