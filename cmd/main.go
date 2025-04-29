package main

import (
	"os"

	"itv-go/internal/router"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// @title Movie API
// @version 1.0
// @host localhost:8080
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @Security BearerAuth
func main() {
	app := fx.New(
		fx.Provide(
			router.NewRouter,
		),
		fx.Invoke(func(r *gin.Engine) {
			port := os.Getenv("PORT")
			if port == "" {
				port = "8080"
			}
			if err := r.Run(":" + port); err != nil {
				panic("failed to start server: " + err.Error())
			}
		}),
	)

	app.Run()
}
