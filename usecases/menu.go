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
	// GetProducts(ctx context.Context, params product.ProductsRequest) ([]model.Product, error)
	CreateMenu(ctx context.Context, params menu.CreateMenuRequest) error
}

func NewMenuUC(menuRepo repository.MenuRepository) Menu {
	return &menuUC{
		menuRepo: menuRepo,
	}
}

// func (u *productUC) GetProducts(ctx context.Context, params product.ProductsRequest) ([]model.Product, error) {

// 	if err := params.Validate(); err != nil {
// 		return nil, err
// 	}

// 	products, err := u.repo.GetProducts(ctx, &model.Product{
// 		ProductName:     params.SearchName,
// 		ProductCategory: params.ProductCategory,
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

//		return products, nil
//	}
func (u *menuUC) CreateMenu(ctx context.Context, params menu.CreateMenuRequest) error {

	if err := params.Validate(); err != nil {
		return err
	}

	err := u.menuRepo.InsertMenu(ctx, &model.Menu{
		Name:     params.Name,
		Price:    params.Price,
		ImageUrl: params.ImageUrl,
	})

	if err != nil {
		return err
	}

	return nil
}
