package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"log"
	"web/common"
	"web/controller"
	"web/database"
	"web/router"
	"web/service"
)

// RunServer function to start the server
func RunServer(lifecycle fx.Lifecycle, r *gin.Engine) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Println("Starting server on :8080")
				if err := r.Run(":8080"); err != nil {
					log.Fatal("Failed to start server:", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Stopping server gracefully...")
			return nil
		},
	})
}

// @title           My Gin API
// @version         1.0
// @description     This is a sample Gin server with Bearer JWT auth.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.myapi.com/support
// @contact.email  support@myapi.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your bearer token in the format: Bearer <token>
func main() {
	app := fx.New(
		router.RouterModule,
		database.DatabaseModule,
		controller.ControllerModule,
		controller.AuthControllerModule,
		common.AuthModule,
		service.UserModule,
		fx.Invoke(RunServer),
	)

	app.Run()
}
