package router

import (
	"github.com/labstack/echo"
)

// Skip ...
func Skip(c echo.Context) bool {
	skippedPath := prefixSkipMapper(map[string]bool{
		"/login":    true,
		"/register": true,
	}, Prefix)
	if skippedPath[c.Path()] {
		return true
	}
	return false
}

func prefixSkipMapper(router map[string]bool, prefix string) map[string]bool {
	var withPrefixRouter = make(map[string]bool)
	for key, value := range router {
		withPrefixRouter[prefix+key] = value
	}
	return withPrefixRouter
}