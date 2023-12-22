package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	err := menuDelivery.menuUC.CreateMenu(ctx, request)
	if err != nil {
		res := shared.BuildErrorResponse("Failed Adding New Menu!", err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := shared.BuildResponse("Success Adding New Menu!", nil)
	ctx.JSON(http.StatusCreated, res)
}
