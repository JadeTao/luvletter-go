package router

import (
	"github.com/labstack/echo"
)

// Skip ...
func Skip(c echo.Context) bool {
	skippedPath := map[string]bool{
		"/login":    true,
		"/register": true,
	}
	if skippedPath[c.Path()] {
		return true
	}
	return false
}
