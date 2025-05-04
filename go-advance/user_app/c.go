package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID    int    `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var db *gorm.DB

func initDB() {
	dbConn := "root:hello@(127.0.0.1:3306)/testDb"
	db, err := gorm.Open("mysql", dbConn)
	if err != nil {
		log.Fatal("err")
	}
	db.AutoMigrate(&User{})
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	m := mux.Vars(r)
	var id string = m["id"]
	var u User
	if db.First(&u, id).RecordNotFound() {
		http.Error(w, "No user found ", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

func isAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
	})
}
func main() {
	r := mux.NewRouter()
	r.Use(isAuthorized)
	r.HandleFunc("/users/{id}", getUserById)
}
