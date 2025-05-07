package main

import (
	"net/http"
	"time"

	healthz_router "github.com/williamkoller/divine-beast/internal/healthz/router"
	order_router "github.com/williamkoller/divine-beast/internal/order/router"
	shared_router "github.com/williamkoller/divine-beast/internal/shared/router"
	user_router "github.com/williamkoller/divine-beast/internal/user/router"
)

func main() {

	healthzRoutes := healthz_router.Routes()
	userRoutes := user_router.Routes()
	orderRoutes := order_router.Routes()

	// Registro
	allRoutes := append(healthzRoutes, userRoutes...)
	allRoutes = append(allRoutes, orderRoutes...)
	r := shared_router.RegisterRoutes(allRoutes)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
