// @description:
// @file: validate.go
// @date: 2021/12/21

package utils

import (
	"sync"

	"learning/logger"

	"github.com/go-playground/validator/v10"
)

var validate struct {
	instance *validator.Validate
	once     sync.Once
}

func NewValidate() *validator.Validate {
	validate.once.Do(func() {
		validate.instance = validator.New()
	})
	return validate.instance
}

func LogValidateErrors(errs error) {
	if _, ok := errs.(*validator.InvalidValidationError); ok {
		logger.Error(errs)
		return
	}

	for _, err := range errs.(validator.ValidationErrors) {
		logger.Error(err)
	}
}
