package httperror

import "net/http"

func New(err error) HTTPError {
	return HTTPError{
        Err: err,
        StatusCode: http.StatusInternalServerError,
    }
}

func WithStatusCode(err error, statusCode int) HTTPError {
    return HTTPError{
        Err: err,
        StatusCode: statusCode,
    }
}

type HTTPError struct {
	Err        error
	Msg        string
	StatusCode int
}

func (h HTTPError) Error() string {
	return ""
}

func (h HTTPError) Unwrap() error {
    return h.Err
}
