package main

import (
	"database/sql"
	"log"

	mysql "github.com/go-sql-driver/mysql"
)

type Student struct {
	ID     int
	Active bool
	Name   string
	Grade  int
	Score  float64
}

type NullableStudent struct {
	ID     int
	Active sql.NullBool
	Name   sql.NullString
	Grade  sql.NullInt64
	Score  sql.NullFloat64
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
		s := NullableStudent{}
		if err := rows.Scan(&s.ID, &s.Active, &s.Name, &s.Grade, &s.Score); err != nil {
			log.Fatal(err)
		}

		st := Student{
			ID:     s.ID,
			Active: s.Active.Bool,
			Name:   s.Name.String,
			Grade:  int(s.Grade.Int64),
			Score:  s.Score.Float64,
		}
		log.Println(st)
		/*
		   2019/08/12 09:17:52 {1 true John Doe 65535 0.009876}
		   2019/08/12 09:17:52 {2 false  0 0}
		   2019/08/12 09:17:52 {3 false Akira Toriyama 0 100}
		*/
	}
}
