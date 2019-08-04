package riot

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

var (
	errMsg []byte = []byte(`placeholder`)
	retry  string = `10`
)

func CreateResponse(code int) *http.Response {
	resp := &http.Response{}
	resp.Body = ioutil.NopCloser(bytes.NewReader(errMsg))
	resp.StatusCode = code
	resp.Header = http.Header{}
	resp.Header.Add(`Retry-After`, retry)

	return resp
}

func TestRequestErrorMsg(t *testing.T) {
	re := NewRequestError(CreateResponse(401))
	if !strings.Contains(re.String(), errMessages[401]) {
		t.Errorf("Did not contain expected message")
	}
}
