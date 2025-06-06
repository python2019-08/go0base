package fwgintest

import (
	"go0base/fw-gin-test/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.InitRoutes(r)
	r.Run(":8080")
}
