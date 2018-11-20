package main

import (
	"luvletter/conf"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"luvletter/app/user"
	"luvletter/custom"
	"luvletter/router"
)

func main() {
	// Echo instance
	e := echo.New()

	e.HTTPErrorHandler = custom.HTTPErrorHandler
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Skipper:    router.Skip,
		SigningKey: []byte("secret"),
		Claims:     &user.JwtCustomClaims{},
		ErrorHandler: func(err error) error {
			return custom.NewHTTPError(http.StatusUnauthorized, "missing or invalid token", "请重新登录")
		},
	}))

	for path, handler := range router.GETRouters {
		e.GET(path, handler)
	}

	for path, handler := range router.POSTRouters {
		e.POST(path, handler)
	}
	port := ":4000"
	if conf.Conf.Mode == "production" {
		port = ":80"
	}
	e.File("/", "public/index.html")
	e.File("/favicon.ico", "public/favicon.ico")
	e.Logger.Fatal(e.Start(port))
}
