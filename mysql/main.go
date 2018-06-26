package main

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type UserItem struct {
	account    string
	password   string
	score      int
	nickname   string
	createtime string
}

func AddOne(item UserItem) error {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/game")
	checkErr(err)
	defer db.Close()

	stmt, err := db.Prepare("INSERT userinfo SET account=?,password=?,score=?,nickname=?,createtime=?")
	checkErr(err)

	res, err := stmt.Exec(item.account, item.password, item.score, item.nickname, item.createtime)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	return nil
}

func IsExist(account string) bool {

	return true
}

func main() {
	item := UserItem{"test002", "tttttt", 0, "小龙", strings.Split(time.Now().String(), ".")[0]}
	AddOne(item)

	// db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/game")
	// checkErr(err)

	// // insert
	// stmt, err := db.Prepare("INSERT userinfo SET account=?,password=?,created=?")
	// checkErr(err)

	// res, err := stmt.Exec("astaxie", "tools", "2008-12-2 22:06:44")
	// checkErr(err)

	// id, err := res.LastInsertId()
	// checkErr(err)

	// fmt.Println(id)
	// // update
	// stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	// checkErr(err)

	// res, err = stmt.Exec("astaxieupdate", id)
	// checkErr(err)

	// affect, err := res.RowsAffected()
	// checkErr(err)

	// fmt.Println(affect)

	// // query
	// rows, err := db.Query("SELECT * FROM userinfo")
	// checkErr(err)

	// for rows.Next() {
	// 	var uid int
	// 	var username string
	// 	var department string
	// 	var created string
	// 	err = rows.Scan(&uid, &username, &department, &created)
	// 	checkErr(err)
	// 	fmt.Println(uid)
	// 	fmt.Println(username)
	// 	fmt.Println(department)
	// 	fmt.Println(created)
	// }

	// // delete
	// stmt, err = db.Prepare("delete from userinfo where uid=?")
	// checkErr(err)

	// res, err = stmt.Exec(id)
	// checkErr(err)

	// affect, err = res.RowsAffected()
	// checkErr(err)

	// fmt.Println(affect)

	// db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
