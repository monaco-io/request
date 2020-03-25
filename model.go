package request

import "time"

const (
	ApplicationJSON               ContentType = "application/json"
	ApplicationXWwwFormURLEncoded ContentType = "application/x-www-form-urlencoded"
	MultipartFormData             ContentType = "multipart/form-data"
)

type ContentType string

type Client struct {
	URL         string
	Method      string
	Header      map[string]string
	Params      map[string]string
	Body        []byte
	Auth        Auth
	Timeout     time.Duration // second
	ContentType ContentType
}

type Auth struct {
	Username string
	Password string
}
