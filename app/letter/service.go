package letter

import (
	"database/sql"
	"luvletter/conf"
	"strings"
)

// SaveLetter ...
func SaveLetter(l Letter) error {
	db, err := sql.Open("mysql", conf.DBConfig)
	stmt, err := db.Prepare(`INSERT INTO letter (account, content, mood, nickname, tag,create_time) VALUES (?, ?, ?, ?, ?, ?)`)
	res, err := stmt.Exec(l.Account, l.Content, l.Mood, l.Nickname, strings.Join(l.Tag, ","), l.CreateTime)
	defer stmt.Close()
	_, err = res.LastInsertId()
	return err
}
