package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func ValidateJwtAccountAccessToken(tokenString string) (jwt.MapClaims, error) {
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

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidToken{tokenString}
	}

	correctIssuer := claims.VerifyIssuer("departamento", true)
	createdBeforeUse := claims.VerifyIssuedAt(time.Now().Unix(), true)
	notExpired := claims.VerifyExpiresAt(time.Now().Unix(), true)
	isAccountToken := claims["typ"] == "acc"
	if !correctIssuer || !createdBeforeUse || !notExpired || !isAccountToken {
		return nil, ErrInvalidToken{tokenString}
	}

	return claims, err
}

func ValidateJwtAccountRefreshToken(tokenString string) (jwt.MapClaims, error) {
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

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken{tokenString}
	}

	correctIssuer := claims.VerifyIssuer("departamento", true)
	createdBeforeUse := claims.VerifyIssuedAt(time.Now().Unix(), true)
	notExpired := claims.VerifyExpiresAt(time.Now().Unix(), true)
	isRefreshToken := claims["typ"] == "ref"
	if !correctIssuer || !createdBeforeUse || !notExpired || !isRefreshToken {
		return nil, ErrInvalidToken{tokenString}
	}

	return claims, err
}
