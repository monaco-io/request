package benchmark

import (
	"net/http"
	"testing"
)

func TestWork(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://www.baidu.com", nil)
	w := Work{
		Request: req,
		N:       100,
		C:       10,
	}
	w.Init()
	w.Run()
}
