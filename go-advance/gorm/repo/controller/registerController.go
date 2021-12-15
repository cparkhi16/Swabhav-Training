package controller

import (
	s "app/service"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func RegisterRoutesForUser(userService *s.UserService, router *mux.Router) {
	router.HandleFunc("/users/{id}/passport", GetUserPassport(userService)).Methods("GET")
	router.HandleFunc("/users", GetAllUsers(userService)).Methods("GET")
	log.Fatal(http.ListenAndServe(":9000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"POST", "PUT", "DELETE"}), handlers.AllowedOrigins([]string{"abc.com"}))(router)))
}
