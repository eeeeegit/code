package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// gorm.Model
type UserItem struct {
	Account  string `gorm:"primary_key"`
	Password string
	Score    int
	NickName string `gorm:"default:'blue'"`
	Create   time.Time
	Login    time.Time
}

func init() {
	db, err := gorm.Open("mysql", "root:123456@/game?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	if db.HasTable(&UserItem{}) {
		db.AutoMigrate(&UserItem{})
	} else {
		db.CreateTable(&UserItem{})
		db.Create(&UserItem{Account: "test001", Password: "666666", Score: 200, Create: time.Now()})
		db.Create(&UserItem{Account: "test002", Password: "666666", Score: 200, Create: time.Now()})
	}

	db.Close()
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@/game?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	}

	var users []UserItem
	db.Find(&users)

	for _, item := range users {
		fmt.Println(item.Account, item.Password, item.Score)
	}

	db.Close()
}
