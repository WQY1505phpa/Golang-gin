package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)
var (
	DbHelper *gorm.DB
	err     error
)
func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/space?charset=utf8mb4&parseTime=True&loc=Local"
	DbHelper, err = gorm.Open("mysql",dsn)
	if err !=nil {
		log.Fatal(err.Error)
	}
	//最大连接数
	DbHelper.DB().SetMaxOpenConns(100)
	//闲置连接数
	DbHelper.DB().SetMaxIdleConns(20)
	//最大连接周期
	DbHelper.DB().SetConnMaxLifetime(100 * time.Second)
	//打开日志
	//DbHelper.LogModel(true)
	DbHelper.SingularTable(true)
}
