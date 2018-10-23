package user

import (
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
	)

	if err := c.Bind(&u); err != nil {
		return custom.NewHTTPError(http.StatusBadRequest, "error occurred when binding parameters", err.Error())
	}

	u.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	u.UpdateTime = u.CreateTime
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
	resUser.Avatar.Valid = false

	if _, err = TrackUserAction(resUser.Account, "register", ""); err != nil {
		return custom.HTTPTrackError(err)
	}
	return c.JSON(http.StatusOK, resUser)
}

// Login logic
func Login(c echo.Context) error {

	var (
		u   SimpleUser
		res ResUser
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
		res.Avatar = user.Avatar
		res.Nickname = user.Nickname

		if _, err = TrackUserAction(user.Account, "login", ""); err != nil {
			return custom.HTTPTrackError(err)
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
