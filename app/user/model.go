package user

import (
	"luvletter/util"

	jwt "github.com/dgrijalva/jwt-go"
)

// User struc
type User struct {
	ID         int16 `json:"id"`
	Avator     util.NullString
	Account    string
	Name       string
	Nickname   string
	Password   string
	CreateTime int
	UpdateTime int
}

// JwtCustomClaims ...
type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

// SimpleUser ...
type SimpleUser struct {
	Account  string
	Password string
}

// NewUser ...
type NewUser struct {
	Account  string
	NickName string
	Password string
}

// ResUser ...
type ResUser struct {
	Avator   util.NullString `json:"avator"`
	Account  string          `json:"account"`
	Nickname string          `json:"nickname"`
	Token    string          `json:"token"`
}

// TrackAction ...
type TrackAction struct {
	ID      int16
	Account string
	Time    string
	Action  string
	Extra   util.NullString
}
