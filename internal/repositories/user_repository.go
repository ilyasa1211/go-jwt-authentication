package repositories

import (
	"github.com/ilyasa1211/go-jwt-authentication/internal/dto"
	"github.com/ilyasa1211/go-jwt-authentication/internal/entities"
)

type UserRepository interface {
	FindAll() []*entities.User
	FindById(id string) *entities.User
	FindByEmail(email string) *entities.User
	Create(u *dto.CreateUserRequest) error
	UpdateById(id string, u *dto.UpdateUserRequest) error
	DeleteById(id string) error
}
