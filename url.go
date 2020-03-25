package request

import "net/url"

// EncodeURL add and encoded parameters.
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
