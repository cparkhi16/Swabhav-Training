package controller

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"
	"userPassport/hashLogic"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"userPassport/model"
	"userPassport/service"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

//var secretKey = []byte("yogesh")

// Response is a representation of JSON response for JWT
type Response struct {
	Token  string `json:"token"`
	Status string `json:"status"`
	UserId string `json:"userId"`
}

type UserController struct {
	userService     *service.UserService
	passportService *service.PassportService
	courseService   *service.CourseService
	hobbyService    *service.HobbyService
}

func NewUserController(userService *service.UserService, passportService *service.PassportService, courseService *service.CourseService, hobbyService *service.HobbyService) *UserController {
	return &UserController{
		userService:     userService,
		passportService: passportService,
		courseService:   courseService,
		hobbyService:    hobbyService,
	}
}

func (uc *UserController) RegisterUserRoutes(router *mux.Router) {
	uc.userService.Logger.Info().Msg("Registered User Routes")
	router.HandleFunc("/login", uc.LoginTokenHandler).Methods("POST")
	router.HandleFunc("/checkToken", uc.CheckToken).Methods("POST")
	subRoute := router.PathPrefix("/users").Subrouter()
	subRoute.HandleFunc("/", uc.getUsers).Methods("GET")
	subRoute.HandleFunc("/", uc.getUsers).Queries("limit", "{limit:[0-9]+}", "pageNo", "{pageNo:[0-9]+}").Methods("GET")
	subRoute.HandleFunc("/", uc.createUser).Methods("POST")
	subRoute.HandleFunc("/{id}", uc.updateUser).Methods("PUT")
	subRoute.HandleFunc("/{id}", uc.getUser).Methods("GET")
	subRoute.HandleFunc("/{id}", uc.deleteUser).Methods("DELETE")
	subRoute.HandleFunc("/{id}/passports", uc.getPassportByUserId).Methods("GET")
	subRoute.HandleFunc("/{id}/course/{courseId}", uc.deleteUserCourse).Methods("DELETE")
}

type TokenInfo struct {
	Token string `json:"token"`
}

func (uc *UserController) deleteUserCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(uc.userService.GetUsersCount()))
	params := mux.Vars(r)
	id, err := uuid.FromString(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("ID invalid!")
		return
	}
	courseId := params["courseId"]
	var user = model.User{Base: model.Base{ID: id}}
	err2 := uc.userService.DeleteUserCourse(&user, courseId)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error in getting users")
		return
	}
	json.NewEncoder(w).Encode("deleted course")
}

func (uc *UserController) CheckToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tokenInfo TokenInfo
	json.NewDecoder(r.Body).Decode(&tokenInfo)

	token, err := jwt.Parse(tokenInfo.Token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return secretKey, nil
	})

	if err != nil {
		//http.Error(w, "Forbidden", http.StatusForbidden)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid token")
		return
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// If token is valid
		// We found the token in our map
		//log.Printf("Authenticated user %s\n", claims)
		// Pass down the request to the next middleware (or final handler)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("valid token")
	} else {
		//http.Error(w, "Forbidden", http.StatusForbidden)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("invalid tokenn")
		return
	}
}

func validateSecretAnswer(user model.User,secretAnswer string)bool{
	//fmt.Println("*****************************secretAnswerGiven-",secretAnswer)
	var pk rsa.PrivateKey
	json.Unmarshal([]byte(user.PrivateKey),&pk)
	//fmt.Println("*****************************user.SecretAnswer-",user.SecretAnswer,pk)
	ppk:=&pk
	userSecretAnswer, _ := ppk.Decrypt(nil, []byte(user.SecretAnswer), &rsa.OAEPOptions{Hash: crypto.SHA512})
	//fmt.Println("*****************************userSecretAnswer-",string(userSecretAnswer),err)
	if string(userSecretAnswer)==secretAnswer{
		return true
	}
	return false
}

// LoginHandler validates the user credentials
func (uc *UserController) LoginTokenHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
		uc.userService.Logger.Error().Err(err).Msg("Please pass the data as URL form encoded")
		return
	}
	var loginInfo model.LoginData
	json.NewDecoder(r.Body).Decode(&loginInfo)
	email := loginInfo.Email
	user, ok := uc.getUserFromEmail(email)
	password := hashLogic.HashString(loginInfo.Password + user.FirstName + user.LastName + user.Passport.Country)
	secretAnswer:=loginInfo.SecretAnswer
	if !validateSecretAnswer(user,secretAnswer){
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		uc.userService.Logger.Error().Msg("Invalid Credentials")
		return
	}
	fmt.Println("secretAnswer-",secretAnswer)
	uc.userService.Logger.Debug().Str("Email-", email).Str("Password-", password).Msg("User Credentials")
	userId := user.ID.String()
	originalPassword := user.Password
	if ok {
		if password == originalPassword {
			// Create a claims map
			claims := jwt.MapClaims{
				"email":    email,
				"IssuedAt": time.Now().Unix(),
			}
			//claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(secretKey)
			if err != nil {
				w.WriteHeader(http.StatusBadGateway)
				w.Write([]byte(err.Error()))
			}
			response := Response{Token: tokenString, Status: "success", UserId: userId}
			responseJSON, _ := json.Marshal(response)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write(responseJSON)

		} else {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			uc.userService.Logger.Error().Msg("Invalid Credentials")
			return
		}
	} else {
		http.Error(w, "User is not found", http.StatusBadRequest)
		uc.userService.Logger.Error().Msg("User is not found")
		return
	}
}

