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
	if err != nil {
		return err
	}
	res, err := stmt.Exec(t.Account, t.Name, 0, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return err
	}
	defer stmt.Close()
	id, err := res.LastInsertId()
	t.ID = id
	return err
}

// AddCount ...
func AddCount(name string) error {
	db := conf.GetDB()
	stmt, err := db.Prepare(`UPDATE tag SET count=count+1, last_used_time=? WHERE name=?`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(time.Now().Format("2006-01-02 15:04:05"), name)
	defer stmt.Close()
	return err
}

// AddCountInBatch ...
func AddCountInBatch(names []string) error {
	db := conf.GetDB()
	stmt, err := db.Prepare(`UPDATE tag SET count=count+1, last_used_time=? WHERE name IN (?)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(time.Now().Format("2006-01-02 15:04:05"), strings.Join(names, ","))
	defer stmt.Close()
	return err
}

// FindAll ...
func FindAll() ([]string, error) {
	res := make([]string, 0)
	db := conf.GetDB()
	rows, err := db.Query(`SELECT name FROM tag`)
	if err != nil {
		return res, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			l string
		)
		rows.Scan(&l)
		res = append(res, l)
	}
	return res, err
}
