package common

import (
	"context"

	jwtGolang "github.com/golang-jwt/jwt/v5"
	"github.com/nhatth/api-service/pkg/jwt"
)

type JWTProvider interface {
	IssueToken(ctx context.Context, id, sub string) (*jwt.TokenDetails, error)
	ParseToken(ctx context.Context, tokenStr string) (*jwtGolang.RegisteredClaims, error)
}
