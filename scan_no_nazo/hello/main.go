package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

type Student struct {
	ID     int
	Active bool
	Name   string
	Grade  int
	Score  float64
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
		s := Student{}
		if err := rows.StructScan(&s); err != nil {
			log.Fatal(err)
		}
		log.Println(s)
	}
	/*
	   ===== 実行結果 =====
	   2019/08/07 18:08:37 {1 true John Doe 65535 0.009876}
	   2019/08/07 18:08:37 sql: Scan error on column index 1, name "active": sql/driver: couldn't convert <nil> (<nil>) into type bool

	   students.active はNULLを許可する列で、そのScanに失敗する。
	*/
}
