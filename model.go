package request

type Client struct {
	URL    string
	Method string
	Header map[string]string
	Params map[string]string
	Body   []byte
}
