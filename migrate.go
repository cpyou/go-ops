package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-ops/config"
	"go-ops/models/users"
)


func main()  {
	c := config.GetConfig()
	var (
		dbType, dbName, user, password, host, tablePrefix string
	)
	mysql := c.MySQL

	dbType = mysql.Dbtype
	dbName = mysql.Dbname
	user = mysql.Username
	password = mysql.Password
	host = mysql.HOST
	tablePrefix = mysql.Prefix

	var db *gorm.DB
	db, _ = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName, ))
	db.SingularTable(true) // 表名单数

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		if defaultTableName != "casbin_rule" {
			return  defaultTableName
		}
		return tablePrefix + defaultTableName
	}

	db.AutoMigrate(&users.User{})
	db.AutoMigrate(&users.UserProfile{})
}

