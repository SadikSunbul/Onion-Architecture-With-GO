package queries

import (
	"github.com/SadikSunbul/Onion-Architecture-With-GO/internal/domain"
	"github.com/SadikSunbul/Onion-Architecture-With-GO/internal/infrastructure/repositories"
)

type GetUserQuery struct {
	UserRepository *repositories.UserRepository
}

func (q *GetUserQuery) Handle(id int) (*domain.User, error) {
	return q.UserRepository.GetUserByID(id)
}
