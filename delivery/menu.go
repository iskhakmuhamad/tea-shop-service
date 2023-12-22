package delivery

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/iskhakmuhamad/teaservice/model"
	"github.com/iskhakmuhamad/teaservice/model/menu"
	"github.com/iskhakmuhamad/teaservice/shared"
	"github.com/iskhakmuhamad/teaservice/usecases"
)

type menuDelivery struct {
	menuUC  usecases.Menu
	tokenUC usecases.Token
}

type MenuDelivery interface {
	CreateMenu(ctx *gin.Context)
	GetMenus(ctx *gin.Context)
}

func NewMenuDelivery(menuUC usecases.Menu, tokenUC usecases.Token) MenuDelivery {
	return &menuDelivery{
		menuUC:  menuUC,
		tokenUC: tokenUC,
	}
}

func (menuDelivery *menuDelivery) CreateMenu(ctx *gin.Context) {
	var (
		request menu.CreateMenuRequest
	)

	if err := ctx.Bind(&request); err != nil {
		res := shared.BuildErrorResponse("Failed to process request", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, err := menuDelivery.tokenUC.ValidateToken(authHeader)
	if err != nil {
		response := shared.BuildErrorResponse("Malformat Token", err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := int64(claims["id"].(float64))
	if userID == 0 {
		response := shared.BuildErrorResponse("Doesnt have permission", "Doesnt Get User ID")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	} else {
		request.UserID = userID
	}

	err = menuDelivery.menuUC.CreateMenu(ctx, request)

	if err != nil {
		res := shared.BuildErrorResponse("Failed Adding New Menu!", err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := shared.BuildResponse("Success Adding New Menu!", nil)
	ctx.JSON(http.StatusCreated, res)
}

func (menuDelivery *menuDelivery) GetMenus(ctx *gin.Context) {
	var (
		request model.Menu
	)

	authHeader := ctx.GetHeader("Authorization")
	token, err := menuDelivery.tokenUC.ValidateToken(authHeader)
	if err != nil {
		response := shared.BuildErrorResponse("Malformat Token", err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := int64(claims["id"].(float64))
	if userID == 0 {
		response := shared.BuildErrorResponse("Doesnt have permission", "Doesnt Get User ID")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	} else {
		request.UserID = userID
	}

	menus, err := menuDelivery.menuUC.GetMenus(ctx, request)

	if err != nil {
		res := shared.BuildErrorResponse("Failed Get Menus", err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := shared.BuildResponse("Success get menus", menus)
	ctx.JSON(http.StatusOK, res)
}
