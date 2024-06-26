package repositories

import (
	"github.com/SadikSunbul/Onion-Architecture-With-GO/internal/domain"
	"github.com/SadikSunbul/Onion-Architecture-With-GO/internal/infrastructure/database"
)

type UserRepository struct{}

func (r *UserRepository) GetUserByID(id int) (*domain.User, error) {
	var user domain.User
	query := "SELECT id, name, email FROM users WHERE id = $1"
	err := database.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateUser(user *domain.User) error {
	query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
	err := database.DB.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}
