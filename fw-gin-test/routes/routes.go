package routes

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.Engine) {
	// r.Use() // 全局中间件注册
	initApi(r)

	initCourse(r)
}
