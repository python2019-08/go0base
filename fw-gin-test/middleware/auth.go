package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 权限检查
func AuthCheck(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userName, _ := c.Get("user_name")
	fmt.Printf("AuthCheck: user_id=%s,user_name=%s\n", userId, userName)
	c.Next()
}

var token = "123456"

func TokenCheck(c *gin.Context) {
	accessToken := c.Request.Header.Get("access_toke")
	if accessToken != token {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "token 检查失败",
		})

		c.AbortWithError(http.StatusInternalServerError, errors.New("token 检查失败"))
	}

	// if check ok
	c.Set("user_name", "nick")
	c.Set("user_id", "10001")
	// Next should be used only inside middleware.
	// It executes the pending handlers in the chain inside the calling handler.
	c.Next()
}
