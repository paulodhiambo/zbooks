package service

import (
	"userservice/pkg/v1/user/repository"
)

type UserService struct {
	UserRepository repository.UserRepositoryImpl
}

type UserServiceImpl interface {
}

func NewUserService(r repository.UserRepositoryImpl) UserServiceImpl {
	return &UserService{
		UserRepository: r,
	}
}
