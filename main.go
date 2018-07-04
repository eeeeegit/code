package main

import (
	db_sql "./db_sql"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {

}

func main() {
	db_sql.Test()
}
