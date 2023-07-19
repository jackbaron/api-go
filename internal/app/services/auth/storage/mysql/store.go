package mysql

import (
	"context"
	"time"

	"github.com/nhatth/api-service/internal/app/services/auth/entity"
	errorPkg "github.com/nhatth/api-service/pkg/errors"
	"github.com/nhatth/api-service/pkg/jwt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type mysqlStore struct {
	db *gorm.DB
}

var table = "users"

func NewMySQLStore(db *gorm.DB) *mysqlStore {
	return &mysqlStore{db: db}
}

func (store *mysqlStore) AddNewUser(ctx context.Context, data *entity.AuthRegister) error {
	data.CreatedAt = time.Now().UTC()

	if err := store.db.Table(table).Create(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (store *mysqlStore) GetUser(ctx context.Context, email string) (*entity.AuthUser, error) {
	var data entity.AuthUser

	if err := store.db.Table(table).Where("email = ?", email).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorPkg.ErrRecordNotFound
		}

		return nil, errors.WithStack(err)
	}
	return &data, nil
}

func (store *mysqlStore) StoreAccessToken(ctx context.Context, data *jwt.TokenDetails, tid, sub string) (*entity.OauthAccessToken, error) {

	accessToken := entity.OauthAccessToken{
		Sub:       sub,
		Tid:       tid,
		ExpiredAt: *data.AccessTokenExpired,
		OauthRefreshToken: entity.OauthRefreshToken{
			ExpiredAt: *data.RefreshTokenExpired,
		},
	}

	if err := store.db.Create(&accessToken).Error; err != nil {

		return nil, errors.WithStack(err)
	}

	return &accessToken, nil
}
