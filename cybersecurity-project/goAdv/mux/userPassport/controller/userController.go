package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"userPassport/model"
	"userPassport/repository"
	"userPassport/service"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

var repo1 repository.Repository
var userService *service.UserService
var passportService *service.PassportService
var db *gorm.DB
var users = map[string]string{"naren": "passme", "admin": "password"}
var secretKey = []byte("yogesh")
var logger *zerolog.Logger

// Response is a representation of JSON response for JWT
type Response struct {
	Token  string `json:"token"`
	Status string `json:"status"`
}

func ControllerInit(dB *gorm.DB, loggerins *zerolog.Logger) {
	logger = loggerins
	db = dB
	repo1 = repository.NewRepository()
	userService = service.NewUserService(repo1, db)
	passportService = service.NewPassportService(repo1, db)
}

func RegisterRoutes(router *mux.Router) {
	router.Use(CheckAuth)
	subRoute := router.PathPrefix("/users").Subrouter()
	subRoute.HandleFunc("/", GetUsers).Methods("GET")
	subRoute.HandleFunc("/", GetUsers).Queries("limit", "{limit:[0-9]+}", "pageNo", "{pageNo:[0-9]+}").Methods("GET")
	subRoute.HandleFunc("/", CreateUser).Methods("POST")
	subRoute.HandleFunc("/{id}", UpdateUser).Methods("PUT")
	subRoute.HandleFunc("/{id}", GetUser).Methods("GET")
	subRoute.HandleFunc("/{id}", DeleteUser).Methods("DELETE")
	subRoute.HandleFunc("/{id}/passports", GetPassportByUserId).Methods("GET")
}

// Middleware function, which will be called for each request
func CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info().Msg(r.URL.Path)
		if r.URL.Path == "/login" {
			next.ServeHTTP(w, r)
			return
		}
		tokenString, err := request.HeaderExtractor{"access_token"}.ExtractToken(r)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return secretKey, nil
		})
		log.Println(reflect.TypeOf(token))
		if err != nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// If token is valid
			// We found the token in our map
			log.Printf("Authenticated user %s\n", claims)

			// Pass down the request to the next middleware (or final handler)
			next.ServeHTTP(w, r)

		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

// LoginHandler validates the user credentials
func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
		return
	}
	log.Println(r.Form)
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	log.Println("username ", username)
	log.Println("password", password)
	if originalPassword, ok := users[username]; ok {
		if password == originalPassword {
			// Create a claims map
			claims := jwt.MapClaims{
				"username": username,
				"IssuedAt": time.Now().Unix(),
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(secretKey)
			if err != nil {
				w.WriteHeader(http.StatusBadGateway)
				w.Write([]byte(err.Error()))
			}
			response := Response{Token: tokenString, Status: "success"}
			responseJSON, _ := json.Marshal(response)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write(responseJSON)

		} else {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			return
		}
	} else {
		http.Error(w, "User is not found", http.StatusNotFound)
		return
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount()))
	limit, _ := strconv.Atoi(r.FormValue("limit"))   /// baseurl/?limit=12&pageno=2
	pageNo, _ := strconv.Atoi(r.FormValue("pageNo")) //12 limit-3 pageno-2   //4,5,6
	// filter := r.FormValue("filter")
	// s := strings.Split(filter, ",")
	// fmt.Println(s)
	params := r.URL.Query()      // baseurl/hobby=ghf&hobby=trd&hobby=tre
	fmt.Println(params["hobby"]) //ghf,trd,tre
	offset := limit * (pageNo - 1)
	fmt.Println(limit, pageNo)
	var users []model.User
	userService.GetAllUsers(&users, limit, offset)
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount()))
	params := mux.Vars(r)
	id, _ := uuid.FromString(params["id"])
	var str []string
	str = append(str, "Passport")
	var user model.User
	userService.GetUserById(&user, id, str)
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount()))
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	userService.CreateUser(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount()))
	params := mux.Vars(r)
	id, _ := uuid.FromString(params["id"])
	var updatedUser model.User
	updatedUser.ID = id
	json.NewDecoder(r.Body).Decode(&updatedUser)
	userService.UpdateUser(updatedUser)
	json.NewEncoder(w).Encode(updatedUser)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount()))
	params := mux.Vars(r)
	idToDelete, _ := uuid.FromString(params["id"])
	userService.DeleteUser(idToDelete)
	w.WriteHeader(http.StatusNoContent)
}

func GetPassportByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount()))
	params := mux.Vars(r)
	id, _ := uuid.FromString(params["id"])
	var passport model.Passport
	passportService.GetPassportByUserId(&passport, id)
	json.NewEncoder(w).Encode(passport)
}
