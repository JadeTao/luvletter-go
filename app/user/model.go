package user

import (
	"luvletter/util"

	jwt "github.com/dgrijalva/jwt-go"
)

// User struc
type User struct {
	ID            int64 `json:"id"`
	Avator        util.NullString
	Account       string
	Name          string
	Nickname      string
	Password      string
	CreateTime    string
	UpdateTime    string
	LastLoginTime string
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

// ResUser ...
type ResUser struct {
	Avator   util.NullString `json:"avator"`
	Account  string          `json:"account"`
	Nickname string          `json:"nickname"`
}
