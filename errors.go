package glaw

import (
	"fmt"
	"net/http"
)

type RequestError struct {
	Code  int
	Retry string
}

func (re *RequestError) Error() string {
	if re.Retry != `` {
		return fmt.Sprintf("api limit reached: retry %s", re.Retry)
	}

	return fmt.Sprintf("%d: %s", re.Code, http.StatusText(re.Code))

}

func (re *RequestError) String() string {
	return re.Error()
}

func NewRequestError(resp *http.Response) *RequestError {
	return &RequestError{
		Code:  resp.StatusCode,
		Retry: resp.Header.Get(`Retry-After`),
	}
}
