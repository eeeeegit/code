package db_sql

import (
	"time"
)

type UserInfo struct {
	Account    string `gorm:"primary_key"`
	Password   string
	Score      int
	NickName   string `gorm:"default:'blue'"`
	CreateTime time.Time
	LoginTime  time.Time `gorm:"default:NOW()"`
}
