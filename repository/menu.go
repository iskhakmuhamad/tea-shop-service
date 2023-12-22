package repository

import (
	"context"

	"github.com/iskhakmuhamad/teaservice/model"
	"gorm.io/gorm"
)

type menuRepository struct {
	qry *gorm.DB
}

type MenuRepository interface {
	InsertMenu(ctx context.Context, params *model.Menu) error
	GetMenus(ctx context.Context, params model.Menu) ([]model.Menu, error)
}

func NewMenuRepository(
	db *gorm.DB,
) MenuRepository {
	return &menuRepository{
		qry: db,
	}
}

func (r *menuRepository) InsertMenu(ctx context.Context, params *model.Menu) error {
	if err := r.qry.Model(&model.Menu{}).Create(params).Error; err != nil {
		return err
	}
	return nil
}

func (r *menuRepository) GetMenus(ctx context.Context, params model.Menu) ([]model.Menu, error) {
	var (
		menus []model.Menu
	)

	db := r.qry.Model(model.Menu{})

	if params.UserID != 0 {
		db = db.Where("user_id = ?", params.UserID)
	}

	if err := db.Find(&menus).Error; err != nil {
		return nil, err
	}

	return menus, nil
}
