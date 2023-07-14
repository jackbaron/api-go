package bussines

import (
	"context"

	"github.com/nhatth/api-service/internal/app/services/auth/entity"
	errorPkg "github.com/nhatth/api-service/pkg/errors"
)

type AuthRepository interface {
	AddNewUser(ctx context.Context, data *entity.AuthRegister) error
	GetUser(ctx context.Context, email string) (*entity.AuthUser, error)
}

type Hashser interface {
	RandomStr(length int) (string, error)
	HashPassword(salt, password string) (string, error)
	CompareHashPassword(hashedPassword, salt, password string) bool
}

type business struct {
	repository AuthRepository
	hasher     Hashser
}

func NewAuthBusiness(repository AuthRepository, hasher Hashser) *business {
	return &business{repository: repository, hasher: hasher}
}

func (bus *business) Register(ctx context.Context, data *entity.AuthRegister) error {
	if err := data.Validate(); err != nil {
		return err
	}

	_, err := bus.repository.GetUser(ctx, data.Email)

	if err == nil {
		return errorPkg.ErrBadRequest.WithError(entity.ErrEmailHasExisted.Error())
	} else if err != errorPkg.ErrRecordNotFound {

		return errorPkg.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	salt, err := bus.hasher.RandomStr(16)

	if err != nil {
		return errorPkg.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	passwordHashed, err := bus.hasher.HashPassword(salt, data.Password)

	if err != nil {
		return errorPkg.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	data.Password = passwordHashed

	data.Salt = salt

	if err := bus.repository.AddNewUser(ctx, data); err != nil {
		return errorPkg.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	return nil
}
