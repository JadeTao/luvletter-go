package user

import (
	"database/sql"
)

// User struc
type User struct {
	ID int `json:"id"`
	avator sql.NullString 
	username string
	name string
	nickname string
	password string
	email string
	roleID string
	createTime string
	updateTime string
}
