package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"luvletter/conf"
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

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// // Route => handler
	// e.GET("/", func(c echo.Context) error {
	// 	db, err := sql.Open("mysql", "root:12150112@tcp(127.0.0.1:3306)/luvletter?charset=utf8")

	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		response := Excuse{ID: "", Error: "true", Quote: ""}
	// 		return c.JSON(http.StatusInternalServerError, response)
	// 	}
	// 	defer db.Close()

	// 	var quote string
	// 	var id string
	// 	err = db.QueryRow("SELECT id, quote FROM excuses ORDER BY RAND() LIMIT 1").Scan(&id, &quote)

	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	fmt.Println(quote)
	// 	response := Excuse{ID: id, Error: "false", Quote: quote}
	// 	return c.JSON(http.StatusOK, response)
	// })

	for path, handler := range router.GETRouters {
		e.GET(path, handler)
	}

	for path, handler := range router.POSTRouters {
		e.POST(path, handler)
	}

	fmt.Print(conf.DBConfig)

	foo := map[string]int{
		"a": 1,
		"b": 2,
	}

	const a = "a"
	fmt.Println(foo[a])
	e.Logger.Fatal(e.Start(":4000"))
}
