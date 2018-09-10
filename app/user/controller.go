package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"luvletter/conf"
	"luvletter/util"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// Excuse error message
type Excuse struct {
	Error string `json:"error"`
	ID    string `json:"id"`
	Quote string `json:"quote"`
}

// Register 注册
func Register(c echo.Context) error {
	db, err := sql.Open("mysql", conf.DBConfig)
	util.Check(err)

	stmt, err := db.Prepare(`INSERT INTO user (avator, username, name, nickname, password, email, role_id) VALUES (?, ?, ?, ?, ?, ?, ?)`)
	util.Check(err)

	res, err := stmt.Exec(nil, "test", "foo", "jader", "112681", "test@go.com", "1")
	util.Check(err)

	id, err := res.LastInsertId()
	util.Check(err)

	fmt.Println(id)
	stmt.Close()

	return nil
}

// Login logic
// func Login(c echo.Context) error {
// 	// username := c.FormValue("username")
// 	// password := c.FormValue("password")
// 	db, err := sql.Open("mysql", conf.DBConfig)
// 	util.Check(err)

// 	stmt, err := db.Prepare(`SELECT * FROM user`)
// 	util.Check(err)
// 	defer stmt.Close()

// 	rows, err := stmt.Query()
// 	util.Check(err)

// 	defer rows.Close()

// 	var (
// 		ret User
// 	)

// 	const timeFormat = "2010-12-15 12:12:12"

// 	for rows.Next() {
// 		err = rows.Scan(&ret.ID, &ret.Avator, &ret.Username, &ret.Name, &ret.Nickname, &ret.Password, &ret.Email, &ret.RoleID, &ret.CreateTime, &ret.UpdateTime, &ret.LastLoginTime)
// 		util.Check(err)
// 	}

// 	return c.JSON(http.StatusOK, ret)
// }

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
	util.Check(err)

	db, err := sql.Open("mysql", conf.DBConfig)
	util.Check(err)

	fmt.Println(u)
	row := db.QueryRow(`SELECT avator,account,nickname FROM user WHERE account=?`, u.Account)
	util.Check(err)

	err = row.Scan(&res.Avator, &res.Account, &res.Nickname)
	util.Check(err)
	if res.Account != "" {
		return c.JSON(http.StatusOK, res)
	}
	jsonByte, err := json.Marshal(&res)
	fmt.Println(string(jsonByte))

	if u.Account == "jon" && u.Password == "snow" {

		// Set custom claims
		claims := &jwtCustomClaims{
			"Jon Snow",
			true,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		util.Check(err)

		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}
