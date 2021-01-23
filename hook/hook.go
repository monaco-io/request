package hook

import "net/http"

// Hook ...
type Hook interface {
	BeforeRequest(*http.Request) (*http.Response, error)

	AfterRequest(*http.Request, *http.Response, error) (*http.Response, error)
}

func applyBeforeReqHooks(req *http.Request, hooks []Hook) (
	resp *http.Response, err error) {
	for _, hook := range hooks {
		resp, err = hook.BeforeRequest(req)
		if resp != nil || err != nil {
			return
		}
	}
	return
}

func applyAfterReqHooks(
	req *http.Request, resp *http.Response, err error, hooks []Hook) (
	newResp *http.Response, newErr error) {
	for _, hook := range hooks {
		newResp, newErr = hook.AfterRequest(req, resp, err)
		if newResp != nil || newErr != nil {
			return
		}
	}
	return
}
