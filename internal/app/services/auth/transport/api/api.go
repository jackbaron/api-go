package api

import (
	"context"
	"net/http"

	"github.com/nhatth/api-service/internal/app/helpers"
	"github.com/nhatth/api-service/internal/app/services/auth/entity"
)

type Bussines interface {
	Register(ctx context.Context, data *entity.AuthRegister) error
}

type api struct {
	bussines Bussines
}

func NewAuthAPI(bussines Bussines) *api {
	return &api{bussines: bussines}
}

func (api *api) RegisterHdl(w http.ResponseWriter, r *http.Request) {
	helpers.SendMessageSuccessWithOutPayLoad(w, r, 200)
}
