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
	type user struct {
		Account  string
		NickName string
		Password string
	}
	var (
		u       user
		resUser ResUser
	)

	if err := c.Bind(&u); err != nil {
		return custom.NewHTTPError(http.StatusBadRequest, "error occurred when binding parameters", err.Error())
	}
	db, err := sql.Open("mysql", conf.DBConfig)
	stmt, err := db.Prepare(`INSERT INTO user (account, nickname, password) VALUES (?, ?, ?)`)
	res, err := stmt.Exec(u.Account, u.NickName, u.Password)
	defer stmt.Close()
	_, err = res.LastInsertId()
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

	type user struct {
		Account  string
		Password string
	}

	var (
		u   user
		res ResUser
	)

	err := c.Bind(&u)
	if err != nil || u.Account == "" || u.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	db, err := sql.Open("mysql", conf.DBConfig)
	row := db.QueryRow(`SELECT avator,account,nickname FROM user WHERE account=?`, u.Account)
	err = row.Scan(&res.Avator, &res.Account, &res.Nickname)

	if err == nil && res.Account != "" {

		if res.Token, err = GenerateToken(res.Account, true); err != nil {
			return custom.NewHTTPError(
				http.StatusInternalServerError,
				"error occurred when generating token",
				err.Error(),
			)
		}

		return c.JSON(http.StatusOK, res)
	}

	return echo.ErrUnauthorized
}
