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
		return custom.BadRequestError("binding parameters error", err)
	}

	trace, err = user.TrackUserAction(t.Account, "create tag", "")
	if err != nil {
		return custom.HTTPTrackError(err)
	}
	t.CreateTime = trace.Time

	err = SaveTag(&t)

	if err != nil {
		return custom.InternalServerError("saving tag error", err)
	}

	return c.JSON(http.StatusOK, t)
}
