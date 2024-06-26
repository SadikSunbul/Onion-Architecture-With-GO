package services

import (
	"github.com/SadikSunbul/Onion-Architecture-With-GO/internal/application/commands"
	"github.com/SadikSunbul/Onion-Architecture-With-GO/internal/application/queries"
	"github.com/SadikSunbul/Onion-Architecture-With-GO/internal/domain"
	"github.com/SadikSunbul/Onion-Architecture-With-GO/internal/infrastructure/repositories"
)

type UserService struct {
	CreateUserCommand *commands.CreateUserCommand
	GetUserQuery      *queries.GetUserQuery
}

func NewUserService() *UserService {
	repo := &repositories.UserRepository{}
	return &UserService{
		CreateUserCommand: &commands.CreateUserCommand{UserRepository: repo},
		GetUserQuery:      &queries.GetUserQuery{UserRepository: repo},
	}
}

func (s *UserService) CreateUser(user *domain.User) error {
	return s.CreateUserCommand.Handle(user)
}

func (s *UserService) GetUser(id int) (*domain.User, error) {
	return s.GetUserQuery.Handle(id)
}
