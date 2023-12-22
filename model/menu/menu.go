package menu

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateMenuRequest struct {
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	ImageUrl string `json:"image_url"`
}

func (params *CreateMenuRequest) Validate() error {
	if err := validation.Validate(params.Name, validation.Required); err != nil {
		return errors.New("name must be filled")
	}
	if err := validation.Validate(params.Price, validation.Required); err != nil {
		return errors.New("price must be filled")
	}
	return nil
}
