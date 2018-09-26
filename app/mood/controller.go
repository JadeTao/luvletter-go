package mood

import (
	"luvletter/app/user"
	"luvletter/custom"
	"net/http"

	"github.com/labstack/echo"
)

// Save ...
func Save(c echo.Context) error {
	var (
		m     Mood
		trace user.TrackAction
		err   error
	)

	if err = c.Bind(&m); err != nil {
		return custom.NewHTTPError(http.StatusBadRequest, "error occurred when binding parameters", err.Error())
	}

	trace, err = user.TrackUserAction(m.Account, "create mood", "")
	m.CreateTime = trace.Time

	err = SaveMood(&m)

	if err != nil {
		return custom.NewHTTPError(
			http.StatusInternalServerError,
			"error occurred when saving tag",
			err.Error(),
		)
	}

	return c.JSON(http.StatusOK, m)
}
