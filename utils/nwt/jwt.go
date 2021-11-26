package nwt

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	UserId    int64  `json:"uid,omitempty"`
	BrowserId string `json:"bid,omitempty"`
}

func NewTokenWithStandardClaims(secret string, browserId string, id int64, expires time.Time) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "link.nanoit.kr",
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
