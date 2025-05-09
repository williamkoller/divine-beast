package user_router

import (
	"github.com/gin-gonic/gin"
	handler_adduser "github.com/williamkoller/divine-beast/internal/user/handler/add-user"
)

var routerName = "users"

func RegisterUserRoutes(r *gin.Engine) {
	userRoutes := r.Group("/" + routerName)

	userRoutes.POST("", handler_adduser.AddUser)

}
