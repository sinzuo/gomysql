package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	db, err := sqlx.Open("mysql", "stu:1234qwer@tcp(10.0.0.241:3307)/mytest?charset=utf8")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = db
}

func main() {
	rows, err := Db.Query("SELECT username,password,email FROM userinfo")
	if err != nil {
		fmt.Println("query failed,error： ", err)
		return
	}
	for rows.Next() { //循环结果
		var username, password, email string
		err = rows.Scan(&username, &password, &email)
		println(username, password, email)
	}

}

//wd 123 wd@163.com
//jack 1222 jack@165.com
