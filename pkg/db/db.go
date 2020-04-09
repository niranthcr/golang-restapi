package db

import (
	"database/sql"
)

type DbService struct {
	Db *sql.DB
}

func GetDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:helloworld@tcp(127.0.0.1:3308)/testapp")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (d *DbService) insert(val1, val2 string) error {
	insert, err := d.Db.Query("INSERT INTO tags VALUES ('%d', '%w')", val1, val2)
	if err != nil {
		return err
	}
	defer insert.Close()
	return nil
}

func (d *DbService) GetTagData() (*[]map[string]string, error) {
	results, err := d.Db.Query("SELECT id, name from tags")
	if err != nil {
		return nil, err
	}
	var t = []map[string]string{}
	for results.Next() {
		var id string
		var name string
		err = results.Scan(&id, &name)
		if err != nil {
			return nil, err
		}
		var m = map[string]string{}
		m["id"] = id
		m["name"] = name
		t = append(t, m)
	}
	return &t, nil
}
