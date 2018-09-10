package router

import (
	"github.com/labstack/echo"

	"luvletter/app/user"
)

// GETRouters RouterConfig for GET.
var GETRouters = map[string]echo.HandlerFunc{
	"/register": user.Register,
}

// POSTRouters RouterConfig for POST.
var POSTRouters = map[string]echo.HandlerFunc{
	"/login": user.Login,
}
