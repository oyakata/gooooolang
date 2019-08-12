package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type Student struct {
	ID     int
	Active bool
	Name   string
	Grade  int
	Score  float64
}

func main() {
	cfg := mysql.NewConfig()
	cfg.User = "root"
	cfg.Net = "tcp"
	cfg.Addr = "mysql:3306"
	cfg.DBName = "scan_no_nazo"

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT id, active, name, grade, score FROM students`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		s := Student{}
		if err := rows.Scan(&s.ID, &s.Active, &s.Name, &s.Grade, &s.Score); err != nil {
			log.Fatal(err)
		}
		log.Println(s)
	}
	/*
	        ===== 実行結果 =====
	2019/08/12 08:45:26 {1 true John Doe 65535 0.009876}
	2019/08/12 08:45:26 sql: Scan error on column index 1, name "active": sql/driver: couldn't convert <nil> (<nil>) into type bool

	        students.active はNULLを許可する列で、そのScanに失敗する。
	*/
}
