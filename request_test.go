package request

import (
	"testing"
)

func TestClient_Send(t *testing.T) {
	c := Client{
		URL:    "http://localhost:8080/ping",
		Method: "POST",
		// BodyString: "A=A&B=2",
		BodyJSON: []byte(`{"A":"A"}`),
		Header:   map[string]string{"Content-Type": "application/x-www-form-urlencoded"},
	}
	t.Log(c.Send().String())
}
