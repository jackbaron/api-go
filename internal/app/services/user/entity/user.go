package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `gorm:"size:255; not null"`
	LastName  string `gorm:"size:255; not null"`
	Salt      string `gorm:"size:255; not null"`
	Email     string `gorm:"unique; index:idx_email; not null"`
	Password  string `gorm:"size:255; not null"`
}
