package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

func main() {
	var err error
	const file string = "History.db"
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	defer db.Close()
	rows, err := db.Query("select name from sqlite_master where type = 'table'")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	type str struct {
		Name string
	}
	result := []str{}

	for rows.Next() {
		s := str{}
		err := rows.Scan(&s.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		result = append(result, s)
	}
	for _, s := range result {
		fmt.Println(s.Name)
	}
}
