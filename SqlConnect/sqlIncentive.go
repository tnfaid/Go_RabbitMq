package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/incentiverules")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	query := "SELECT * FROM fulfillment_error_code where code = ?"
	rows, err := db.Query(query, "30RV-0008")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var code string
		var description string
		var retry int
		var last_event string
		if err := rows.Scan(&code, &description, &retry, &last_event); err != nil {
			panic(err.Error())
		}
		fmt.Printf("Code dari esb: %s, Description: %s, retry: %d, last_event: %s \n", code, description, retry, last_event)
	}

	db.Close()
}
