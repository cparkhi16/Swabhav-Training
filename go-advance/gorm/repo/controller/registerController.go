package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var mySigningKey = []byte("captainjacksparrowsayshi")

func RegisterRoutesForUser(uc *UserController, router *mux.Router) {
	router.Use(isAuthorized)
	router.HandleFunc("/user/passport", GetTestTwo).Methods("GET")
	router.HandleFunc("/user/{id}", GetTestOne).Methods("GET")
	router.HandleFunc("/users/{id}/passport", uc.GetUserPassport).Methods("GET")
	router.HandleFunc("/users/token", uc.GetUserToken).Methods("GET")
	router.HandleFunc("/users/{id}/passport", uc.UpdateUserPassportDetail).Methods("PUT")
	router.HandleFunc("/users/{id}/passport", uc.AddPassportForUser).Methods("POST")
	router.HandleFunc("/users/{id}/passport", uc.DeletePassportDetailsForUser).Methods("DELETE")
	router.HandleFunc("/users", uc.GetAllUsers).Methods("GET")
	router.HandleFunc("/users", uc.CreateUser).Methods("POST")
}

func isAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/users/token" {
			next.ServeHTTP(w, r)
			return
		}
		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				next.ServeHTTP(w, r)
			}
		} else {

			fmt.Fprintf(w, "Not Authorized")
		}
	})
}
func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "Chinmay Parkhi"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
