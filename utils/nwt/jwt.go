package nwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"strings"
	"time"
)

var (
	ErrNilBrowserId  = errors.New("browser id is nil")
	ErrNilUserId     = errors.New("user id is nil")
	ErrNilIssuer     = errors.New("issuer is nil")
	ErrInvalidIssuer = errors.New("invalid issuer")
)

type CustomClaims struct {
	jwt.RegisteredClaims
	UserId    int64  `json:"uid,omitempty"`
	BrowserId string `json:"bid,omitempty"`
}

func (c CustomClaims) Validate() error {
	if len(c.BrowserId) <= 0 {
		return ErrNilBrowserId
	}

	if c.UserId <= 0 {
		return ErrNilUserId
	}

	if len(c.Issuer) <= 0 {
		return ErrNilIssuer
	}

	if !strings.EqualFold(c.Issuer, os.Getenv("JWT_ISSUER")) {
		return ErrInvalidIssuer
	}

	return nil
}

func NewTokenWithStandardClaims(secret string, browserId string, id int64, expires time.Time) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    os.Getenv("JWT_ISSUER"),
			ExpiresAt: jwt.NewNumericDate(expires),
		},
		UserId:    id,
		BrowserId: browserId,
	}).SignedString([]byte(secret))
}

func ParseErrorCheck(secret string, tokenString string, browserId string) (issuer string, userId int64, err error) {
	claims := CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", 0, err
	}

	if token.Valid {
		if claims.UserId <= 0 {
			return "", 0, jwt.NewValidationError("invalid claims: user id is 0", jwt.ValidationErrorClaimsInvalid)
		}

		if len(claims.Issuer) <= 0 {
			return "", 0, jwt.NewValidationError("invalid claims: issuer is null", jwt.ValidationErrorClaimsInvalid)
		}

		if len(claims.BrowserId) <= 0 {
			return "", 0, jwt.NewValidationError("invalid claims: browser id is null", jwt.ValidationErrorClaimsInvalid)
		}

		return claims.Issuer, claims.UserId, nil
	} else {
		return "", 0, jwt.NewValidationError("invalid claims", jwt.ValidationErrorClaimsInvalid)
	}
}
