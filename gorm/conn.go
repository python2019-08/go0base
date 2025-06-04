package gorm

// https://gorm.io/zh_CN/docs/connecting_to_the_database.html

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var dsn = "root:123456@tcp(192.168.1.108:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"

func init() {
	fmt.Println("gorm/conn.go:init()..start")
	defer fmt.Println("gorm/conn.go:init()..end")

	var err error
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln(err)
	}

	// set connection pool
	setPool(DB)
}

func setPool(db *gorm.DB) {
	sqLDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqLDB.SetMaxOpenConns(10)
	sqLDB.SetMaxIdleConns(5)
	sqLDB.SetConnMaxLifetime(time.Hour)

}
