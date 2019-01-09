package user

import (
	"errors"
	"net/http"
	"time"

	"github.com/JadeTao/luvletter-go/custom"

	"github.com/labstack/echo"
)

// Register 注册
func Register(c echo.Context) error {

	var (
		u       NewUser
		resUser ResUser
	)

	if err := c.Bind(&u); err != nil {
		return custom.BadRequestError("binding parameters error", err)
	}

	u.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	u.UpdateTime = u.CreateTime
	err := SaveUser(u)

	if err != nil {
		return custom.InternalServerError("processing database error", err)
	}

	if resUser.Token, err = GenerateToken(resUser.Account, true); err != nil {
		return custom.InternalServerError("generating token error", err)

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
		return custom.BadRequestError("binding parameters error", err)
	}
	if u.Account == "" || u.Password == "" {
		return custom.BadRequestError("empty account or password", errors.New("empty account or password"))
	}

	user, err := GetUserByAccount(u.Account)

	if user.Password != u.Password {
		return custom.NewHTTPError(http.StatusForbidden, "wrong user infomation", "请验证您的账户和密码")
	}

	if err == nil && user.Account != "" {
		if res.Token, err = GenerateToken(user.Account, true); err != nil {
			return custom.InternalServerError("generating token error", err)
		}
		res.Account = user.Account
		res.Avatar = user.Avatar
		res.Nickname = user.Nickname

		if _, err = TrackUserAction(user.Account, "login", ""); err != nil {
			return custom.HTTPTrackError(err)
		}
		UpdateLastLoginTime(u.Account)
		return c.JSON(http.StatusOK, res)
	} else if err != nil {
		return custom.InternalServerError("processing database error", err)
	}

	return echo.ErrUnauthorized
}
