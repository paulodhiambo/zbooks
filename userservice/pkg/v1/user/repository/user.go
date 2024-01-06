package repository

import "gorm.io/gorm"

type UserRepository struct {
	DB *gorm.DB
}

type UserRepositoryImpl interface {
}

func NewUserRepository(db *gorm.DB) UserRepositoryImpl {
	return &UserRepository{
		DB: db,
	}
}
