package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type Address struct {
	gorm.Model
	UserID  uint   `gorm:"uniqueIndex" json:"userID"` // Foreign key to User
	Street  string `gorm:"not null" json:"street"`
	City    string `gorm:"not null" json:"city"`
	State   string `gorm:"not null" json:"state"`
	ZipCode string `gorm:"not null" json:"zipCode"`
}

type UserRole string

const (
	RoleAdmin      UserRole = "admin"
	RoleCustomer   UserRole = "customer"
	RoleSuperAdmin UserRole = "superadmin"
)

// User represents the user model
type User struct {
	gorm.Model
	Username      string    `gorm:"unique;not null" json:"username"`
	Email         string    `gorm:"unique;not null" json:"email"`
	Password      string    `gorm:"not null" json:"password"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	IsActive      bool      `json:"isActive"`
	DeactivatedAt time.Time `json:"deactivatedAt,omitempty"`
	DOB           time.Time `json:"dob"`
	Roles         []Role    `gorm:"many2many:user_roles" json:"roles"`
}

// Permission represents the permission model
type Permission struct {
	gorm.Model
	Name        string `gorm:"unique;not null" json:"name"`
	Description string `json:"description"`
}

// Role represents the role model
type Role struct {
	gorm.Model
	Name        UserRole     `gorm:"unique;not null" json:"name"`
	Permissions []Permission `gorm:"many2many:role_permissions" json:"permissions"`
	Users       []User       `gorm:"many2many:user_roles" json:"users"`
}

// BeforeCreate is a GORM hook that is called before creating a new record
func (u *User) BeforeCreate(tx *gorm.DB) error {
	// Hash the password before creating the user
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	u.IsActive = true
	return nil
}
