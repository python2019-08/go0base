package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Register",
	})
}
