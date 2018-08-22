package gorm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"exchange-currency/config"
)

var (
	dbConfig  = config.Config.DB
	mysqlConn *gorm.DB
	err       error
	err1       error
)

// initialize database
func init() {
	if dbConfig.Driver == "mysql" {
		setupMysqlConn()
	}
}

// setupMysqlConn: setup mysql database connection using the configuration from config.yml
func setupMysqlConn() {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	mysqlConn, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	err = mysqlConn.DB().Ping()
	if err != nil {
		panic(err)
	}
	mysqlConn.LogMode(true)
}

// MysqlConn: return mysql connection from gorm ORM
func MysqlConn() *gorm.DB {
	return mysqlConn
}
