package tag

import (
	"luvletter/app/user"
	"luvletter/custom"
	"net/http"

	"github.com/labstack/echo"
)

// Save ...
func Save(c echo.Context) error {
	var (
		t     Tag
		trace user.TrackAction
		err   error
	)

	if err = c.Bind(&t); err != nil {
		return custom.NewHTTPError(http.StatusBadRequest, "error occurred when binding parameters", err.Error())
	}

	trace, err = user.TrackUserAction(t.Account, "create tag", "")
	t.CreateTime = trace.Time

	err = SaveTag(&t)

	if err != nil {
		return custom.NewHTTPError(
			http.StatusInternalServerError,
			"error occurred when saving tag",
			err.Error(),
		)
	}

	return c.JSON(http.StatusOK, t)
}
