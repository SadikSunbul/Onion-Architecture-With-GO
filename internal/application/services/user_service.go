package services

import (
	"github.com/SadikSunbul/Onion-Architecture-With-GO/internal/application/commands"
	"github.com/SadikSunbul/Onion-Architecture-With-GO/internal/application/queries"
	"github.com/SadikSunbul/Onion-Architecture-With-GO/internal/domain"
	"github.com/SadikSunbul/Onion-Architecture-With-GO/internal/infrastructure/repositories"
)

// UserService kullanıcı islemlerini yapar
type UserService struct {
	CreateUserCommand *commands.CreateUserCommand // kullanıcı bilgilerini veritabanına kaydeder
	GetUserQuery      *queries.GetUserQuery       // kullanıcı bilgilerini veritabanından alır
}

// NewUserService kullanıcı islemlerini yapar
func NewUserService() *UserService {
	repo := &repositories.UserRepository{} // kullanıcı bilgilerini veritabanına kaydeder
	return &UserService{
		CreateUserCommand: &commands.CreateUserCommand{UserRepository: repo}, // kullanıcı bilgilerini veritabanına kaydeder
		GetUserQuery:      &queries.GetUserQuery{UserRepository: repo},       // kullanıcı bilgilerini veritabanından alır
	}
}

// CreateUser kullanıcı bilgilerini veritabanına kaydeder
func (s *UserService) CreateUser(user *domain.User) error {
	return s.CreateUserCommand.Handle(user) // kullanıcı bilgilerini veritabanına kaydeder
}

// GetUser kullanıcı bilgilerini veritabanından alır
func (s *UserService) GetUser(id int) (*domain.User, error) {
	return s.GetUserQuery.Handle(id) // kullanıcı bilgilerini veritabanından alır
}
