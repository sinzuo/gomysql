package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	db, err := sqlx.Open("mysql", "root:jiangyibo@tcp(127.0.0.1:4306)/mytest?charset=utf8")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = db
}

func main() {
	tx, err := Db.Beginx()
	_, err = tx.Exec("insert into userinfo(username,password) values(?,?)", "Rose", "2223")
	if err != nil {
		tx.Rollback()
	}
	_, err = tx.Exec("insert into userinfo(username,password) values(?,?)", "Mick", 222)
	if err != nil {
		fmt.Println("exec sql error:", err)
		tx.Rollback()
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println("commit error")
	}

}
