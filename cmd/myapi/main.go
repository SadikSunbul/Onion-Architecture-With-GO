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

type Config struct {
	Database struct {
		Connection string `yaml:"connection"`
	} `yaml:"database"`
}

func main() {

	configFile, err := os.ReadFile("../../config.yaml")
	if err != nil {
		log.Fatalf("config.yaml dosyası okunamadı: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("Yapılandırma dosyası ayrıştırılamadı: %v", err)
	}

	conn, _ := url.Parse(config.Database.Connection)
	conn.RawQuery = "sslmode=verify-ca;sslrootcert=ca.pem"
	database.InitDB(conn.String())

	//postgresql database bağlantısını burada sağla değilse çalışmaz
	//database.InitDB("connectionstring")

	userService := services.NewUserService()
	userHandler := &handlers.UserHandler{UserService: userService}

	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)

	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
