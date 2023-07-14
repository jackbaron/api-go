package mysql

import (
	"context"

	"github.com/nhatth/api-service/internal/app/services/auth/entity"
	errorPkg "github.com/nhatth/api-service/pkg/errors"
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
