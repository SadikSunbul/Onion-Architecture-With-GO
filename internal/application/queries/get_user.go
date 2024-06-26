package queries

import (
	"github.com/SadikSunbul/Onion-Architecture-With-GO/internal/domain"
	"github.com/SadikSunbul/Onion-Architecture-With-GO/internal/infrastructure/repositories"
)

// GetUserQuery veritabanından kullanıcı bilgilerini getirir
type GetUserQuery struct {
	UserRepository *repositories.UserRepository // kullanıcı bilgilerini veritabanından alır
}

// Handle veritabanından kullanıcı bilgilerini getirir
func (q *GetUserQuery) Handle(id int) (*domain.User, error) {
	return q.UserRepository.GetUserByID(id) // veritabanından kullanıcı bilgilerini getirir
}
