package main

import (
	"log"

	"github.com/gin-gonic/gin"
	healthz_router "github.com/williamkoller/divine-beast/internal/healthz/router"
	order_router "github.com/williamkoller/divine-beast/internal/order/router"
	user_router "github.com/williamkoller/divine-beast/internal/user/router"
)

func main() {
	r := gin.Default()

	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	healthz_router.RegisterHealthzRouter(r)
	user_router.RegisterUserRoutes(r)
	order_router.RegisterOrderRoutes(r)

	log.Println("Starting server on port 8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
