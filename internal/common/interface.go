package common

import (
	"context"

	"github.com/nhatth/api-service/pkg/jwt"
)

type JWTProvider interface {
	IssueToken(ctx context.Context, id, sub string) (*jwt.TokenDetails, error)
}
