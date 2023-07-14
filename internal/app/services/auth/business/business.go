package bussines

import (
	"context"

	"github.com/nhatth/api-service/internal/app/services/auth/entity"
)

type AuthRepository interface {
	AddNewUser(ctx context.Context, data *entity.AuthRegister) error
	GetUser(ctx context.Context, email string) (*entity.AuthUser, error)
}

type business struct {
	repository AuthRepository
}

func NewAuthBusiness(repository AuthRepository) *business {
	return &business{repository: repository}
}

func (bus *business) Register(ctx context.Context, data *entity.AuthRegister) error {
	return nil
}
