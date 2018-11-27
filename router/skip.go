package router

import (
	"luvletter/util"

	"github.com/labstack/echo"
)

// Skip ...
func Skip(c echo.Context) bool {

	white := []string{APILogin, APIRegister, APIAvatar}

	whiteWithPrefix := prefixSkipMapper(white, Prefix)
	whiteWithPrefix = append(whiteWithPrefix, "/favicon.ico")
	for _, path := range whiteWithPrefix {
		if util.ComparePath(c.Path(), path) {
			return true
		}
	}
	return false
}

// func prefixSkipMapper(router map[string]bool, prefix string) map[string]bool {
// 	var withPrefixRouter = make(map[string]bool)
// 	for key, value := range router {
// 		withPrefixRouter[prefix+key] = value
// 	}
// 	return withPrefixRouter
// }

func prefixSkipMapper(router []string, prefix string) []string {
	length := len(router)
	routerWithPrefix := make([]string, length)
	for index, value := range router {
		routerWithPrefix[index] = prefix + value
	}

	return routerWithPrefix
}
