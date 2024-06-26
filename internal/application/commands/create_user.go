package commands

import (
	"github.com/SadikSunbul/Onion-Architecture-With-GO/internal/domain"
	"github.com/SadikSunbul/Onion-Architecture-With-GO/internal/infrastructure/repositories"
)

// CreateUserCommand veritabanına kullanıcı bilgilerini kaydeder
type CreateUserCommand struct {
	UserRepository *repositories.UserRepository // kullanıcı bilgilerini veritabanına kaydeder
}

// Handle veritabanına kullanıcı bilgilerini kaydeder
func (c *CreateUserCommand) Handle(user *domain.User) error {
	return c.UserRepository.CreateUser(user) // veritabanına kullanıcı bilgilerini kaydeder
}
