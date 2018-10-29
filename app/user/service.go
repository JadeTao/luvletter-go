package user

import (
	"database/sql"
	"luvletter/conf"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateToken ...
func GenerateToken(account string, admin bool) (string, error) {
	// Set custom claims
	claims := &JwtCustomClaims{
		account, // Account
		admin,   // Admin
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	return token.SignedString([]byte("secret"))
}

// SaveUser ...
func SaveUser(u NewUser) error {
	db, err := sql.Open("mysql", conf.DBConfig)
	stmt, err := db.Prepare(`INSERT INTO user (account, nickname, password) VALUES (?, ?, ?)`)
	_, err = stmt.Exec(u.Account, u.NickName, u.Password)
	defer stmt.Close()
	return err
}

// GetUserByID ...
func GetUserByID(id int16) (User, error) {
	var res User
	db, err := sql.Open("mysql", conf.DBConfig)
	row := db.QueryRow(`SELECT id, avatar, account, nickname, password,create_time,update_time FROM user WHERE id=?`, id)
	err = row.Scan(&res.ID, &res.Avatar, &res.Account, &res.Nickname, &res.Password, &res.CreateTime, &res.UpdateTime)
	return res, err
}

// GetUserByAccount ...
func GetUserByAccount(account string) (User, error) {
	var res User
	db, err := sql.Open("mysql", conf.DBConfig)
	row := db.QueryRow(`SELECT id, avatar, account, nickname, password,create_time,update_time FROM user WHERE account=?`, account)
	err = row.Scan(&res.ID, &res.Avatar, &res.Account, &res.Nickname, &res.Password, &res.CreateTime, &res.UpdateTime)
	return res, err
}

// UpdateUser ...
func UpdateUser(u User) error {
	db, err := sql.Open("mysql", conf.DBConfig)
	stmt, err := db.Prepare(`UPDATE user SET avatar=?,nickname=?,password=?,update_time=? WHERE id=?`)
	_, err = stmt.Exec(u.Avatar, u.Nickname, u.Password, u.UpdateTime, u.ID)
	defer stmt.Close()
	return err
}

// TrackUserAction ...
func TrackUserAction(account string, action string, extra string) (TrackAction, error) {
	var (
		track TrackAction
	)
	track.Account = account
	track.Action = action
	track.Time = time.Now().Format("2006-01-02 15:04:05")
	track.Extra = ""

	db, err := sql.Open("mysql", conf.DBConfig)
	stmt, err := db.Prepare(`INSERT INTO trace (account, time, action, extra) VALUES (?, ?, ?, ?)`)
	if err != nil {
		return track, err
	}
	_, err = stmt.Exec(track.Account, track.Time, track.Action, track.Extra)
	defer stmt.Close()
	return track, err
}
