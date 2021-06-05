package main

import (
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"log"
)

func selectData(db *sql.DB) {
	var id int
	var name string
	var age int
	rows, err := db.Query(`SELECT * From user where id = 1`)
	if err != nil {
		log.Fatalf("insert data error: %v\n", err)
		return
	}
	for rows.Next() {
		rows.Scan(&id, &age, &name)
		if err != nil {
			if err == sql.ErrNoRows{
				errors.Wrap(err, "no such data")
			}
		}
		log.Printf("get data, id: %d, name: %s, age: %d", id, name, age)
	}

}