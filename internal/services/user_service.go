package services

import (
	"encoding/json"
	"net/http"

	"github.com/ilyasa1211/go-jwt-authentication/internal/dto"
	"github.com/ilyasa1211/go-jwt-authentication/internal/entities"
	"github.com/ilyasa1211/go-jwt-authentication/internal/repositories"
	"github.com/ilyasa1211/go-jwt-authentication/internal/utils"
)

type UserService struct {
	r repositories.UserRepository
}

func NewUserService(r repositories.UserRepository) *UserService {
	return &UserService{r}
}

func (s *UserService) FindAll() []*entities.User {
	return s.r.FindAll()
}
func (s *UserService) FindById(r *http.Request) *entities.User {
	id := r.PathValue("id")

	return s.r.FindById(id)
}
func (s *UserService) Create(r *http.Request) error {
	var dto dto.CreateUserRequest

	json.NewDecoder(r.Body).Decode(&dto)

	if pass, err := utils.HashPass(dto.Password); err != nil {
		return err
	} else {
		dto.Password = pass
	}

	return s.r.Create(&dto)
}
func (s *UserService) UpdateById(r *http.Request) error {
	id := r.PathValue("id")

	var dto dto.UpdateUserRequest
	json.NewDecoder(r.Body).Decode(&dto)

	return s.r.UpdateById(id, &dto)
}
func (s *UserService) DeleteById(r *http.Request) error {
	id := r.PathValue("id")

	return s.r.DeleteById(id)
}
