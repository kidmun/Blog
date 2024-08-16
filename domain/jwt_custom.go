package domain

import (
	"github.com/dgrijalva/jwt-go"
)

type JwtCustomClaims struct {
	UserName string `json:"userName"`
	ID   string `json:"id"`
	jwt.StandardClaims
}

type JwtCustomRefreshClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}