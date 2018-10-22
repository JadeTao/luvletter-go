package main

import (
	"luvletter/conf"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"luvletter/app/user"
	"luvletter/custom"
	"luvletter/router"
)

// Excuse error
type Excuse struct {
	Error string `json:"error"`
	ID    string `json:"id"`
	Quote string `json:"quote"`
}

func main() {
	// Echo instance
	e := echo.New()

	e.HTTPErrorHandler = custom.HTTPErrorHandler
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Skipper:    router.Skip,
		SigningKey: []byte("secret"),
		Claims:     &user.JwtCustomClaims{},
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
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
	e.Logger.Fatal(e.Start(port))
}
