package router

import (
	"luvletter/app/mood"
	"luvletter/app/tag"

	"github.com/labstack/echo"

	"luvletter/app/letter"
	"luvletter/app/user"
)

// GETRouters RouterConfig for GET.
var GETRouters = map[string]echo.HandlerFunc{
	"/letter": letter.GetAll,
}

// POSTRouters RouterConfig for POST.
var POSTRouters = map[string]echo.HandlerFunc{
	"/login":    user.Login,
	"/register": user.Register,
	"/letter":   letter.Save,
	"/tag":      tag.Save,
	"/mood":     mood.Save,
}
