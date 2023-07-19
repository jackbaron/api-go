package bussines

import (
	"context"

	"github.com/google/uuid"

	"github.com/nhatth/api-service/internal/app/pkg/uid"
	"github.com/nhatth/api-service/internal/app/services/auth/entity"
	"github.com/nhatth/api-service/internal/common"
	errorPkg "github.com/nhatth/api-service/pkg/errors"
	"github.com/nhatth/api-service/pkg/jwt"
)

type AuthRepository interface {
	AddNewUser(ctx context.Context, data *entity.AuthRegister) error
	GetUser(ctx context.Context, email string) (*entity.AuthUser, error)
	StoreAccessToken(ctx context.Context, data *jwt.TokenDetails, tid, sub string) (*entity.OauthAccessToken, error)
}

type Hashser interface {
	RandomStr(length int) (string, error)
	HashPassword(salt, password string) (string, error)
	CompareHashPassword(hashedPassword, salt, password string) bool
}

type business struct {
	repository AuthRepository
	hasher     Hashser
	jwt        common.JWTProvider
}

func NewAuthBusiness(repository AuthRepository, hasher Hashser, jwt common.JWTProvider) *business {
	return &business{repository: repository, hasher: hasher, jwt: jwt}
}

func (bus *business) Register(ctx context.Context, data *entity.AuthRegister) (map[string]string, error) {

	errors, errorValidate := data.Validate()

	if errorValidate {
		return errors, errorPkg.ErrBadRequest.WithError(entity.ErrorValidateFailed.Error())
	}

	_, err := bus.repository.GetUser(ctx, data.Email)

	if err == nil {
		return errors, errorPkg.ErrBadRequest.WithError(entity.ErrEmailHasExisted.Error())
	} else if err != errorPkg.ErrRecordNotFound {

		return errors, errorPkg.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	salt, err := bus.hasher.RandomStr(16)

	if err != nil {
		return errors, errorPkg.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	passwordHashed, err := bus.hasher.HashPassword(salt, data.Password)

	if err != nil {
		return errors, errorPkg.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	data.Password = passwordHashed

	data.Salt = salt

	if err := bus.repository.AddNewUser(ctx, data); err != nil {
		return errors, errorPkg.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	return errors, nil
}

func (bus *business) Login(ctx context.Context, data *entity.AuthEmailPassword) (*entity.TokenResponse, map[string]string, error) {

	errors, errorValidate := data.Validate()

	if errorValidate {
		return nil, errors, errorPkg.ErrBadRequest.WithError(entity.ErrorValidateFailed.Error())
	}

	authData, err := bus.repository.GetUser(ctx, data.Email)

	if err != nil {
		return nil, errors, errorPkg.ErrBadRequest.WithError(entity.ErrLoginFailed.Error()).WithDebug(err.Error())
	}

	if !bus.hasher.CompareHashPassword(authData.Password, authData.Salt, data.Password) {

		return nil, errors, errorPkg.ErrBadRequest.WithError(entity.ErrLoginFailed.Error())
	}

	uid := uid.NewUID(uint32(authData.Id), 1, 1)

	sub := uid.String()

	tid := uuid.New().String()

	token, err := bus.jwt.IssueToken(ctx, tid, sub)

	if err != nil {

		return nil, errors, errorPkg.ErrBadRequest.WithError(entity.ErrLoginFailed.Error()).WithDebug(err.Error())
	}

	_, err = bus.repository.StoreAccessToken(ctx, token, tid, sub)

	if err != nil {
		return nil, errors, errorPkg.ErrBadRequest.WithError(entity.ErrLoginFailed.Error()).WithDebug(err.Error())
	}

	return &entity.TokenResponse{
		TokenType: "Bearer",
		AccessToken: entity.Token{
			Token:     token.AccessToken,
			ExpiredIn: token.AccessTokenExpired,
		},
		RefreshToken: entity.Token{
			Token:     token.RefreshToken,
			ExpiredIn: token.RefreshTokenExpired,
		},
	}, errors, nil

}
