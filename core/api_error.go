package core

import (
	"fmt"
	"net/http"
)

// APIError is a lightweight wrapper around the standard error
// interface that preserves the status code from the RPC, if any.
type APIError struct {
	err error

	StatusCode int         `json:"-"`
	Header     http.Header `json:"-"`
}

// NewAPIError constructs a new API error.
func NewAPIError(statusCode int, header http.Header, err error) *APIError {
	return &APIError{
		err:        err,
		Header:     header,
		StatusCode: statusCode,
	}
}

// Unwrap returns the underlying error. This also makes the error compatible
// with errors.As and errors.Is.
func (a *APIError) Unwrap() error {
	if a == nil {
		return nil
	}
	return a.err
}

// Error returns the API error's message.
func (a *APIError) Error() string {
	if a == nil || (a.err == nil && a.StatusCode == 0) {
		return ""
	}
	if a.err == nil {
		return fmt.Sprintf("%d", a.StatusCode)
	}
	if a.StatusCode == 0 {
		return a.err.Error()
	}
	return fmt.Sprintf("%d: %s", a.StatusCode, a.err.Error())
}

// ResponseDecodeError is returned (wrapped in an *APIError that carries the
// HTTP status code and headers) when a successful response body cannot be
// decoded into the expected response type. Body holds the raw response body
// and Err the underlying decoding error.
type ResponseDecodeError struct {
	Target string
	Body   []byte
	Err    error
}

// Error returns a description of the decoding failure, including the start
// of the raw response body.
func (e *ResponseDecodeError) Error() string {
	body := e.Body
	suffix := ""
	if len(body) > 500 {
		body = body[:500]
		suffix = "..."
	}
	return fmt.Sprintf("failed to decode response body into %s: %v; body: %s%s", e.Target, e.Err, body, suffix)
}

// Unwrap returns the underlying decoding error.
func (e *ResponseDecodeError) Unwrap() error {
	return e.Err
}
