package tag

import (
	"luvletter/conf"
	"strings"
	"time"
)

// SaveTag ...
func SaveTag(t *Tag) error {
	db := conf.GetDB()
	stmt, err := db.Prepare(`INSERT INTO tag (account, name, count, create_time) VALUES (?, ?, ?, ?)`)
	res, err := stmt.Exec(t.Account, t.Name, 1, t.CreateTime)
	defer stmt.Close()
	id, err := res.LastInsertId()
	t.ID = id
	return err
}

// AddCount ...
func AddCount(name string) error {
	db := conf.GetDB()
	stmt, err := db.Prepare(`UPDATE tag SET count=count+1, last_used_time=? WHERE name=?`)
	_, err = stmt.Exec(time.Now().Format("2006-01-02 15:04:05"), name)
	defer stmt.Close()
	return err
}

// AddCountInBatch ...
func AddCountInBatch(names []string) error {
	db := conf.GetDB()
	stmt, err := db.Prepare(`UPDATE tag SET count=count+1, last_used_time=? WHERE name IN (?)`)
	_, err = stmt.Exec(time.Now().Format("2006-01-02 15:04:05"), strings.Join(names, ","))
	defer stmt.Close()
	return err
}
