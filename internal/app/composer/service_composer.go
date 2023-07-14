package composer

import (
	"net/http"

	AuthBussines "github.com/nhatth/api-service/internal/app/services/auth/business"
	authSQLStore "github.com/nhatth/api-service/internal/app/services/auth/storage/mysql"
	authAPI "github.com/nhatth/api-service/internal/app/services/auth/transport/api"
	"gorm.io/gorm"
)

type AuthService interface {
	RegisterHdl(w http.ResponseWriter, r *http.Request)
}

func ComposeAuthAPIService(db *gorm.DB) AuthService {

	authStore := authSQLStore.NewMySQLStore(db)

	bus := AuthBussines.NewAuthBusiness(authStore)

	sericeAPI := authAPI.NewAuthAPI(bus)

	return sericeAPI
}
