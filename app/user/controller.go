package user

import (
	"database/sql"
	"fmt"
	"luvletter/conf"
	"luvletter/util"

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
func Login(c echo.Context) error {
	// username := c.FormValue("username")
	// password := c.FormValue("password")
	db, err := sql.Open("mysql", conf.DBConfig)
	util.Check(err)

	stmt, err := db.Prepare(`SELECT * FROM user`)
	util.Check(err)
	defer stmt.Close()

	rows, err := stmt.Query()
	util.Check(err)

	defer rows.Close()
	var ret User

	for rows.Next() {
		err = rows.Scan(&ret.ID, &ret.avator, &ret.username, &ret.name, &ret.nickname, &ret.password, &ret.email, &ret.roleID, &ret.createTime, &ret.updateTime)
		util.Check(err)
	}
	fmt.Println(ret)

	return nil
}
