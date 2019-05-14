package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

type stu struct {
	Username   string `db:"username"`
	Password   string `db:"password"`
	Department string `db:"department"`
	Email      string `db:"email"`
}

func init() {
	db, err := sqlx.Open("mysql", "root:jiangyibo@tcp(127.0.0.1:4306)/mytest?charset=utf8")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = db
}

func main() {
	var stus []stu
	err := Db.Select(&stus, "SELECT username,password,email FROM userinfo")
	if err != nil {
		fmt.Println("Select error", err)
	}
	fmt.Printf("this is Select res:%v\n", stus)
	var s stu
	err1 := Db.Get(&s, "SELECT username,password,email FROM userinfo where uid = ?", 2)
	if err1 != nil {
		fmt.Println("GET error :", err1)
	} else {
		fmt.Printf("this is GET res:%v", s)
	}
}

//this is Select res:[{wd 123  wd@163.com} {jack 1222  jack@165.com}]
//this is GET res:{jack 1222  jack@165.com}
