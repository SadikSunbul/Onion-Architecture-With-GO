package main

import (
	"github.com/SadikSunbul/Onion-Architecture-With-GO/internal/application/services"
	"github.com/SadikSunbul/Onion-Architecture-With-GO/internal/interfaces/handlers"
	"github.com/SadikSunbul/Onion-Architecture-With-GO/internal/interfaces/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

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
