package mood

import (
	"net/http"

	"github.com/JadeTao/luvletter-go/app/user"
	"github.com/JadeTao/luvletter-go/custom"

	jwt "github.com/dgrijalva/jwt-go"
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

// GetAll ...
func GetAll(c echo.Context) error {
	var (
		all []string
		err error
	)
	userInfo := c.Get("user").(*jwt.Token)
	claims := userInfo.Claims.(*user.JwtCustomClaims)
	account := claims.Account

	_, err = user.TrackUserAction(account, "get moods", "")
	if err != nil {
		return custom.HTTPTrackError(err)
	}
	if all, err = FindAll(); err != nil {
		return custom.BadRequestError("get all moods error", err)
	}
	return c.JSON(http.StatusOK, all)
}
