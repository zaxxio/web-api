package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"

	"web/common"
	"web/controller"
	_ "web/docs" // Import Swagger docs
)

func NewRouter(
	userController *controller.UserController,
	authController *controller.AuthController,
) *gin.Engine {

	router := gin.Default()

	router.GET("/swagger-ui/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// CORS configuration (example for a frontend at http://localhost:3000)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Adjust as needed
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	publicGroup := router.Group("")
	{
		publicGroup.POST("/auth/signin", authController.SignIn)
		publicGroup.POST("/auth/signup", authController.SignUp)
	}

	privateGroup := router.Group("").Use(common.JWTMiddleware())
	{
		privateGroup.POST("/users", userController.CreateUser)
		privateGroup.GET("/users", userController.GetUsers)
	}

	return router
}

var RouterModule = fx.Module(
	"router",
	fx.Provide(NewRouter),
)
