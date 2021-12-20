package controller

import (
	lr "app/logger"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var mySigningKey = []byte("captainjacksparrowsayshi")
var logger = lr.GetLogger()

func RegisterRoutesForUser(uc *UserController, router *mux.Router) {
	router.Use(isAuthorized)
	router.HandleFunc("/users/{id}/passport", uc.GetUserPassport).Methods("GET")
	router.HandleFunc("/users/{id}/hobbies", uc.GetAllUserHobbies).Methods("GET")
	router.HandleFunc("/users/{id}/hobbies", uc.AddHobbiesForUser).Methods("PUT")
	router.HandleFunc("/users/{id}/hobbies", uc.DeleteHobbiesForUser).Methods("DELETE")
	router.HandleFunc("/users/token", uc.GetUserToken).Methods("GET")
	router.HandleFunc("/users/{id}/passport", uc.UpdateUserPassportDetail).Methods("PUT")
	router.HandleFunc("/users/{id}/passport", uc.AddPassportForUser).Methods("POST")
	router.HandleFunc("/users/{id}/passport", uc.DeletePassportDetailsForUser).Methods("DELETE")
	router.HandleFunc("/users", uc.GetAllUsers).Methods("GET")
	router.HandleFunc("/users", uc.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}/", uc.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/{id}/", uc.UpdateUser).Methods("PUT")
}
func RegisterRoutesForCourse(cc *CourseController, router *mux.Router) {
	router.Use(isAuthorized)
	router.HandleFunc("/courses", cc.GetAllCourses).Methods("GET")
	router.HandleFunc("/courses", cc.CreateCourse).Methods("POST")
	router.HandleFunc("/course/{id}", cc.DeleteCourse).Methods("DELETE")
	router.HandleFunc("/course/{id}", cc.UpdateCourse).Methods("PUT")
}
func RegisterRoutesForHobby(hc *HobbyController, router *mux.Router) {
	router.Use(isAuthorized)
	router.HandleFunc("/hobby/{id}", hc.GetHobby).Methods("GET")
	router.HandleFunc("/hobbies", hc.GetAllHobbies).Methods("GET")
	router.HandleFunc("/hobby/{id}", hc.DeleteHobby).Methods("DELETE")
	router.HandleFunc("/hobby/{id}", hc.UpdateHobby).Methods("PUT")
}

//client id
//237013399860-4rjcn57jg6md4f35r86919linl8pt2eq.apps.googleusercontent.com

//secret
//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
func isAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		if r.URL.Path == "/users/token" || (r.URL.Path == "/users" && r.Method == "POST") || r.URL.Path == "/users/"+params["id"]+"/hobbies" ||
			r.URL.Path == "/courses" && r.Method == "GET" || r.URL.Path == "/hobbies" {
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
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	fmt.Println("Claims for authenticated user ", claims)
	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
