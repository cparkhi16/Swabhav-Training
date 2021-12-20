package controller

import (
	h "app/hash"
	m "app/model"
	s "app/service"
	"encoding/json"
	"fmt"
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
	params := mux.Vars(r)
	id, _ := uuid.FromString(params["id"])
	w.Header().Set("Content-Type", "application/json")
	userPassportDetail := uc.us.GetPassportIDForUser(id)
	json.NewEncoder(w).Encode(userPassportDetail)

}
func (uc *UserController) GetUserToken(w http.ResponseWriter, r *http.Request) {
	var users []m.User
	users = uc.us.GetUsers(&users, []string{})
	email := r.FormValue("email")
	password := r.FormValue("password")
	validUser := false
	for _, val := range users {
		if val.Email == email {
			if h.ComparePasswords(val.Password, password) {
				validUser = true
				break
			}
		}
	}
	if validUser {
		validToken, err := GenerateJWT()
		if err != nil {
			fmt.Println("Failed to generate token")
		}
		fmt.Fprint(w, validToken)
	} else {
		fmt.Fprintf(w, "Not a valid user")
	}

}
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser m.User
	er := json.NewDecoder(r.Body).Decode(&newUser)

	if er != nil {
		uc.us.Logger.Error().Msg("Error in user JSON decoding")
	}
	uc.us.AddUser(&newUser)

}
func (uc *UserController) UpdateUserPassportDetail(w http.ResponseWriter, r *http.Request) {
	var updateUser m.User
	er := json.NewDecoder(r.Body).Decode(&updateUser)
	if er != nil {
		//log.Fatal("Here in update passport detail", er)
		uc.us.Logger.Error().Msgf("Error while decoding user passport details", er)
	}
	params := mux.Vars(r)
	id, _ := uuid.FromString(params["id"])
	updateUser.ID = id
	e := uc.us.UpdateUser(&updateUser)
	if e != nil {
		//log.Fatal("Error updating passport detail", e)
		uc.us.Logger.Error().Msgf("Error while updating user passport details", er)
	}

}
func (uc *UserController) AddPassportForUser(w http.ResponseWriter, r *http.Request) {
	var updateUser m.User
	er := json.NewDecoder(r.Body).Decode(&updateUser)
	if er != nil {
		uc.us.Logger.Error().Msgf("Error in decoding passport JSON %v", er)
	}
	params := mux.Vars(r)
	id, _ := uuid.FromString(params["id"])
	updateUser.ID = id
	//newPassportID := updateUser.Passport.PassportID
	//fmt.Println("Got PASSPORT ID ", newPassportID)
	e := uc.us.UpdateUser(&updateUser)
	if e != nil {
		uc.us.Logger.Error().Msgf("Error updating passport detail %v", e)
	}

}
func (uc *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []m.User
	w.Header().Set("Content-Type", "application/json")
	// query params
	page := r.FormValue("page")
	limit := r.FormValue("limit")
	hobby := r.FormValue("hobby")
	hob := strings.Split(hobby, ",")
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
func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser m.User
	er := json.NewDecoder(r.Body).Decode(&updateUser)
	if er != nil {
		uc.us.Logger.Error().Msgf("Error in decoding user JSON", er)
	}
	params := mux.Vars(r)
	id, _ := uuid.FromString(params["id"])
	zeroUUID, _ := uuid.FromString("00000000-0000-0000-0000-000000000000")
	if id != zeroUUID {
		updateUser.ID = id
		e := uc.us.UpdateUser(&updateUser)
		if e != nil {
			uc.us.Logger.Error().Msgf("Error updating user detail %v", e)
		}
	} else {
		uc.us.Logger.Error().Msg("Please give a User ID in params")
	}
}
func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var deleteUser m.User
	er := json.NewDecoder(r.Body).Decode(&deleteUser)
	if er != nil {
		uc.us.Logger.Error().Msgf("Error in decoding user JSON", er)
	}
	params := mux.Vars(r)
	id, _ := uuid.FromString(params["id"])
	zeroUUID, _ := uuid.FromString("00000000-0000-0000-0000-000000000000")
	if id != zeroUUID {
		deleteUser.ID = id
		uc.us.DeleteUser(&deleteUser)
	} else {
		uc.us.Logger.Error().Msg("Please enter a User ID in params")
	}
}
func (uc *UserController) DeletePassportDetailsForUser(w http.ResponseWriter, r *http.Request) {
	var deleteUserPassportDetail m.User
	er := json.NewDecoder(r.Body).Decode(&deleteUserPassportDetail)
	if er != nil {
		uc.us.Logger.Error().Msgf("Error in decoding user JSON", er)
	}
	params := mux.Vars(r)
	id, _ := uuid.FromString(params["id"])
	deleteUserPassportDetail.ID = id
	p := []string{"Passport"}
	e := uc.us.FindAndDeletePassport(&deleteUserPassportDetail, p)
	if e != nil {
		//log.Fatal("Error deleting passport detail", e)
		uc.us.Logger.Error().Msgf("Error deleting passport detail ", e)
	}

}

func (uc *UserController) GetAllUserHobbies(w http.ResponseWriter, r *http.Request) {
	var user m.User
	params := mux.Vars(r)
	id, _ := uuid.FromString(params["id"])
	user.ID = id
	hobbies := uc.us.GetUserHobbies(&user)
	w.Header().Set("User-Hobbies-Count", strconv.Itoa(len(hobbies)))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hobbies)
}

func (uc *UserController) AddHobbiesForUser(w http.ResponseWriter, r *http.Request) {
	var user m.User
	er := json.NewDecoder(r.Body).Decode(&user)
	if er != nil {
		uc.us.Logger.Error().Msgf("Error in decoding user hobby JSON", er)
	}
	params := mux.Vars(r)
	id, _ := uuid.FromString(params["id"])
	user.ID = id
	e := uc.us.AddUserHobbies(&user)
	if e != nil {
		uc.us.Logger.Error().Msgf("Error updating user hobbies %v", e)
	}

}

func (uc *UserController) DeleteHobbiesForUser(w http.ResponseWriter, r *http.Request) {
	var user m.User
	er := json.NewDecoder(r.Body).Decode(&user)
	if er != nil {
		uc.us.Logger.Error().Msgf("Error in decoding user hobby JSON", er)
	}
	params := mux.Vars(r)
	id, _ := uuid.FromString(params["id"])
	user.ID = id
	e := uc.us.DeleteUserHobbies(&user)
	if e != nil {
		uc.us.Logger.Error().Msgf("Error deleting user hobbies %v", e)
	}
}
