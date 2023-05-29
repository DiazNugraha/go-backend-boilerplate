package config

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY = []byte("secretapp")

type JWTClaim struct {
	Email string
	jwt.RegisteredClaims
}
