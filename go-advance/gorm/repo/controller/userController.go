package controller

import (
	m "app/model"
	s "app/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

func GetUserPassport(us *s.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Get Request called !")
		params := mux.Vars(r)
		id, _ := uuid.FromString(params["id"])
		w.Header().Set("Content-Type", "application/json")
		userPassportDetail := us.GetPassportIDForUser(id)
		json.NewEncoder(w).Encode(userPassportDetail)
	}
}
func CreateUser(us *s.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newUser m.User
		er := json.NewDecoder(r.Body).Decode(&newUser)
		if er != nil {
			log.Fatal("Here", er)
		}
		us.AddUser(&newUser)
	}
}
func UpdateUserPassportDetail(us *s.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var updateUser m.User
		er := json.NewDecoder(r.Body).Decode(&updateUser)
		if er != nil {
			log.Fatal("Here in update passport detail", er)
		}
		params := mux.Vars(r)
		id, _ := uuid.FromString(params["id"])
		updateUser.ID = id
		newPassportID := updateUser.Passport.PassportID
		fmt.Println("New passport ID", newPassportID)
		e := us.UpdateUser(&updateUser)
		if e != nil {
			log.Fatal("Error updating passport detail", e)
		}
	}
}
func AddPassportForUser(us *s.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var updateUser m.User
		er := json.NewDecoder(r.Body).Decode(&updateUser)
		if er != nil {
			log.Fatal("Here in add passport detail", er)
		}
		params := mux.Vars(r)
		id, _ := uuid.FromString(params["id"])
		updateUser.ID = id
		newPassportID := updateUser.Passport.PassportID
		fmt.Println("Got PASSPORT ID ", newPassportID)
		e := us.UpdateUser(&updateUser)
		if e != nil {
			log.Fatal("Error updating passport detail", e)
		}
	}
}
func GetAllUsers(us *s.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Get Request called  for users!")
		var users []m.User
		//w.Header().Add("Access-Control-Expose-Header", "Content-Type")
		//w.Header().Add("Access-Control-Expose-Header", "User-Count")
		w.Header().Set("Content-Type", "application/json")
		//p := []string{"Hobbies", "Courses"}
		//us.GetUsers(&users, p)
		//fmt.Println("Users slice len ", len(users))
		//json.NewEncoder(w).Encode(users)

		// query params
		page := r.FormValue("page")
		limit := r.FormValue("limit")
		hobby := r.FormValue("hobby")
		if page != "" && limit != "" {
			pageInt, _ := strconv.Atoi(page)
			limitInt, _ := strconv.Atoi(limit)

			//http://localhost:9000/users?limit=2&page=1&hobby=Cycling
			//http://localhost:9000/users?limit=2&page=1
			users = us.GetAllUsersWithPagination(pageInt, limitInt, hobby, &users)

			w.Header().Set("User-Count", strconv.Itoa(len(users)))
			fmt.Println("Users slice len ", len(users))
			json.NewEncoder(w).Encode(users)
		} else {
			p := []string{"Hobbies", "Courses"}
			us.GetUsers(&users, p)
			fmt.Println("Users slice len ", len(users))
			w.Header().Set("User-Count", strconv.Itoa(len(users)))
			json.NewEncoder(w).Encode(users)
		}
	}
}

func DeletePassportDetailsForUser(us *s.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var deleteUserPassportDetail m.User
		params := mux.Vars(r)
		id, _ := uuid.FromString(params["id"])
		deleteUserPassportDetail.ID = id
		p := []string{"Passport"}
		e := us.FindAndDeletePassport(&deleteUserPassportDetail, p)
		if e != nil {
			log.Fatal("Error deleting passport detail", e)
		}
	}
}
