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
	rows, err := Db.Queryx("SELECT username,password,email FROM userinfo")
	if err != nil {
		fmt.Println("Qeryx failed,error： ", err)
		return
	}
	for rows.Next() { //循环结果
		var stu1 stu
		err = rows.StructScan(&stu1) // 转换为结构体
		fmt.Println("stuct data：", stu1.Username, stu1.Password)
	}
}

//stuct data： wd 123
//stuct data： jack 1222
