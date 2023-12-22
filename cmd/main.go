package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iskhakmuhamad/teaservice/config"
	"github.com/iskhakmuhamad/teaservice/delivery"
	"github.com/iskhakmuhamad/teaservice/middleware"
	"github.com/iskhakmuhamad/teaservice/repository"
	"github.com/iskhakmuhamad/teaservice/usecases"
	"gorm.io/gorm"
)

var (
	db       *gorm.DB                  = config.SetupDatabaseConnection()
	userRepo repository.UserRepository = repository.NewUserRepository(db)
	menuRepo repository.MenuRepository = repository.NewMenuRepository(db)

	tokenUC usecases.Token = usecases.NewTokenUc()
	authUC  usecases.Auth  = usecases.NewAuthUC(userRepo)
	menuUC  usecases.Menu  = usecases.NewMenuUC(menuRepo)

	authDelivery delivery.AuthDelivery = delivery.NewAuthDelivery(authUC, tokenUC)
	menuDelivery delivery.MenuDelivery = delivery.NewMenuDelivery(menuUC, tokenUC)
)

func main() {
	r := gin.Default()

	apiRoutes := r.Group("api/v1/")
	{
		authRoutes := apiRoutes.Group("auth")
		{
			authRoutes.POST("/register", authDelivery.Register)
			authRoutes.POST("/login", authDelivery.Login)
			authRoutes.POST("/logout", authDelivery.Logout)
		}
		menuRoutes := apiRoutes.Group("menu")
		{
			menuRoutes.POST("/", middleware.AuthorizeJWT(tokenUC), menuDelivery.CreateMenu)
		}
	}
	r.Run()
}
