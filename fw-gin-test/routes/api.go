package routes

import (
	"go0base/fw-gin-test/web"

	"github.com/gin-gonic/gin"
)

func initApi(r *gin.Engine) {

	// http://localhost:8080/api
	api := r.Group("/api")

	// http://localhost:8080/api/v1
	v1 := api.Group("/v1")
	v1.GET("/ping", web.Ping)

	// http://localhost:8080/api/v1/login
	// var i int = 10
	// v1.POST("/login", func(ctx *gin.Context) {
	// 	fmt.Println(i)
	// })
	v1.POST("/login", web.Login)
	// http://localhost:8080/api/v1/register
	v1.POST("/register", web.Register)
}
