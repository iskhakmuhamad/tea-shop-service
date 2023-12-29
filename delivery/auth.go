package delivery

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/iskhakmuhamad/teaservice/model"
	"github.com/iskhakmuhamad/teaservice/model/auth"
	"github.com/iskhakmuhamad/teaservice/model/token"
	"github.com/iskhakmuhamad/teaservice/shared"
	"github.com/iskhakmuhamad/teaservice/usecases"
)

type authDelivery struct {
	authUC  usecases.Auth
	tokenUC usecases.Token
}

type AuthDelivery interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

func NewAuthDelivery(
	authUC usecases.Auth,
	tokenUC usecases.Token,
) AuthDelivery {
	return &authDelivery{
		authUC:  authUC,
		tokenUC: tokenUC,
	}
}

func (authDelivery *authDelivery) Register(ctx *gin.Context) {
	var (
		params auth.RegisterRequest
		err    error
	)

	err = ctx.ShouldBind(&params)
	if err != nil {
		res := shared.BuildErrorResponse("failed", "Failed to process request")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	err = authDelivery.authUC.Register(ctx, params)
	if err != nil {
		res := shared.BuildErrorResponse("Register Failed!", err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := shared.BuildResponse("Register Success!", nil)
	ctx.JSON(http.StatusCreated, res)
}

func (authDelivery *authDelivery) Login(ctx *gin.Context) {

	var (
		response *token.ResultResponse
		login    auth.LoginRequest
		user     *model.User
		err      error
	)

	err = ctx.ShouldBind(&login)
	if err != nil {
		res := shared.BuildErrorResponse("Failed to process request", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	user, err = authDelivery.authUC.Login(ctx, login)
	if err != nil {
		res := shared.BuildErrorResponse("Login Failed!", err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	response, _ = authDelivery.tokenUC.GenerateToken(ctx, user)

	res := shared.BuildResponse("Login Success!", response)
	ctx.JSON(http.StatusOK, res)
}

func (authDelivery *authDelivery) Logout(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, err := authDelivery.tokenUC.ValidateToken(authHeader)
	if err != nil {
		res := shared.BuildErrorResponse("Logout Failed!", err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	email := fmt.Sprintf("%v", claims["email"])
	err = authDelivery.authUC.Logout(ctx, email)
	if err != nil {
		res := shared.BuildErrorResponse("Logout Failed!", err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := shared.BuildResponse("Logout Success!", nil)
	ctx.JSON(http.StatusOK, res)
}
