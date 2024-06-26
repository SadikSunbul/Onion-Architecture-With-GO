package repositories

import (
	"github.com/SadikSunbul/Onion-Architecture-With-GO/internal/domain"
	"github.com/SadikSunbul/Onion-Architecture-With-GO/internal/infrastructure/database"
)

type UserRepository struct{}

// GetUserByID veritabanından kullanıcı bilgilerini getirir
func (r *UserRepository) GetUserByID(id int) (*domain.User, error) {
	var user domain.User                                                           // kullanıcı
	query := "SELECT id, name, email FROM users WHERE id = $1"                     // veritabanına gönderilcek sorgu
	err := database.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email) // veritabanından veri alınır
	if err != nil {                                                                // hata varsa
		return nil, err // hata döndürür
	}
	return &user, nil // kullanıcı bilgileri döndürür
}

// CreateUser veritabanına kullanıcı bilgilerini kaydeder
func (r *UserRepository) CreateUser(user *domain.User) error {
	query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"  // veritabanına gönderilecek sorgu
	err := database.DB.QueryRow(query, user.Name, user.Email).Scan(&user.ID) // veritabanına sorgu gonderilir
	if err != nil {                                                          // hata varsa
		return err // hata döndürür
	}
	return nil // hata yoksa
}
