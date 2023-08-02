package entity

import (
	"time"

	"gorm.io/gorm"
)

type AuthUser struct {
	Email     string
	FirstName string
	LastName  string
	Salt      string
	Password  string
	Id        int64
}

type OauthAccessToken struct {
	gorm.Model
	Revoked           bool
	Sub               string            `gorm:"size:255; not null"`
	Tid               string            `gorm:"size:255; not null"`
	OauthRefreshToken OauthRefreshToken `gorm:"foreignKey:AccessTokenID; contraint:OnDelete:CASCADE"`
	ExpiredAt         *time.Time        `gorm:"datetime; not null"`
}

type OauthRefreshToken struct {
	gorm.Model
	AccessTokenID uint
	Revoked       bool
	ExpiredAt     *time.Time `gorm:"datetime; not null"`
}
