package tag

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

// GetAll ...
func GetAll(c echo.Context) error {
	var (
		all []string
		err error
	)
	userInfo := c.Get("user").(*jwt.Token)
	claims := userInfo.Claims.(*user.JwtCustomClaims)
	account := claims.Account

	_, err = user.TrackUserAction(account, "get tags", "")
	if err != nil {
		return custom.HTTPTrackError(err)
	}
	if all, err = FindAll(); err != nil {
		return custom.BadRequestError("get all letters error", err)
	}
	return c.JSON(http.StatusOK, all)
}
