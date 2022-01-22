package controller

import (
	h "app/hash"
	"app/model"
	m "app/model"
	s "app/service"
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type UserController struct {
	us *s.UserService
}

func NewUserController(us *s.UserService) *UserController {
	return &UserController{us: us}
}
func isValidMailFormat(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
func (uc *UserController) GetUserPassport(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, erID := uuid.FromString(params["id"])
	if erID == nil {
		w.Header().Set("Content-Type", "application/json")
		userPassportDetail := uc.us.GetPassportIDForUser(id)
		json.NewEncoder(w).Encode(userPassportDetail)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Incorrect UUID ")
	}

}
func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	var user m.User
	params := mux.Vars(r)
	id, erID := uuid.FromString(params["id"])
	if erID == nil {
		w.Header().Set("Content-Type", "application/json")
		userDetail, _ := uc.us.GetUserById(&user, id, []string{"Hobbies", "Courses", "Passport"})
		json.NewEncoder(w).Encode(userDetail)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Incorrect UUID ")
	}

}
func (uc *UserController) ValidateToken(w http.ResponseWriter, r *http.Request) {
	type Token struct {
		Token string
	}
	type Response struct {
		IsValidToken bool
	}
	var t Token
	er := json.NewDecoder(r.Body).Decode(&t)
	if er != nil {
		uc.us.Logger.Error().Msg("Error in Token JSON decoding")
	}
	_, err := jwt.Parse(t.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}
		return mySigningKey, nil
	})
	var res Response
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		res = Response{IsValidToken: false}
		//fmt.Fprintf(w, err.Error())
	} else {
		res = Response{IsValidToken: true}
	}
	json.NewEncoder(w).Encode(res)
}
func (uc *UserController) GetUserToken(w http.ResponseWriter, r *http.Request) {
	var users []m.User
	users = uc.us.GetUsers(&users, []string{})
	var login m.Login
	type Response struct {
		Token string
		ID    string
	}
	er := json.NewDecoder(r.Body).Decode(&login)
	if er != nil {
		uc.us.Logger.Error().Msg("Error in login JSON decoding")
	}
	validUser := false
	var userID uuid.UUID
	for _, val := range users {
		if val.Email == login.Email {
			if h.ComparePasswords(val.Password, login.Password) {
				userID = val.ID
				validUser = true
				break
			}
		}
	}
	if validUser {
		validToken, err := generateJWT()
		r := Response{Token: validToken, ID: userID.String()}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(r)
		if err != nil {
			fmt.Println("Failed to generate token")
		}
		//fmt.Fprint(w, validToken)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Invalid credentials")
	}

}
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser m.User
	//fmt.Println("Here in create user")
	type Response struct {
		Token  string
		UserID string
	}
	er := json.NewDecoder(r.Body).Decode(&newUser)
	if er != nil {
		w.WriteHeader(http.StatusBadRequest)
		uc.us.Logger.Error().Msg("Error in user JSON decoding")
	}
	c := uc.us.GetUsersCount("email = ?", newUser.Email)
	if c == 0 {
		if isValidMailFormat(newUser.Email) {
			uc.us.AddUser(&newUser)
			token, _ := generateJWT()
			w.Header().Set("Content-Type", "application/json")
			res := Response{Token: token, UserID: newUser.ID.String()}
			json.NewEncoder(w).Encode(res)
		} else {
			uc.us.Logger.Error().Msg("Invalid email format")
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "User already exists with this email address!!")
	}

}
func (uc *UserController) UpdateUserPassportDetail(w http.ResponseWriter, r *http.Request) {
	var updateUser m.User
	er := json.NewDecoder(r.Body).Decode(&updateUser)
	if er != nil {
		w.WriteHeader(http.StatusBadRequest)
		uc.us.Logger.Error().Msgf("Error while decoding user passport details", er)
	}
	params := mux.Vars(r)
	id, erID := uuid.FromString(params["id"])
	if erID == nil {
		updateUser.ID = id
		e := uc.us.UpdateUser(&updateUser)
		if e != nil {
			w.WriteHeader(http.StatusBadRequest)
			uc.us.Logger.Error().Msgf("Error while updating user passport details", er)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Incorrect UUID ")
	}

}
func (uc *UserController) DeleteCourseForUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, erID := uuid.FromString(params["id"])
	courseid, cerID := uuid.FromString(params["courseid"])
	if erID != nil || cerID != nil {
		uc.us.Logger.Error().Msg("Error in converting id ")
		return
	}
	var user model.User
	user.ID = id
	err := uc.us.UnEnrollCourse(&user, courseid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		uc.us.Logger.Error().Msgf("Error deleting asscoiated course %v ", err)
		return
	}
}
func (uc *UserController) AddPassportForUser(w http.ResponseWriter, r *http.Request) {
	var updateUser m.User
	er := json.NewDecoder(r.Body).Decode(&updateUser)
	fmt.Println("Passport ID ", updateUser.Passport.PassportID)
	if er != nil {
		uc.us.Logger.Error().Msgf("Error in decoding passport JSON %v", er)
	}
	params := mux.Vars(r)
	id, erID := uuid.FromString(params["id"])
	if erID == nil {
		updateUser.ID = id
		e := uc.us.UpdateUser(&updateUser)
		if e != nil {
			w.WriteHeader(http.StatusBadRequest)
			uc.us.Logger.Error().Msgf("Error updating passport detail %v", e)
			fmt.Fprintf(w, e.Error())
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Incorrect UUID ")
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
		if limitInt < 0 && pageInt >= 2 {
			return
		}
		//http://localhost:9000/users?limit=2&page=1&hobby=Cycling
		//http://localhost:9000/users?limit=2&page=1
		users = uc.us.GetAllUsersWithPagination(pageInt, limitInt, hob, &users)
		w.Header().Set("User-Count", strconv.Itoa(len(users)))
		fmt.Println("Users slice len ", len(users))
		json.NewEncoder(w).Encode(users)
	} else {
		p := []string{"Hobbies", "Courses", "Passport"}
		uc.us.GetUsers(&users, p)
		fmt.Println("Users slice len ", len(users))
		w.Header().Set("User-Count", strconv.Itoa(len(users)))
		json.NewEncoder(w).Encode(users)
	}

}
func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In update user")
	var updateUser m.User
	er := json.NewDecoder(r.Body).Decode(&updateUser)
	if er != nil {
		uc.us.Logger.Error().Msgf("Error in decoding user JSON", er)
	}
	params := mux.Vars(r)
	id, erID := uuid.FromString(params["id"])
	if erID == nil {
		if id != uuid.Nil {
			updateUser.ID = id
			e := uc.us.UpdateUser(&updateUser)
			if e != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, e.Error())
				uc.us.Logger.Error().Msgf("Error updating user detail %v", e)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			uc.us.Logger.Error().Msg("Please give a User ID in params")
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Incorrect UUID ")
	}
}
func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete user ")
	var deleteUser m.User
	params := mux.Vars(r)
	id, erID := uuid.FromString(params["id"])
	hardDelete := r.FormValue("hardDelete")
	if erID == nil {
		if id != uuid.Nil {
			deleteUser.ID = id
			deleteErr := uc.us.DeleteUser(&deleteUser, hardDelete)
			if deleteErr != nil {
				w.WriteHeader(http.StatusBadRequest)
				uc.us.Logger.Error().Msg("Error deleting user ")
				fmt.Fprintf(w, deleteErr.Error())
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			uc.us.Logger.Error().Msg("Please enter a User ID in params")
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Incorrect UUID ")
	}
}
func (uc *UserController) DeletePassportDetailsForUser(w http.ResponseWriter, r *http.Request) {
	var deleteUserPassportDetail m.User
	params := mux.Vars(r)
	id, erID := uuid.FromString(params["id"])
	if erID == nil {
		deleteUserPassportDetail.ID = id
		e := uc.us.FindAndDeletePassport(&deleteUserPassportDetail)
		if e != nil {
			w.WriteHeader(http.StatusBadRequest)
			uc.us.Logger.Error().Msgf("Error deleting passport detail ", e)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Incorrect UUID ")
	}

}

func (uc *UserController) GetAllUserHobbies(w http.ResponseWriter, r *http.Request) {
	var user m.User
	params := mux.Vars(r)
	id, erID := uuid.FromString(params["id"])
	if erID == nil {
		user.ID = id
		hobbies := uc.us.GetUserHobbies(&user)
		w.Header().Set("User-Hobbies-Count", strconv.Itoa(len(hobbies)))
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(hobbies)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Incorrect UUID ")
	}
}

func (uc *UserController) AddHobbiesForUser(w http.ResponseWriter, r *http.Request) {
	var user m.User
	er := json.NewDecoder(r.Body).Decode(&user)
	if er != nil {
		uc.us.Logger.Error().Msgf("Error in decoding user hobby JSON", er)
	}
	params := mux.Vars(r)
	id, erID := uuid.FromString(params["id"])
	if erID == nil {
		user.ID = id
		e := uc.us.AddUserHobbies(&user)
		if e != nil {
			w.WriteHeader(http.StatusBadRequest)
			uc.us.Logger.Error().Msgf("Error updating user hobbies %v", e)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Incorrect UUID ")
	}

}

func (uc *UserController) DeleteHobbiesForUser(w http.ResponseWriter, r *http.Request) {
	var user m.User
	params := mux.Vars(r)
	id, erID := uuid.FromString(params["id"])
	if erID == nil {
		user.ID = id
		e := uc.us.DeleteUserHobbies(&user)
		if e != nil {
			w.WriteHeader(http.StatusBadRequest)
			uc.us.Logger.Error().Msgf("Error deleting user hobbies %v", e)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Incorrect UUID ")
	}
}
