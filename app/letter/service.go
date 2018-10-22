package letter

import (
	"database/sql"
	"luvletter/conf"
	"luvletter/util"
	"strings"
)

// SaveLetter ...
func SaveLetter(l *Letter) error {
	db, err := sql.Open("mysql", conf.DBConfig)
	stmt, err := db.Prepare(`INSERT INTO letter (account, content, mood, nickname, tags, create_time) VALUES (?, ?, ?, ?, ?, ?)`)
	res, err := stmt.Exec(l.Account, l.Content, l.Mood, l.Nickname, strings.Join(l.Tags, ","), l.CreateTime)
	defer stmt.Close()
	id, err := res.LastInsertId()
	l.ID = id
	return err
}

// FindAll ...
func FindAll() ([]Letter, error) {
	res := make([]Letter, 0)
	db, err := sql.Open("mysql", conf.DBConfig)
	rows, err := db.Query(`SELECT id, account, nickname, content, create_time, mood, tags FROM letter`)
	defer rows.Close()
	for rows.Next() {
		var (
			l   Letter
			tags string
		)
		rows.Scan(&l.ID, &l.Account, &l.Nickname, &l.Content, &l.CreateTime, &l.Mood, &tags)
		l.Tags = strings.FieldsFunc(tags, util.Split(','))
		res = append(res, l)
	}
	return res, err
}
