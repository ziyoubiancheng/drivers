package main

import (
	"github.com/ziyoubiancheng/drivers"
	"github.com/ziyoubiancheng/drivers/pkg/database/mysql"
	"github.com/ziyoubiancheng/drivers/pkg/logger"
)

var cfg = `
[drivers.mysql.default]
    debug = true
    level = "panic"
    network = "tcp"
    dialect = "mysql"
    addr = "127.0.0.1:3306"
    username = "root"
    password = "root"
    db = "shop"
    charset = "utf8"
    parseTime = "True"
    loc = "Local"
    timeout = "1s"
    readTimeout = "1s"
    writeTimeout = "1s"
    maxOpenConns = 30
    maxIdleConns = 10
    connMaxLifetime = "300s"

`
var (
	Db *mysql.Client
)

func main() {
	if err := drivers.Container(
		[]byte(cfg),
		mysql.Register,
		logger.Register,
	); err != nil {
		panic(err)
	}

	initCaller()
	type User struct {
		Uid  int
		Name string
	}
	u := User{}
	Db.Table("user").Where("uid=?", 1).Find(&u)
}

func initCaller() {
	Db = mysql.Caller("default")
}