func (uc *UserController) getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(uc.userService.GetUsersCount()))
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	pageNo, _ := strconv.Atoi(r.FormValue("pageNo"))
	offset := int(math.Abs(float64(limit))) * (pageNo - 1)
	fmt.Println(limit, offset)
	uc.userService.Logger.Info().Int("limit", limit).Int("pageNo", pageNo).Msg("Get All Users")
	var users []model.User
	if !(limit < 0 && pageNo >= 2) {
		err2 := uc.userService.GetAllUsers(&users, limit, offset)
		if err2 != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Error in getting users")
			return
		}
	}
	json.NewEncoder(w).Encode(users)
}

func (uc *UserController) getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(uc.userService.GetUsersCount()))
	params := mux.Vars(r)
	id, err := uuid.FromString(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("ID invalid!")
		return
	}
	var user model.User
	//var emptyUser model.User
	//emptyId:= uuid.Nil
	uc.userService.GetUserById(&user, id)
	fmt.Println(user)
	if user.ID == uuid.Nil {
		json.NewEncoder(w).Encode("Could not find user!")
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(uc.userService.GetUsersCount()))
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	user.Password = hashLogic.HashString(user.Password + user.FirstName + user.LastName + user.Passport.Country)
	privateKey,_:=rsa.GenerateKey(rand.Reader, 2048)
	privk,_:=json.Marshal(privateKey)
	user.PrivateKey=string(privk)
	publickey := privateKey.PublicKey
	cipher, err5 := rsa.EncryptOAEP(sha512.New(), rand.Reader, &publickey, []byte(user.SecretAnswer), nil)
	user.SecretAnswer=string(cipher)
	fmt.Println("cipher-",err5,len(user.SecretAnswer))
	//fmt.Println("private and public key-",user.Publickey,user.PrivateKey)
	fmt.Println(user.Password)
	var fromDbUser model.User
	isvalid := user.Validate()
	fmt.Println("validity-", isvalid)
	if !isvalid {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid email")
		return
	}
	err := uc.userService.GetUserFromEmail(&fromDbUser, user.Email)
	if err == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("user already exists")
		return
	}
	err2 := uc.userService.CreateUser(&user)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("err2")
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("User-Count", strconv.Itoa(uc.userService.GetUsersCount()))
	params := mux.Vars(r)
	id, err := uuid.FromString(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("ID invalid!")
		return
	}
	var updatedUser model.User
	json.NewDecoder(r.Body).Decode(&updatedUser)
	updatedUser.ID = id
	if !uc.userService.CheckIfUserExists(id) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Could not find user!")
		return
	}
	// var emptyPassport model.Passport
	// if updatedUser.Passport == emptyPassport {
	// 	var passport model.Passport
	// 	uc.passportService.GetPassportByUserId(&passport, id)
	// 	uc.passportService.DeletePassport(passport.ID)
	// }
	//map[string]interface{} solution
	//https://stackoverflow.com/questions/64330504/update-method-does-not-update-zero-value
	err2 := uc.userService.UpdateUser(updatedUser)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error in updating user")
		return
	}
	json.NewEncoder(w).Encode("User updated")
}

func (uc *UserController) deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("User-Count", strconv.Itoa(uc.userService.GetUsersCount()))
	params := mux.Vars(r)
	idToDelete, err := uuid.FromString(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("ID invalid!")
		return
	}
	if !uc.userService.CheckIfUserExists(idToDelete) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Could not find user!")
		return
	}
	var user model.User
	uc.userService.GetUserById(&user, idToDelete)
	uc.passportService.DeletePassport(user.Passport.ID)
	for _, hobby := range user.Hobbies {
		uc.hobbyService.DeleteHobby(hobby.ID)
	}
	err2 := uc.userService.DeleteUser(idToDelete)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error in deleting user")
		return
	}

	json.NewEncoder(w).Encode("User deleted")
}

func (uc *UserController) getPassportByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(uc.userService.GetUsersCount()))
	params := mux.Vars(r)
	id, err := uuid.FromString(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("ID invalid!")
		return
	}
	var passport model.Passport
	if !uc.userService.CheckIfUserExists(id) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Could not find user!")
		return
	}
	err2 := uc.passportService.GetPassportByUserId(&passport, id)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error in getting passport")
		return
	}
	json.NewEncoder(w).Encode(passport)
}

func (uc *UserController) getPasswordFromEmail(email string) (string, string, bool) {
	var user model.User
	err := uc.userService.GetUserFromEmail(&user, email)
	if err != nil {
		return "", "", false
	}
	return user.Password, user.ID.String(), true

}

func (uc *UserController) getUserFromEmail(email string) (model.User, bool) {
	var user model.User
	err := uc.userService.GetUserFromEmail(&user, email)
	if err != nil {
		return model.User{}, false
	}
	return user, true

}

/*
func (uc *UserController) doesUserExist(id uuid.UUID) bool {
	var user model.User
	uc.userService.GetUserById(&user, id)
	if user.ID == uuid.Nil {
		return false
	}
	return true
}*/
