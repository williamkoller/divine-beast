package healthz_router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHealthzRouter(r *gin.Engine) {
	healthzRoutes := r.Group("/healthz")
	healthzRoutes.GET("", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) })

}
