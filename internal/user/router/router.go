package user_router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	shared_router "github.com/williamkoller/divine-beast/internal/shared/router"
)

var routerName = "users"

func Routes() []shared_router.Route {
	return []shared_router.Route{
		{
			Method:  "GET",
			Path:    "/" + routerName,
			Handler: func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "user router"}) },
		},
	}
}
