package controller

import (
	s "app/service"

	"github.com/gorilla/mux"
)

func RegisterRoutesForUser(userService *s.UserService, router *mux.Router) {
	router.HandleFunc("/users/{id}/passport", GetUserPassport(userService)).Methods("GET")
	router.HandleFunc("/users/{id}/passport", UpdateUserPassportDetail(userService)).Methods("PUT")
	router.HandleFunc("/users/{id}/passport", AddPassportForUser(userService)).Methods("POST")
	router.HandleFunc("/users/{id}/passport", DeletePassportDetailsForUser(userService)).Methods("DELETE")
	router.HandleFunc("/users", GetAllUsers(userService)).Methods("GET")
	router.HandleFunc("/users", CreateUser(userService)).Methods("POST")
}
