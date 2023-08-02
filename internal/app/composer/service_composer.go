package composer

import (
	"net/http"

	"github.com/nhatth/api-service/internal/app/helpers"
	AuthBussines "github.com/nhatth/api-service/internal/app/services/auth/business"
	authSQLStore "github.com/nhatth/api-service/internal/app/services/auth/storage/mysql"
	authAPI "github.com/nhatth/api-service/internal/app/services/auth/transport/api"
	"github.com/nhatth/api-service/internal/common"
	"github.com/nhatth/api-service/pkg/jwt"
	"gorm.io/gorm"
)

type AuthService interface {
	RegisterHdl(w http.ResponseWriter, r *http.Request)
	LoginHdl(w http.ResponseWriter, r *http.Request)
	RefreshTokenHl(w http.ResponseWriter, r *http.Request)
}

func ComposeAuthAPIService(db *gorm.DB) AuthService {

	authStore := authSQLStore.NewMySQLStore(db)

	hasher := new(helpers.Hasher)

	jwtComp := jwt.NewJWT(common.JWTstring)

	jwtComp.InitFlags()

	bus := AuthBussines.NewAuthBusiness(authStore, hasher, jwtComp)

	sericeAPI := authAPI.NewAuthAPI(bus)

	return sericeAPI
}
