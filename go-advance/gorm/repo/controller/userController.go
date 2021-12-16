package controller

import (
	m "app/model"
	s "app/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type UserController struct {
	us *s.UserService
}

func NewUserController(us *s.UserService) *UserController {
	return &UserController{us: us}
}
func (uc *UserController) GetUserPassport(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Get Request called !")
	params := mux.Vars(r)
	id, _ := uuid.FromString(params["id"])
	w.Header().Set("Content-Type", "application/json")
	userPassportDetail := uc.us.GetPassportIDForUser(id)
	json.NewEncoder(w).Encode(userPassportDetail)

}
func (uc *UserController) GetUserToken(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Println("Failed to generate token")
	}
	fmt.Fprint(w, validToken)

}
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

	var newUser m.User
	er := json.NewDecoder(r.Body).Decode(&newUser)
	if er != nil {
		log.Fatal("Here", er)
	}
	uc.us.AddUser(&newUser)

}
func (uc *UserController) UpdateUserPassportDetail(w http.ResponseWriter, r *http.Request) {
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
	e := uc.us.UpdateUser(&updateUser)
	if e != nil {
		log.Fatal("Error updating passport detail", e)
	}

}
func (uc *UserController) AddPassportForUser(w http.ResponseWriter, r *http.Request) {
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
	e := uc.us.UpdateUser(&updateUser)
	if e != nil {
		log.Fatal("Error updating passport detail", e)
	}

}
func (uc *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
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
	//var hobbies []string
	hob := strings.Split(hobby, ",")
	/*for i := 0; i < len(hob); i++ {
		hobbies = append(hobbies, strconv.Quote(hob[i]))
	}*/
	fmt.Println("Hobbies comma separated list ", hob)
	if page != "" && limit != "" {
		pageInt, _ := strconv.Atoi(page)
		limitInt, _ := strconv.Atoi(limit)

		//http://localhost:9000/users?limit=2&page=1&hobby=Cycling
		//http://localhost:9000/users?limit=2&page=1
		users = uc.us.GetAllUsersWithPagination(pageInt, limitInt, hob, &users)

		w.Header().Set("User-Count", strconv.Itoa(len(users)))
		fmt.Println("Users slice len ", len(users))
		json.NewEncoder(w).Encode(users)
	} else {
		p := []string{"Hobbies", "Courses"}
		uc.us.GetUsers(&users, p)
		fmt.Println("Users slice len ", len(users))
		w.Header().Set("User-Count", strconv.Itoa(len(users)))
		json.NewEncoder(w).Encode(users)
	}

}

func (uc *UserController) DeletePassportDetailsForUser(w http.ResponseWriter, r *http.Request) {
	var deleteUserPassportDetail m.User
	params := mux.Vars(r)
	id, _ := uuid.FromString(params["id"])
	deleteUserPassportDetail.ID = id
	p := []string{"Passport"}
	e := uc.us.FindAndDeletePassport(&deleteUserPassportDetail, p)
	if e != nil {
		log.Fatal("Error deleting passport detail", e)
	}

}

func GetTestOne(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test One called ")
}

func GetTestTwo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test Two called ")
}
