package jwt

import "github.com/dgrijalva/jwt-go"

type AccountTokenClaims struct {
	Level string `json:"lvl"`
	Type  string `json:"typ"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type RoleTokenClaims struct {
	jwt.StandardClaims
}
