package commands

import (
	"github.com/SadikSunbul/Onion-Architecture-With-GO/internal/domain"
	"github.com/SadikSunbul/Onion-Architecture-With-GO/internal/infrastructure/repositories"
)

type CreateUserCommand struct {
	UserRepository *repositories.UserRepository
}

func (c *CreateUserCommand) Handle(user *domain.User) error {
	return c.UserRepository.CreateUser(user)
}
