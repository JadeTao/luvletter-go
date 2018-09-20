package letter

import (
	"luvletter/app/user"
	"luvletter/custom"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// GetAll ...
func GetAll(c echo.Context) error {
	var (
		all []Letter
		err error
	)
	if all, err = FindAll(); err != nil {
		return custom.NewHTTPError(http.StatusBadRequest, "error occurred when querying all letters", err.Error())
	}
	return c.JSON(http.StatusOK, all)
}

// Save ...
func Save(c echo.Context) error {
	var (
		l     Letter
		trace user.TrackAction
		err   error
	)
	trace.Account = l.Account
	trace.Action = "save letter"
	trace.Time = time.Now().Format("2006-01-02 15:04:05")

	l.CreateTime = trace.Time
	if err = c.Bind(&l); err != nil {
		return custom.NewHTTPError(http.StatusBadRequest, "error occurred when binding parameters", err.Error())
	}
	if err = SaveLetter(&l); err != nil {
		return custom.NewHTTPError(
			http.StatusInternalServerError,
			"error occurred when saving letter",
			err.Error(),
		)
	}

	return c.JSON(http.StatusOK, l)
}
