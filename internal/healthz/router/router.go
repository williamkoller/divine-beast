package healthz_router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	shared_router "github.com/williamkoller/divine-beast/internal/shared/router"
)

func Routes() []shared_router.Route {
	return []shared_router.Route{
		{
			Method:  "GET",
			Path:    "/healthz",
			Handler: func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) },
		},
		{
			Method:  "GET",
			Path:    "/status",
			Handler: func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "running"}) },
		},
	}
}
