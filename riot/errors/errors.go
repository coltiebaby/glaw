package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"io/ioutil"
)

type RequestError struct {
	Message     string
	RiotMessage string
	StatusCode  int
	Wait        int
}

func NewErrorFromString(msg string) *RequestError {
	return &RequestError{Message: msg}
}

func NewRequestError(resp *http.Response) *RequestError {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return NewErrorFromString(fmt.Sprintf("Could not read error from body: %s", err))
	}

	code := resp.StatusCode

	var retry int
	if r := resp.Header.Get(`Retry-After`); r != "" {
		retry, _ = strconv.Atoi(r)
	}

	return &RequestError{
		Message:     getMessage(code),
		RiotMessage: string(body[:]),
		StatusCode:  code,
		Wait:        retry,
	}

}

func (re *RequestError) Error() string {
	builder := &strings.Builder{}
	json.NewEncoder(builder).Encode(re)
	return builder.String()
}

func (re *RequestError) String() string {
	return re.Error()
}

func getMessage(code int) string {
	var (
		msg string
		ok  bool
	)

	if msg, ok = errMessages[code]; ok {
		return msg
	}

	switch {
	case code >= 500 && code < 600:
		msg = "Riot awknowledged us but failed to respond"
	default:
		msg = "Issue with the Riot Server"
	}

	return msg
}

var errMessages = map[int]string{
	http.StatusBadRequest:           `Bad Request. Make sure your url and parameters are correct`,
	http.StatusUnauthorized:         `Unauthorized. Confirm your API key`,
	http.StatusForbidden:            `Forbidden. Incorrect Path or API key`,
	http.StatusNotFound:             `Nah-ah-ah. Didn't say the magic words`,
	http.StatusUnsupportedMediaType: `Unsupported Media Type`,
	http.StatusTooManyRequests:      `TooManyReqeusts: Please try again in %d`,
}
