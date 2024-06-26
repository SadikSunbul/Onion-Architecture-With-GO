package handlers

import (
	"encoding/json"
	"github.com/SadikSunbul/Onion-Architecture-With-GO/internal/application/services"
	"github.com/SadikSunbul/Onion-Architecture-With-GO/internal/domain"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// UserHandler kullanıcı islemlerini yapar
type UserHandler struct {
	UserService *services.UserService // kullanıcı islemlerini yapar
}

// NewUserHandler kullanıcı islemlerini yapar
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User                         // kullanıcı
	err := json.NewDecoder(r.Body).Decode(&user) // json ile kullanıcı bilgilerini alır
	if err != nil {                              // hata varsa
		http.Error(w, err.Error(), http.StatusBadRequest) // hata döndürür
		return
	}

	err = h.UserService.CreateUser(&user) // kullanıcı bilgilerini veritabanına kaydeder
	if err != nil {                       // hata varsa
		http.Error(w, err.Error(), http.StatusInternalServerError) // hata döndürür
		return
	}

	w.Header().Set("Content-Type", "application/json") // json olarak döndürür
	json.NewEncoder(w).Encode(user)                    // kullanıcı bilgileri döndürür
}

// GetUser kullanıcı bilgilerini veritabanından alır
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                 // url parametrelerini alır
	id, err := strconv.Atoi(vars["id"]) // url parametrelerinden idyi alır
	if err != nil {                     // hata varsa
		http.Error(w, err.Error(), http.StatusBadRequest) // hata döndürür
		return
	}

	user, err := h.UserService.GetUser(id) // kullanıcı bilgilerini veritabanından alır
	if err != nil {                        // hata varsa
		http.Error(w, err.Error(), http.StatusInternalServerError) // hata döndürür
		return                                                     // hata döndürür
	}

	w.Header().Set("Content-Type", "application/json") // json olarak döndürür
	json.NewEncoder(w).Encode(user)                    // kullanıcı bilgileri döndürür
}
