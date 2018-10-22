package user

import (
	"luvletter/util"

	jwt "github.com/dgrijalva/jwt-go"
)

// User struc
type User struct {
	ID            int64           `json:"id"`
	Avatar        util.NullString `json:"avatar"`
	Account       string          `json:"account"`
	Name          string          `json:"name"`
	Nickname      string          `json:"nickname"`
	Password      string          `json:"password"`
	CreateTime    string          `json:"createTime"`
	UpdateTime    string          `json:"updateTime"`
	LastLoginTime string          `json:"lastLoginTime"`
}

// JwtCustomClaims ...
type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

// SimpleUser ...
type SimpleUser struct {
	Account       string `json:"account"`
	Password      string `json:"password"`
	LastLoginTime string `json:"lastLoginTime"`
}

// NewUser ...
type NewUser struct {
	Account    string `json:"account"`
	NickName   string `json:"nickname"`
	Password   string `json:"password"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

// ResUser ...
type ResUser struct {
	Avatar   util.NullString `json:"avatar"`
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
	Extra   string
}
