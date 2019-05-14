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
	row := Db.QueryRow("SELECT username,password,email FROM userinfo where uid = ?", 1) // QueryRow返回错误，错误通过Scan返回
	var username, password, email string
	err := row.Scan(&username, &password, &email)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("this is QueryRow res:[%s:%s:%s]\n", username, password, email)
	var s stu
	err1 := Db.QueryRowx("SELECT username,password,email FROM userinfo where uid = ?", 2).StructScan(&s)
	if err1 != nil {
		fmt.Println("QueryRowx error :", err1)
	} else {
		fmt.Printf("this is QueryRowx res:%v", s)
	}
}

//this is QueryRow res:[wd:123:wd@163.com]
//this is QueryRowx res:{jack 1222  jack@165.com}
