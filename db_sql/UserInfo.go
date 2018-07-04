package db_sql

import (
	"time"
)

type IUserInfo interface {
	//根据参数，自动完成数据库查询
	Retrieve(params map[string]interface{}, args ...interface{}) map[string]interface{}
	//根据参数，自动完成数据库插入
	Create(params map[string]interface{}, args ...interface{}) map[string]interface{}
	//根据参数，自动完成数据库更新（只支持单条）
	Update(params map[string]interface{}, args ...interface{}) map[string]interface{}
	//根据参数，自动完成数据库删除（只支持单条）
	Delete(params map[string]interface{}, args ...interface{}) map[string]interface{}
	//手写查询sql支持
	QuerySql(sql string, values []interface{}, params map[string]interface{}) map[string]interface{}
	//手写非查询sql支持
	ExecSql(sql string, values []interface{}) map[string]interface{}
	//批量插入或更新
	InsertBatch(tablename string, els []interface{}) map[string]interface{}
	//事务支持
	TransGo(objs map[string]interface{}) map[string]interface{}
}

type UserInfo struct {
	Account    string `gorm:"primary_key"`
	Password   string
	Score      int
	NickName   string `gorm:"default:'blue'"`
	CreateTime time.Time
	LoginTime  time.Time `gorm:"default:NOW()"`
}
