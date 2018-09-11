package custom

// HTTPError ...
type HTTPError struct {
	code    int
	Key     string `json:"error"`
	Message string `json:"message"`
}

// NewHTTPError ...
func NewHTTPError(code int, key string, msg string) *HTTPError {
	return &HTTPError{
		code:    code,
		Key:     key,
		Message: msg,
	}
}

// Error makes it compatible with `error` interface.
func (e *HTTPError) Error() string {
	return e.Key + ": " + e.Message
}
