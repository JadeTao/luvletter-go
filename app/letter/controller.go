package letter

import (
	"luvletter/custom"
	"net/http"

	"github.com/labstack/echo"
)

// Save ...
func Save(c echo.Context) error {
	var (
		l   Letter
		err error
	)
	if err = c.Bind(&l); err != nil {
		return custom.NewHTTPError(http.StatusBadRequest, "error occurred when binding parameters", err.Error())
	}
	SaveLetter(l)
	return err
}
