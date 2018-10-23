package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
)

var (
	dbConn *sql.DB
	err    error
)
//初始化数据库的连接
//具体配置见配置文件
func init() {

	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	dbConn, err = sql.Open("mysql", dsn)
	//dbConn, err = sql.Open("mysql", "root:w199547@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}