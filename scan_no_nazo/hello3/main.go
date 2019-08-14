package main

import (
	"database/sql"
	"encoding/json"
	"log"

	mysql "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Student struct {
	ID     int
	Active *bool
	Name   *string
	Grade  *int
	Score  *float64
}

func main() {
	cfg := mysql.NewConfig()
	cfg.User = "root"
	cfg.Net = "tcp"
	cfg.Addr = "mysql:3306"
	cfg.DBName = "scan_no_nazo"

	conn, err := sql.Open("mysql", cfg.FormatDSN())
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
		s := Student{}
		if err := rows.StructScan(&s); err != nil {
			log.Fatal(err)
		}

		b, err := json.Marshal(&s)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(b))
	}
}
