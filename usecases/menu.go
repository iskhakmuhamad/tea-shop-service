package usecases

import (
	"context"

	"github.com/iskhakmuhamad/teaservice/model"
	"github.com/iskhakmuhamad/teaservice/model/menu"
	"github.com/iskhakmuhamad/teaservice/repository"
)

type menuUC struct {
	menuRepo repository.MenuRepository
}

type Menu interface {
	CreateMenu(ctx context.Context, params menu.CreateMenuRequest) error
	GetMenus(ctx context.Context, params model.Menu) ([]model.Menu, error)
	GetMenuByID(ctx context.Context, params model.Menu) (*model.Menu, error)
	UpdateMenu(ctx context.Context, params menu.UpdateMenuRequest) error
}

func NewMenuUC(
	menuRepo repository.MenuRepository) Menu {
	return &menuUC{
		menuRepo: menuRepo,
	}
}

func (u *menuUC) GetMenus(ctx context.Context, params model.Menu) ([]model.Menu, error) {
	menus, err := u.menuRepo.GetMenus(ctx, model.Menu{
		UserID: params.UserID,
	})
	if err != nil {
		return nil, err

	}
	return menus, nil
}

func (u *menuUC) GetMenuByID(ctx context.Context, params model.Menu) (*model.Menu, error) {
	menu, err := u.menuRepo.GetMenuByID(ctx, model.Menu{
		UserID: params.UserID,
	})
	if err != nil {
		return nil, err

	}
	return menu, nil
}

func (u *menuUC) CreateMenu(ctx context.Context, params menu.CreateMenuRequest) error {

	if err := params.Validate(); err != nil {
		return err
	}

	err := u.menuRepo.InsertMenu(ctx, &model.Menu{
		Name:     params.Name,
		Price:    params.Price,
		ImageUrl: params.ImageUrl,
		UserID:   params.UserID,
	})

	if err != nil {
		return err
	}

	return nil
}

func (u *menuUC) UpdateMenu(ctx context.Context, params menu.UpdateMenuRequest) error {
	if err := params.Validate(); err != nil {
		return err
	}
	if err := u.menuRepo.UpdateMenu(ctx, model.Menu{
		ID:       params.MenuID,
		Name:     params.Name,
		Price:    params.Price,
		ImageUrl: params.ImageUrl,
		UserID:   params.UserID,
	}); err != nil {
		return err
	}
	return nil
}
