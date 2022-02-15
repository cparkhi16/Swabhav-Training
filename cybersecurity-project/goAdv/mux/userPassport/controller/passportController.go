package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// var repo2 repository.Repository
// var userService *service.UserService
// var passportService *service.PassportService
// var db *gorm.DB

func PassportControllerInit(dB *gorm.DB) {
	// db = dB
	// repo2 = repository.NewRepository()
	// fmt.Println(repo2)
	// userService = service.NewUserService(repo2, db)
	// fmt.Println(userService)
	// passportService = service.NewPassportService(repo2, db)
	// fmt.Println(passportService)
}

func RegisterRoutess(router *mux.Router) {
	subRoute := router.PathPrefix("/passports").Subrouter()
	subRoute.HandleFunc("/", GetPassports).Methods("GET")
	subRoute.HandleFunc("/", GetPassports).Queries("limit", "{limit:[0-9]+}", "pageNo", "{pageNo:[0-9]+}").Methods("GET")
	//subRoute.HandleFunc("/{id}", UpdatePassport).Methods("PUT")
	subRoute.HandleFunc("/{id}", GetPassport).Methods("GET")
	subRoute.HandleFunc("/{id}", DeletePassport).Methods("DELETE")
}

func GetPassports(w http.ResponseWriter, r *http.Request) {
	/*
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount()))
		limit, _ := strconv.Atoi(r.FormValue("limit"))
		pageNo, _ := strconv.Atoi(r.FormValue("pageNo"))
		offset := limit * (pageNo - 1)
		fmt.Println(limit, pageNo)
		var passports []model.Passport
		userService.GetAllUsers(&passports, limit, offset)
		json.NewEncoder(w).Encode(passports)*/
}

func GetPassport(w http.ResponseWriter, r *http.Request) {
	/*
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount()))
		params := mux.Vars(r)
		id, _ := uuid.FromString(params["id"])
		var str []string
		var passport model.Passport
		userService.GetUserById(&passport, id, str)
		json.NewEncoder(w).Encode(passport)*/
}

func DeletePassport(w http.ResponseWriter, r *http.Request) {
	/*
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount()))
		params := mux.Vars(r)
		idToDelete, _ := uuid.FromString(params["id"])
		passportService.DeletePassport(idToDelete)
		w.WriteHeader(http.StatusNoContent)*/
}
