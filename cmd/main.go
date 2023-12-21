package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iskhakmuhamad/teaservice/config"
	"github.com/iskhakmuhamad/teaservice/delivery"
	"github.com/iskhakmuhamad/teaservice/repository"
	"github.com/iskhakmuhamad/teaservice/usecases"
	"gorm.io/gorm"
)

var (
	db       *gorm.DB                  = config.SetupDatabaseConnection()
	userRepo repository.UserRepository = repository.NewUserRepository(db)

	tokenUC usecases.Token = usecases.NewTokenUc()
	authUC  usecases.Auth  = usecases.NewAuthUC(userRepo)

	authDelivery delivery.AuthDelivery = delivery.NewAuthDelivery(authUC, tokenUC)
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
	}
	r.Run()
}
