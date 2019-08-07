package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

type NullableStudent struct {
	ID     int
	Active sql.NullBool
	Name   sql.NullString
	Grade  sql.NullInt64
	Score  sql.NullFloat64
}

func main() {
	conn, err := sql.Open("mysql", "root:@tcp(mysql:3306)/scan_no_nazo")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	db := sqlx.NewDb(conn, "mysql")
	rows, err := db.Queryx(`SELECT id, active, name, grade, score FROM students`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		s := NullableStudent{}
		if err := rows.StructScan(&s); err != nil {
			log.Fatal(err)
		}
		log.Println(s)
	}
	/*
	   ===== 実行結果 =====

	*/
}
