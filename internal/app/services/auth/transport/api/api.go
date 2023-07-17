package api

import (
	"context"
	"net/http"

	"github.com/nhatth/api-service/internal/app/helpers"
	"github.com/nhatth/api-service/internal/app/services/auth/entity"
)

type Bussines interface {
	Register(ctx context.Context, data *entity.AuthRegister) (map[string]string, error)
}

type api struct {
	bussines Bussines
}

func NewAuthAPI(bussines Bussines) *api {
	return &api{bussines: bussines}
}

func (api *api) RegisterHdl(w http.ResponseWriter, r *http.Request) {
	var data entity.AuthRegister

	helpers.BindingDataBody(r, &data)

	msgErros, err := api.bussines.Register(r.Context(), &data)

	if err != nil {

		helpers.SendErrorResponse(w, r, http.StatusBadRequest, "Failed to validation", msgErros)
		return
	}

	helpers.SendMessageSuccessWithOutPayLoad(w, r, http.StatusOK)
}
