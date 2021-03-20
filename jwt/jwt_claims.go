package jwt

import "github.com/dgrijalva/jwt-go"

type AccountTokenClaims struct {
	Type  string `json:"typ"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type RoleTokenClaims struct {
	Type         string `json:"typ"`
	Email        string `json:"email"`
	Role         string `json:"rol"`
	DepartmentId int64    `json:"dep_id"`
	jwt.StandardClaims
}
