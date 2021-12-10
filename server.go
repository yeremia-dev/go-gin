package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yeremia-dev/go-gin/config"
	"github.com/yeremia-dev/go-gin/controller"
	"github.com/yeremia-dev/go-gin/repository"
	"github.com/yeremia-dev/go-gin/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
)

func main() {
	r := gin.Default()
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	r.Run()
}
