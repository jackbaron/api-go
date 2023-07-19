package jwt

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

const (
	defaultSecret                   = "my-secret"
	defaultTimeSecretExpires        = 60 * 60 * 24 * 7 // 7 days
	defaultRefreshSecret            = "my-refresh-secret"
	defaultTimeRefreshSecretExpires = 60 * 60 * 24 * 8 // 8 days
)

var (
	ErrSecretKeyNotValid     = errors.New("secret key must be in 32 bytes")
	ErrTokenLifeTimeTooShort = errors.New("token life time too short")
)

type TokenDetails struct {
	AccessToken         string
	RefreshToken        string
	AccessTokenExpired  *time.Time
	RefreshTokenExpired *time.Time
}

type jwtx struct {
	id                                string
	secret                            string
	refreshSecret                     string
	expireSecretTokenInSeconds        int
	expireRefreshSecretTokenInSeconds int
}

func NewJWT(id string) *jwtx {
	return &jwtx{id: id}
}

func (j *jwtx) GetID() string {
	return j.id
}

func (j *jwtx) InitFlags() {
	flag.StringVar(&j.secret, "jwt-secret", defaultSecret, "Secret key to sign JWT")
	flag.StringVar(&j.refreshSecret, "jwt-refresh-secret", defaultRefreshSecret, "Refresh secret key to sign JWT")
	flag.IntVar(&j.expireSecretTokenInSeconds, "jwt-exp-secret", defaultTimeSecretExpires, "Number of seconds token will expired")
	flag.IntVar(&j.expireRefreshSecretTokenInSeconds, "jwt-exp-refresh-secret", defaultTimeRefreshSecretExpires, "Number of seconds refresh token will expired")
}

func (j *jwtx) IssueToken(ctx context.Context, id, sub string) (*TokenDetails, error) {
	now := time.Now().UTC()

	tokenDetail := new(TokenDetails)

	accesTokenExpired := jwt.NewNumericDate(now.Add(time.Second * time.Duration(j.expireSecretTokenInSeconds)))
	refrehTokenExpired := jwt.NewNumericDate(now.Add(time.Second * time.Duration(j.expireRefreshSecretTokenInSeconds)))

	//? Generate accesss token
	claims := jwt.RegisteredClaims{
		Subject:   sub,
		ExpiresAt: accesTokenExpired,
		NotBefore: jwt.NewNumericDate(now),
		IssuedAt:  jwt.NewNumericDate(now),
		ID:        id,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenSignedStr, err := t.SignedString([]byte(j.secret))

	if err != nil {

		return tokenDetail, errors.WithStack(err)
	}

	tokenDetail.AccessToken = tokenSignedStr
	tokenDetail.AccessTokenExpired = &accesTokenExpired.Time

	//? Generate refreh token
	claimsRefresh := jwt.RegisteredClaims{
		Subject:   sub,
		ExpiresAt: refrehTokenExpired,
		NotBefore: jwt.NewNumericDate(now),
		IssuedAt:  jwt.NewNumericDate(now),
		ID:        id,
	}

	t = jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)

	refreshTokenSigned, err := t.SignedString([]byte(j.refreshSecret))

	if err != nil {

		return tokenDetail, errors.WithStack(err)
	}

	tokenDetail.RefreshToken = refreshTokenSigned
	tokenDetail.RefreshTokenExpired = &refrehTokenExpired.Time

	return tokenDetail, nil
}

func (j *jwtx) ParseToken(ctx context.Context, tokenStr string) (claims *jwt.RegisteredClaims, err error) {
	var rc jwt.RegisteredClaims

	token, err := jwt.ParseWithClaims(tokenStr, &rc, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.secret), nil
	})

	if !token.Valid {
		return nil, errors.WithStack(err)
	}

	return &rc, nil
}
