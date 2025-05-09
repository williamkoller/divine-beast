package order_router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var routerName = "orders"

func RegisterOrderRoutes(r *gin.Engine) {
	orderRoutes := r.Group("/" + routerName)

	orderRoutes.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

}
