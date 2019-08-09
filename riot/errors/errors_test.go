package errors

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
	re := NewRequestError(CreateResponse(501))
	if !strings.Contains(re.String(), `failed to respond`) {
		t.Errorf("Did not contain expected message")
	}

	re = NewRequestError(CreateResponse(700))
	if !strings.Contains(re.String(), `Server`) {
		t.Errorf("Did not contain expected message")
	}
}

func TestGetMessages(t *testing.T) {
	for code, msg := range errMessages {
		re := NewRequestError(CreateResponse(code))
		if !strings.Contains(re.String(), msg) {
			t.Errorf("Did not contain expected message")
		}
	}

}
