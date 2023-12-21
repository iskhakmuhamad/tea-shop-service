package repository

import (
	"context"

	"github.com/iskhakmuhamad/teaservice/model"
	"gorm.io/gorm"
)

type userRepository struct {
	qry *gorm.DB
}

type UserRepository interface {
	InsertUser(ctx context.Context, params *model.User) error
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		qry: db,
	}
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user *model.User

	if err := r.qry.Model(&user).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) InsertUser(ctx context.Context, params *model.User) error {
	var user *model.User

	if err := r.qry.Model(&user).Create(params).Error; err != nil {
		return err
	}
	return nil
}
