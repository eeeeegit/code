package db_sql

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	Account    string `gorm:"primary_key"`
	Password   string
	Score      int
	NickName   string `gorm:"default:'blue'"`
	CreateTime time.Time
	LoginTime  time.Time `gorm:"default:NOW()"`
}
