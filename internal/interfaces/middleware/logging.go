package middleware

import (
	"log"
	"net/http"
)

// LoggingMiddleware Günlüğe kaydetme için ara yazılım
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { // http.HandlerFunc ile fonksiyonu cagırır
		log.Printf("Request: %s %s", r.Method, r.URL.Path) // istek bilgilerini loga yazdırır
		next.ServeHTTP(w, r)                               // fonksiyonu cagırır
	})
}
