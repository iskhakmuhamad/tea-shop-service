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
