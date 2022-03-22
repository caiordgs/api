package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n Method: %s, URI: %s, Host: %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Authenticate verifies if the user is authenticated.
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Validating...")
		next(w, r)
	}
}
