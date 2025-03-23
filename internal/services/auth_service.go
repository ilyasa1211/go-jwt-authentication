package services

import (
	"encoding/json"
	"net/http"

	"github.com/ilyasa1211/go-jwt-authentication/internal/dto"
	"github.com/ilyasa1211/go-jwt-authentication/internal/repositories"
	"github.com/ilyasa1211/go-jwt-authentication/internal/utils"
)

type AuthService struct {
	r repositories.UserRepository
}

func NewAuthService(r repositories.UserRepository) *AuthService {
	return &AuthService{r}
}

func (s *AuthService) Login(r *http.Request) (string, error) {
	var dto dto.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return "", err
	}

	user := s.r.FindByEmail(dto.Email)

	if err := utils.ComparePass(dto.Password, user.Password); err != nil {
		return "", err
	}

	return utils.GenJWTToken(user), nil
}

func (s *AuthService) Register(r *http.Request) (string, error) {
	var reg dto.RegisterRequest
	var userCreate dto.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&reg); err != nil {
		return "", err
	}

	userCreate = reg.CreateUserRequest

	if pass, err := utils.HashPass(reg.Password); err != nil {
		return "", err
	} else {
		userCreate.Password = pass
	}

	if err := s.r.Create(&userCreate); err != nil {
		return "", err
	}

	user := s.r.FindByEmail(reg.Email)

	if err := utils.ComparePass(reg.Password, user.Password); err != nil {
		return "", err
	}

	return utils.GenJWTToken(user), nil
}
