package service

import "userservice/pkg/v1/permission/repository"

type PermissionService struct {
	PermissionRepository repository.PermissionRepositoryImpl
}

type PermissionServiceImpl interface {
}

func NewPermissionService(r repository.PermissionRepositoryImpl) PermissionServiceImpl {
	return &PermissionService{PermissionRepository: r}
}
