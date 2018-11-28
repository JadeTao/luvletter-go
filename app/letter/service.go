package letter

import (
	"luvletter/conf"
	"luvletter/util"
	"strings"
)

// SaveLetter ...
func SaveLetter(l *Letter) error {
	db := conf.GetDB()
	stmt, err := db.Prepare(`INSERT INTO letter (account, content, mood, nickname, tags, create_time) VALUES (?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(l.Account, l.Content, l.Mood, l.Nickname, strings.Join(l.Tags, ","), l.CreateTime)
	if err != nil {
		return err
	}
	defer stmt.Close()
	id, err := res.LastInsertId()
	l.ID = id
	return err
}

// FindAll ...
func FindAll() ([]Letter, error) {
	res := make([]Letter, 0)
	db := conf.GetDB()
	rows, err := db.Query(`SELECT id, account, nickname, content, create_time, mood, tags FROM letter ORDER BY id DESC`)
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
	db := conf.GetDB()
	rows, err := db.Query(`SELECT id, account, nickname, content, create_time, mood, tags FROM letter ORDER BY id DESC LIMIT ?,?`, position-1, offset)
	if err != nil {
		return nil, err
	}
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
	db := conf.GetDB()
	row := db.QueryRow(`SELECT COUNT(id) FROM letter`)

	err := row.Scan(&length)
	return length, err
}
