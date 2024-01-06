package service

import (
	"userservice/pkg/v1/role/repository"
)

type RoleService struct {
	RoleRepository repository.RoleRepositoryImpl
}

type RoleServiceImpl interface {
}

func NewRoleService(r repository.RoleRepositoryImpl) RoleServiceImpl {
	return &RoleService{
		RoleRepository: r,
	}
}
