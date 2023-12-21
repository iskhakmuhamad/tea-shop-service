package auth

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"Address"`
	WANumber string `json:"WANumber"`
}

func (params *RegisterRequest) Validate() error {
	if err := validation.Validate(params.Name, validation.Required); err != nil {
		return errors.New("name must be filled")
	}
	if err := validation.Validate(params.Email, validation.Required); err != nil {
		return errors.New("email must be filled")
	}
	if err := validation.Validate(params.Password, validation.Required); err != nil {
		return errors.New("password must be filled")
	}
	if err := validation.Validate(params.Address, validation.Required); err != nil {
		return errors.New("address must be filled")
	}
	if err := validation.Validate(params.Password, validation.Length(6, 0)); err != nil {
		return errors.New("password minimal 6 character")
	}
	return nil
}
