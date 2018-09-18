package user

import (
	"fmt"
	"luvletter/custom"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// Register 注册
func Register(c echo.Context) error {

	var (
		u       NewUser
		resUser ResUser
		trace   TrackAction
	)

	if err := c.Bind(&u); err != nil {
		return custom.NewHTTPError(http.StatusBadRequest, "error occurred when binding parameters", err.Error())
	}

	err := SaveUser(u)

	if err != nil {
		return custom.NewHTTPError(
			http.StatusInternalServerError,
			"error occurred when processing database",
			err.Error(),
		)
	}

	if resUser.Token, err = GenerateToken(resUser.Account, true); err != nil {
		return custom.NewHTTPError(
			http.StatusInternalServerError,
			"error occurred when generating token",
			err.Error(),
		)
	}

	resUser.Account = u.Account
	resUser.Nickname = u.NickName
	resUser.Avator.Valid = false

	trace.Account = u.Account
	trace.Time = time.Now().Format("2006-01-02 15:04:05")
	trace.Action = "register"
	trace.Extra.Valid = false
	if err = TrackUserAction(trace); err != nil {

	}
	return c.JSON(http.StatusOK, resUser)
}

// Login logic
func Login(c echo.Context) error {

	var (
		u     SimpleUser
		res   ResUser
		trace TrackAction
	)

	if err := c.Bind(&u); err != nil {
		return custom.NewHTTPError(http.StatusBadRequest, "error occurred when binding parameters", err.Error())
	}
	if u.Account == "" || u.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	user, err := GetUserByAccount(u.Account)

	if err == nil && user.Account != "" {
		if res.Token, err = GenerateToken(res.Account, true); err != nil {
			return custom.NewHTTPError(
				http.StatusInternalServerError,
				"error occurred when generating token",
				err.Error(),
			)
		}
		res.Account = user.Account
		res.Avator = user.Avator
		res.Nickname = user.Nickname

		trace.Action = "login"
		trace.Account = user.Account
		trace.Time = time.Now().Format("2006-01-02 15:04:05")
		trace.Extra.Valid = false
		if err = TrackUserAction(trace); err != nil {
		}
		return c.JSON(http.StatusOK, res)
	} else if err != nil {
		return custom.NewHTTPError(
			http.StatusInternalServerError,
			"error occurred when processing database",
			err.Error(),
		)
	}

	return echo.ErrUnauthorized
}
