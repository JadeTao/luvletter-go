package upload

import (
	"io"
	"net/http"
	"os"
	"path"

	"github.com/labstack/echo"
)

// Avator ...
func Avator(c echo.Context) error {
	//-----------
	// Read file
	//-----------
	account := c.FormValue("account")
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("./assets/avators/" + account + path.Ext(file.Filename))
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}
