package repository

import "gorm.io/gorm"

type RoleRepository struct {
	DB *gorm.DB
}

type RoleRepositoryImpl interface {
}

func NewUserRepository(db *gorm.DB) RoleRepositoryImpl {
	return &RoleRepository{
		DB: db,
	}
}
