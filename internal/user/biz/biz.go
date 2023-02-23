// Package biz .
package biz

import "github.com/samber/do"

type UserUseCase struct{}

func NewUserUseCase(i *do.Injector) (*UserUseCase, error) {
	return &UserUseCase{}, nil
}
