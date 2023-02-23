package service

import (
	"context"

	"github.com/samber/do"

	userv1 "github.com/saltfishpr/go-learning/gen/go/user/v1"
	"github.com/saltfishpr/go-learning/internal/user/biz"
)

type UserService struct {
	userv1.UnimplementedUserServiceServer
	uuc *biz.UserUseCase
}

func NewUserService(i *do.Injector) (*UserService, error) {
	return &UserService{
		uuc: do.MustInvoke[*biz.UserUseCase](i),
	}, nil
}

func (s *UserService) CreateUser(
	ctx context.Context,
	req *userv1.CreateUserRequest,
) (*userv1.User, error) {
	u := &userv1.User{
		Name:     "/user/1",
		Username: "test",
		Email:    "test@example.com",
	}
	return u, nil
}

func (s *UserService) GetUser(
	ctx context.Context,
	req *userv1.GetUserRequest,
) (*userv1.User, error) {
	u := &userv1.User{
		Name:     "/user/1",
		Username: "test",
		Email:    "test@example.com",
	}
	return u, nil
}
