package factory

import "github.com/qrave1/quicknotes/pkg/validator"

func provideValidator() validator.Validator {
	return validator.New()
}
