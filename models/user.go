package models

import (
	"time"

	// "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Username  string         `gorm:"size:50;not null;unique" json:"username"`
	Password  string         `gorm:"size:255;not null" json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Roles     []Role         `gorm:"many2many:user_roles;" json:"roles"`
}

// Role represents a user role in the system
type Role struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"size:50;not null;unique" json:"name"`
	Description string         `gorm:"size:255" json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Users       []User         `gorm:"many2many:user_roles;" json:"-"`
}

