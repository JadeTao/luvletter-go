package letter

import (
	"luvletter/app/mood"
	"luvletter/app/tag"
	"luvletter/app/user"
	"luvletter/custom"
	"net/http"

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

	if err = c.Bind(&l); err != nil {
		return custom.NewHTTPError(http.StatusBadRequest, "error occurred when binding parameters", err.Error())
	}

	trace, err = user.TrackUserAction(l.Account, "create letter", "")
	l.CreateTime = trace.Time

	err = SaveLetter(&l)
	if err != nil {
		return custom.NewHTTPError(
			http.StatusInternalServerError,
			"error occurred when saving letter",
			err.Error(),
		)
	}

	// mood、tag计数
	_ = tag.AddCountInBatch(l.Tags)
	_ = mood.AddCount(l.Mood)

	return c.JSON(http.StatusOK, l)
}
