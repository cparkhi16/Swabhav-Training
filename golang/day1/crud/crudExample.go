package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var db *gorm.DB

func initDB() {
	dsn := "root:yourpassword@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.AutoMigrate(&User{})
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	db.Create(&user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	db.Find(&users)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	if db.First(&user, params["id"]).RecordNotFound() {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	if db.First(&user, params["id"]).RecordNotFound() {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	db.Save(&user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	if db.First(&user, params["id"]).RecordNotFound() {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	db.Delete(&user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}
func isAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
}
func main() {
	initDB()
	r := mux.NewRouter()
	r.Use(isAuthorized)
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	fmt.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
