package main

import (
	"database/sql"
	"log"

	"github.com/GianlucaTurra/todoCli/cmd"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./todo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStatement := `
	create table if not exists todos (
		id integer not null primary key autoincrement,
		name text not null,
		description text,
		dueDate integer default 0,
		completed integer not null default 0
	)
	`

	_, err = db.Exec(sqlStatement)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStatement)
		return
	}

	cmd.Execute()
}
