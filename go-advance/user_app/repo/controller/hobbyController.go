package controller

import (
	m "app/model"
	s "app/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type HobbyController struct {
	us *s.UserService
}

func NewHobbyController(us *s.UserService) *HobbyController {
	return &HobbyController{us: us}
}
func (hc *HobbyController) GetHobby(w http.ResponseWriter, r *http.Request) {
	var hobby m.Hobby
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := uuid.FromString(params["id"])
	hobby.ID = id
	hobby = hc.us.GetHobbyById(&hobby, id)
	json.NewEncoder(w).Encode(hobby)
}
func (hc *HobbyController) DeleteHobby(w http.ResponseWriter, r *http.Request) {
	var hobby m.Hobby
	params := mux.Vars(r)
	id, _ := uuid.FromString(params["id"])
	zeroUUID, _ := uuid.FromString("00000000-0000-0000-0000-000000000000")
	if id != zeroUUID {
		hobby.ID = id
		err := hc.us.DeleteHobbyById(&hobby)
		if err != nil {
			hc.us.Logger.Error().Msgf("Error while deleting hobby by ID %v", err)
		}
	} else {
		hc.us.Logger.Error().Msg("Please give a hobby ID in params")
	}
}
func (hc *HobbyController) UpdateHobby(w http.ResponseWriter, r *http.Request) {
	//http://localhost:9000/hobby/1b7f13bc-5699-433f-8b2f-118ed488292c
	var hobby m.Hobby
	er := json.NewDecoder(r.Body).Decode(&hobby)
	if er != nil {
		hc.us.Logger.Error().Msgf("Error in decoding hobby JSON", er)
	}
	params := mux.Vars(r)
	id, _ := uuid.FromString(params["id"])
	hobby.ID = id
	zeroUUID, _ := uuid.FromString("00000000-0000-0000-0000-000000000000")
	if id != zeroUUID {
		hobby.ID = id
		e := hc.us.UpdateHobbyById(&hobby)
		if e != nil {
			hc.us.Logger.Error().Msgf("Error updating hobby detail %v", e)
		}
	} else {
		hc.us.Logger.Error().Msg("Please give a hobby ID in params")
	}
}
func (hc *HobbyController) GetAllHobbies(w http.ResponseWriter, r *http.Request) {
	var hobbies []m.Hobby
	w.Header().Set("Content-Type", "application/json")
	// query params
	page := r.FormValue("page")
	limit := r.FormValue("limit")
	var pageInt int
	var limitInt int
	if page == "" || limit == "" {
		pageInt = 1
		limitInt = 2
	} else {
		pageInt, _ = strconv.Atoi(page)
		limitInt, _ = strconv.Atoi(limit)
	}
	//http://localhost:9000/courses?limit=2&page=1
	hobbies = hc.us.GetAllHobbiesWithPagination(pageInt, limitInt, &hobbies)
	w.Header().Set("Hobby-Count", strconv.Itoa(len(hobbies)))
	fmt.Println("Hobbies slice len ", len(hobbies))
	json.NewEncoder(w).Encode(hobbies)
}
