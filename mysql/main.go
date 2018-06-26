package main

import (
	"database/sql"
	"fmt"

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
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/game")
	checkErr(err)
	defer db.Close()

	var checkStr string
	err = db.QueryRow("select account from userinfo where account = ?", account).Scan(&checkStr)
	if err == nil {
		return true
	} else {
		return false
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println(IsExist("test003"))
	// item := UserItem{"test002", "tttttt", 0, "小龙", strings.Split(time.Now().String(), ".")[0]}
	// AddOne(item)

}
