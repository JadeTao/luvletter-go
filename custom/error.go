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
