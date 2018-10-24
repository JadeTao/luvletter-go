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
		return custom.BadRequestError("binding parameters error", err)
	}

	trace, err = user.TrackUserAction(m.Account, "create mood", "")
	if err != nil {
		return custom.HTTPTrackError(err)
	}
	m.CreateTime = trace.Time

	err = SaveMood(&m)

	if err != nil {
		return custom.InternalServerError("saving mood error", err)
	}

	return c.JSON(http.StatusOK, m)
}
