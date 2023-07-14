package models

type User struct {
	ID       uint `gorm:primaryKey;default:auto_random`
	FullName string
	Email    string `gorm:"index:email_index;unique"`
	Password string
}
