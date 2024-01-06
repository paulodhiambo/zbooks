package repository

import "gorm.io/gorm"

type PermissionRepository struct {
	DB *gorm.DB
}

type PermissionRepositoryImpl interface {
}

func NewPermissionRepository(db *gorm.DB) PermissionRepositoryImpl {
	return &PermissionRepository{
		DB: db,
	}
}
