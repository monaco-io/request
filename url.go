package request

import "net/url"

// EncodeURL add and encode parameters.
func EncodeURL(baseURL string, p map[string]string) (string, error) {
	URL, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}
	query := URL.Query()
	for k := range p {
		query.Set(k, p[k])
	}
	URL.RawQuery = query.Encode()
	return URL.String(), nil
}

func EncodeParams(p map[string]string) string {
	params := url.Values{}
	for k := range p {
		params.Set(k, p[k])
	}
	return params.Encode()
}
