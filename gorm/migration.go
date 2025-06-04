package gorm

import "fmt"

// https://gorm.io/zh_CN/docs/migration.html

func init() {
	fmt.Println("gorm/migration.go:init()..start")
	defer fmt.Println("gorm/migration.go:init()..end")

	// DB.AutoMigrate(&Teacher{})

	////// DB.Migrator().CurrentDatabase()
}
