package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/*
CREATE TABLE `userinfo` (
    `uid` INT(10) NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(64)  DEFAULT NULL,
    `password` VARCHAR(32)  DEFAULT NULL,
    `department` VARCHAR(64)  DEFAULT NULL,
    `email` varchar(64) DEFAULT NULL,
    PRIMARY KEY (`uid`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8
*/

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
	result, err := Db.Exec("INSERT INTO userinfo (username, password, department,email) VALUES (?, ?, ?,?)", "wd", "123", "it", "wd@163.com")
	if err != nil {
		fmt.Println("insert failed,error： ", err)
		return
	}
	id, _ := result.LastInsertId()
	fmt.Println("insert id is :", id)
	_, err1 := Db.Exec("update userinfo set username = ? where uid = ?", "jack", 1)
	if err1 != nil {
		fmt.Println("update failed error:", err1)
	} else {
		fmt.Println("update success!")
	}
	_, err2 := Db.Exec("delete from userinfo where uid = ? ", 1)
	if err2 != nil {
		fmt.Println("delete error:", err2)
	} else {
		fmt.Println("delete success")
	}

}

//insert id is : 1
//update success!
//delete success
