package validator

import (
	validator2 "github.com/go-playground/validator/v10"
)

type Validator interface {
	Validate(i interface{}) error
}

type ValidatorImpl struct {
	validator *validator2.Validate
}

func (v *ValidatorImpl) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func New() *ValidatorImpl {
	newValidator := validator2.New()
	return &ValidatorImpl{
		newValidator,
	}
}
