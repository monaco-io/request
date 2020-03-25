package request

type Client struct {
	URL    string
	Method string
	Params map[string]string
	Header map[string]string
	Body   []byte
}
