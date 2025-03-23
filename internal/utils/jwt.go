package utils

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ilyasa1211/go-jwt-authentication/internal/configs"
	"github.com/ilyasa1211/go-jwt-authentication/internal/entities"
)

type UserClaim struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func GenJWTToken(u *entities.User) string {
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"id":   u.ID,
		"name": u.Name,
	})

	key := configs.GetJWTSecret()

	s, err := t.SignedString([]byte(key))

	if err != nil {
		panic(err)
	}

	return s
}

func VerifyJWTToken(tokenStr string) (*UserClaim, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(configs.GetJWTSecret()), nil
	}, jwt.WithValidMethods([]string{
		jwt.SigningMethodHS512.Alg(),
	}))

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*UserClaim); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token claims")
}
