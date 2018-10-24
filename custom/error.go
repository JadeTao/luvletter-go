package custom

import (
	"net/http"
)

// HTTPTrackError ...
func HTTPTrackError(err error) *HTTPError {
	return NewHTTPError(
		http.StatusInternalServerError,
		"database error when tracking",
		err.Error(),
	)
}

// InternalServerError ...
func InternalServerError(msg string, err error) *HTTPError {
	return NewHTTPError(
		http.StatusInternalServerError,
		msg,
		err.Error(),
	)
}

// BadRequestError ...
func BadRequestError(msg string, err error) *HTTPError {
	return NewHTTPError(
		http.StatusBadRequest,
		msg,
		err.Error(),
	)
}
