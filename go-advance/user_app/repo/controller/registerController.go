package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var mySigningKey = []byte("captainjacksparrowsayshi")

func (uc *UserController) RegisterRoutesForUser(authRouter *mux.Router, nonAuthRouter *mux.Router) {
	nonAuthRouter.HandleFunc("/login", uc.GetUserToken).Methods("POST")
	nonAuthRouter.HandleFunc("/validateToken", uc.ValidateToken).Methods("POST")
	nonAuthRouter.HandleFunc("/users", uc.CreateUser).Methods("POST")
	nonAuthRouter.HandleFunc("/users/{id}/hobbies", uc.GetAllUserHobbies).Methods("GET")
	nonAuthRouter.HandleFunc("/users/{id}", uc.GetUser).Methods("GET")
	authRouter.Use(isAuthorized)
	authRouter.HandleFunc("/users/{id}/passport", uc.GetUserPassport).Methods("GET")
	authRouter.HandleFunc("/users/{id}/hobbies", uc.AddHobbiesForUser).Methods("PUT")
	authRouter.HandleFunc("/users/{id}/hobbies", uc.DeleteHobbiesForUser).Methods("DELETE")
	authRouter.HandleFunc("/users/{id}/passport", uc.UpdateUserPassportDetail).Methods("PUT")
	authRouter.HandleFunc("/users/{id}/passport", uc.AddPassportForUser).Methods("POST")
	authRouter.HandleFunc("/users/{id}/passport", uc.DeletePassportDetailsForUser).Methods("DELETE")
	authRouter.HandleFunc("/users", uc.GetAllUsers).Methods("GET")
	authRouter.HandleFunc("/users/{id}", uc.DeleteUser).Methods("DELETE")
	authRouter.HandleFunc("/users/{id}", uc.UpdateUser).Methods("PUT")
}
func (cc *CourseController) RegisterRoutesForCourse(authRouter *mux.Router, nonAuthRouter *mux.Router) {
	nonAuthRouter.HandleFunc("/courses", cc.GetAllCourses).Methods("GET")
	authRouter.Use(isAuthorized)
	authRouter.HandleFunc("/courses", cc.CreateCourse).Methods("POST")
	authRouter.HandleFunc("/courses/{id}", cc.DeleteCourse).Methods("DELETE")
	authRouter.HandleFunc("/courses/{id}", cc.UpdateCourse).Methods("PUT")
}
func (hc *HobbyController) RegisterRoutesForHobby(authRouter *mux.Router, nonAuthRouter *mux.Router) {
	nonAuthRouter.HandleFunc("/hobbies", hc.GetAllHobbies).Methods("GET")
	authRouter.Use(isAuthorized)
	authRouter.HandleFunc("/hobbies/{id}", hc.GetHobby).Methods("GET")
	authRouter.HandleFunc("/hobbies/{id}", hc.DeleteHobby).Methods("DELETE")
	authRouter.HandleFunc("/hobbies/{id}", hc.UpdateHobby).Methods("PUT")
}

//client id
//237013399860-4rjcn57jg6md4f35r86919linl8pt2eq.apps.googleusercontent.com

//secret
//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
func isAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
func generateJWT() (string, error) {
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
