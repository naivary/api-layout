package httperror

type HTTPError struct {
    StatusCode int
}

func (h HTTPError) Error() string {
    return ""
}


