package user

import (
	"database/sql"
	"luvletter/conf"
	"luvletter/custom"
	"net/http"

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

	err := AddUser(u)

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
	return c.JSON(http.StatusOK, resUser)
}

// Login logic
func Login(c echo.Context) error {

	var (
		u   SimpleUser
		res ResUser
	)

	err := c.Bind(&u)
	if err != nil || u.Account == "" || u.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	db, err := sql.Open("mysql", conf.DBConfig)
	row := db.QueryRow(`SELECT avator,account,nickname FROM user WHERE account=?`, u.Account)
	err = row.Scan(&res.Avator, &res.Account, &res.Nickname)

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

		return c.JSON(http.StatusOK, res)
	}

	return echo.ErrUnauthorized
}
