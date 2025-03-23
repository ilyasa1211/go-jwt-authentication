package repositories

import (
	"database/sql"

	"github.com/ilyasa1211/go-jwt-authentication/internal/dto"
	"github.com/ilyasa1211/go-jwt-authentication/internal/entities"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) FindAll() []*entities.User {
	rows, err := r.Db.Query("SELECT * FROM users")

	if err != nil {
		panic(err)
	}

	var users []*entities.User

	for rows.Next() {
		var user entities.User
		rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)

		users = append(users, &user)
	}

	return users
}
func (r *UserRepository) FindById(id string) *entities.User {
	rows, err := r.Db.Query("SELECT * FROM users WHERE id = ?", id)

	if err != nil {
		panic(err)
	}

	var user *entities.User

	if rows.Next() {
		rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	}

	return user
}

func (r *UserRepository) FindByEmail(email string) *entities.User {
	rows, err := r.Db.Query("SELECT * FROM users WHERE email = ?", email)

	if err != nil {
		panic(err)
	}

	var user entities.User

	if rows.Next() {
		rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	}

	return &user
}

func (r *UserRepository) Create(u *dto.CreateUserRequest) error {
	if _, err := r.Db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", u.Name, u.Email, u.Password); err != nil {
		return err
	}

	return nil
}
func (r *UserRepository) UpdateById(id string, u *dto.UpdateUserRequest) error {
	if _, err := r.Db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", u.Name, u.Email, id); err != nil {
		return err
	}

	return nil
}
func (r *UserRepository) DeleteById(id string) error {
	if _, err := r.Db.Exec("DELETE FROM users WHERE id = ?", id); err != nil {
		return err
	}

	return nil
}
