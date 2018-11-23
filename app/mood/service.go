package mood

import (
	"luvletter/conf"
	"time"
)

// SaveMood ...
func SaveMood(m *Mood) error {
	db := conf.GetDB()
	stmt, err := db.Prepare(`INSERT INTO mood (account, name, count, create_time) VALUES (?, ?, ?, ?)`)
	res, err := stmt.Exec(m.Account, m.Name, 0, m.CreateTime)
	defer stmt.Close()
	id, err := res.LastInsertId()
	m.ID = id
	return err
}

// AddCount ...
func AddCount(name string) error {
	db := conf.GetDB()
	stmt, err := db.Prepare(`UPDATE mood SET count=count+1, last_used_time=? WHERE name=?`)
	_, err = stmt.Exec(time.Now().Format("2006-01-02 15:04:05"), name)

	defer stmt.Close()
	return err
}
