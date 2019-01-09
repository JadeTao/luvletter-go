package custom

import (
	"net/http"

	"github.com/JadeTao/luvletter-go/conf"

	"github.com/labstack/echo"
)

// HTTPErrorHandler customize echo's HTTP error handler.
func HTTPErrorHandler(err error, c echo.Context) {
	var (
		code = http.StatusInternalServerError
		key  = "ServerError"
		msg  string
	)

	if he, ok := err.(*HTTPError); ok {
		code = he.code
		key = he.Key
		msg = he.Message
	} else if ee, ok := err.(*echo.HTTPError); ok {
		code = ee.Code
		key = http.StatusText(code)
		msg = key
	} else if conf.Conf.Mode == "dev" {
		msg = err.Error()
	} else {
		msg = http.StatusText(code)
	}

	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			err := c.NoContent(code)
			if err != nil {
				c.Logger().Error(err)
			}
		} else {
			err := c.JSON(code, NewHTTPError(code, key, msg))
			if err != nil {
				c.Logger().Error(err)
			}
		}
	}
}
