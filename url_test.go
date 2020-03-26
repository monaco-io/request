package request

import "testing"

func TestEncodeURL(t *testing.T) {
	type args struct {
		baseURL string
		p       map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "TestEncodeURL_0",
			args: args{
				baseURL: "https://google.com",
				p:       map[string]string{"a": "1", "b": "2"},
			},
			want:    "https://google.com?a=1&b=2",
			wantErr: false,
		},
		{
			name: "TestEncodeURL_1",
			args: args{
				baseURL: "https://google.com/path",
				p:       map[string]string{"a": "1", "b": "2"},
			},
			want:    "https://google.com/path?a=1&b=2",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncodeURL(tt.args.baseURL, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EncodeURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}
