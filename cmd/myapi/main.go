package main

import (
	"github.com/SadikSunbul/Onion-Architecture-With-GO/internal/application/services"
	"github.com/SadikSunbul/Onion-Architecture-With-GO/internal/infrastructure/database"
	"github.com/SadikSunbul/Onion-Architecture-With-GO/internal/interfaces/handlers"
	"github.com/SadikSunbul/Onion-Architecture-With-GO/internal/interfaces/middleware"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
	"log"
	"net/http"
	"net/url"
	"os"
)

// Config yapılandırma dosyası
type Config struct {
	Database struct { // veritabanı bilgileri
		Connection string `yaml:"connection"` // veritabanı bağlantısı
	} `yaml:"database"` // veritabanı bilgileri
}

func main() {

	configFile, err := os.ReadFile("../../config.yaml") // config dosyasını okur
	if err != nil {
		log.Fatalf("config.yaml dosyası okunamadı: %v", err) // hata döndürür
	}

	var config Config                         // yapılandırma dosyası
	err = yaml.Unmarshal(configFile, &config) // yapılandırma dosyasını okur
	if err != nil {                           // hata varsa
		log.Fatalf("Yapılandırma dosyası ayrıştırılamadı: %v", err)
	}

	conn, _ := url.Parse(config.Database.Connection)       // veritabanı bağlantısı
	conn.RawQuery = "sslmode=verify-ca;sslrootcert=ca.pem" // ssl bilgileri
	database.InitDB(conn.String())                         // veritabanını başlatır

	userService := services.NewUserService()                       // kullanıcı islemlerini yapar
	userHandler := &handlers.UserHandler{UserService: userService} // kullanıcı islemlerini yapar

	r := mux.NewRouter()                // router oluşturur
	r.Use(middleware.LoggingMiddleware) // middleware ekle

	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")  // kullanıcı bilgilerini veritabanına kaydeder
	r.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET") // kullanıcı bilgilerini veritabanından alır

	log.Println("Server started at :8080")     // sunucu baslatılır
	log.Fatal(http.ListenAndServe(":8080", r)) // sunucu baslatılır
}
