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
	if err != nil {
		return res, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			l    Letter
			tags string
		)
		rows.Scan(&l.ID, &l.Account, &l.Nickname, &l.Content, &l.CreateTime, &l.Mood, &l.Tags)
		l.Tags = strings.FieldsFunc(tags, util.Split(','))
		res = append(res, l)
	}
	return res, err
}

// FindPage ...
func FindPage(position int64, offset int64) ([]Letter, error) {
	res := make([]Letter, 0)
	db, err := sql.Open("mysql", conf.DBConfig)
	rows, err := db.Query(`SELECT id, account, nickname, content, create_time, mood, tags FROM letter LIMIT ?,?`, (position-1)*offset, offset)
	defer rows.Close()
	for rows.Next() {
		var (
			l    Letter
			tags string
		)
		rows.Scan(&l.ID, &l.Account, &l.Nickname, &l.Content, &l.CreateTime, &l.Mood, &tags)
		l.Tags = strings.FieldsFunc(tags, util.Split(','))
		res = append(res, l)
	}
	return res, err
}

// FindNumber ...
func FindNumber() (int64, error) {
	var length int64
	db, err := sql.Open("mysql", conf.DBConfig)
	row := db.QueryRow(`SELECT COUNT(id) FROM letter`)

	row.Scan(&length)
	return length, err
}
