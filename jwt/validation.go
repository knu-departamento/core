package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func ValidateJwtAccountAccessToken(tokenString string) (*AccountTokenClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_SIGNING_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*AccountTokenClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	if !(!claims.VerifyIssuer("departamento", true) || time.Now().Unix() <= claims.IssuedAt || time.Now().Unix() >= claims.ExpiresAt || claims.Level != "account" || claims.Type != "access") {
		return nil, errors.New("invalid token")
	}

	return claims, err
}

func ValidateJwtAccountRefreshToken(tokenString string) (*AccountTokenClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_SIGNING_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(AccountTokenClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	if !(!claims.VerifyIssuer("departamento", true) || time.Now().Unix() <= claims.IssuedAt || time.Now().Unix() >= claims.ExpiresAt || claims.Level != "account" || claims.Type != "refresh") {
		return nil, errors.New("invalid token")
	}

	return &claims, err
}
