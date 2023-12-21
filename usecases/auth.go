package usecases

import (
	"context"
	"errors"

	"github.com/iskhakmuhamad/teaservice/model"
	"github.com/iskhakmuhamad/teaservice/model/auth"
	"github.com/iskhakmuhamad/teaservice/repository"
	"github.com/iskhakmuhamad/teaservice/shared"
)

type authUC struct {
	userRepo repository.UserRepository
}

type Auth interface {
	Register(ctx context.Context, param auth.RegisterRequest) error
	Login(ctx context.Context, params auth.LoginRequest) (*model.User, error)
	Logout(ctx context.Context, email string) error
}

func NewAuthUC(userRepo repository.UserRepository) Auth {
	return &authUC{
		userRepo: userRepo,
	}
}

func (u *authUC) Register(ctx context.Context, param auth.RegisterRequest) error {
	var (
		encryptedPassword string
		err               error
		user              *model.User
	)
	if err = param.Validate(); err != nil {
		return err
	}
	user, _ = u.userRepo.GetUserByEmail(ctx, param.Email)
	if user != nil {
		return errors.New("email already used")
	}

	encryptedPassword, err = shared.EncryptPassword(param.Password)
	if err != nil {
		return err
	}

	err = u.userRepo.InsertUser(ctx, &model.User{
		Name:     param.Name,
		Email:    param.Email,
		Address:  param.Address,
		Password: encryptedPassword,
		WANumber: param.WANumber,
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *authUC) Login(ctx context.Context, params auth.LoginRequest) (*model.User, error) {

	var (
		err  error
		user *model.User
	)

	if err = params.Validate(); err != nil {
		return nil, err
	}

	user, err = u.userRepo.GetUserByEmail(ctx, params.Email)
	if err != nil || user == nil {
		return nil, errors.New("email not found")
	}

	err = shared.CheckPassword(params.Password, user.Password)
	if err != nil {
		return nil, errors.New("wrong password")
	}

	user, _ = u.userRepo.GetUserByEmail(ctx, params.Email)

	return user, nil
}

func (u *authUC) Logout(ctx context.Context, email string) error {

	// var (
	// 	err  error
	// 	user *models.User
	// )

	// user, _ = u.repo.GetUserByEmail(ctx, email)

	return nil
}
