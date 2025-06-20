package main

import (
	"go0base/cmd"
)

//go:generate swag init --parseDependency --parseDepth=6 --instanceName admin -o ./docs/go0base

// @title go0base API
// @version 2.0.0
// @description 基于Gin + Vue + Element UI的前后端分离权限管理系统的接口文档
// @description 谢谢！
// @license.name MIT
// @license.url https://github.com/go-admin-team/go-admin/blob/master/LICENSE.md

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	app := cmd.NewApp()
	app.Execute()
}

// func main() {
// 	logrusdemo.DemoLogrus()

// 	fmt.Println("----main()....start")
// 	defer fmt.Println("----main()....end")
// 	go_pool.Go_Pool()

// 	s1 := make([]byte, 20)
// 	s1[1] = 99
// 	s1[2] = 100
// 	fmt.Println(s1)
// 	test.Test_struct()

// 	// gorm.CreateRecord()
// 	gorm.Query()

// 	fmt.Println("Hello, World!")
// }
