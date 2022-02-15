package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"userPassport/model"
	"userPassport/service"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type HobbyController struct {
	hobbyService *service.HobbyService
}

func NewHobbyController(hobbyService *service.HobbyService) *HobbyController {
	return &HobbyController{
		hobbyService: hobbyService,
	}
}

func (hc *HobbyController) RegisterHobbyRoutes(router *mux.Router) {
	hc.hobbyService.Logger.Info().Msg("Registered Hobby Routes")
	subRoute := router.PathPrefix("/hobbies").Subrouter()
	subRoute.HandleFunc("/", hc.getHobbies).Methods("GET")
	subRoute.HandleFunc("/", hc.getHobbies).Queries("limit", "{limit:[0-9]+}", "pageNo", "{pageNo:[0-9]+}").Methods("GET")
	subRoute.HandleFunc("/{id}", hc.updateHobby).Methods("PUT")
	subRoute.HandleFunc("/", hc.addHobby).Methods("POST")
	subRoute.HandleFunc("/{id}", hc.deleteHobby).Methods("DELETE")
	//subRoute.HandleFunc("/{id}", hc.getHobby).Methods("GET")
	subRoute.HandleFunc("/{userId}", hc.getHobbiesOfUser).Methods("GET")
}

func (hc *HobbyController) getHobbies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Hobby-Count", strconv.Itoa(hc.hobbyService.GetHobbyCount()))
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	pageNo, _ := strconv.Atoi(r.FormValue("pageNo"))
	offset := limit * (pageNo - 1)
	hc.hobbyService.Logger.Info().Int("limit", limit).Int("pageNo", pageNo).Msg("Get All Hobbies")
	var hobbies []model.Hobby
	hc.hobbyService.GetAllHobbies(&hobbies, limit, offset)
	json.NewEncoder(w).Encode(hobbies)
}

func (uc *HobbyController) getHobby(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Hobby-Count", strconv.Itoa(uc.hobbyService.GetHobbyCount()))
	params := mux.Vars(r)
	id, err := uuid.FromString(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("ID invalid!")
		return
	}
	var hobby model.Hobby
	err2 := uc.hobbyService.GetHobbyById(&hobby, id)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Could on find Hobby")
		return
	}
	json.NewEncoder(w).Encode(hobby)
}

func (uc *HobbyController) getHobbiesOfUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Hobby-Count", strconv.Itoa(uc.hobbyService.GetHobbyCount()))
	params := mux.Vars(r)
	id, err := uuid.FromString(params["userId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("ID invalid!")
		return
	}
	var hobby []model.Hobby
	err2 := uc.hobbyService.GetHobbyByUserId(&hobby, id)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Could on find Hobbies")
		return
	}
	json.NewEncoder(w).Encode(hobby)
}

func (hc *HobbyController) deleteHobby(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Course-Count", strconv.Itoa(cc.courseService.GetCoursesCount()))
	params := mux.Vars(r)
	idToDelete, err := uuid.FromString(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("ID invalid!")
		return
	}
	if !hc.hobbyService.CheckIfHobbyExists(idToDelete) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Could not get hobby!")
		return
	}
	err2 := hc.hobbyService.DeleteHobby(idToDelete)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error in deleting hobby")
		return
	}
	json.NewEncoder(w).Encode("Hobby deleted")
}

func (hc *HobbyController) updateHobby(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Hobby-Count", strconv.Itoa(hc.hobbyService.GetHobbyCount()))
	params := mux.Vars(r)
	id, err := uuid.FromString(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("ID invalid!")
		return
	}
	var updatedHobby model.Hobby
	json.NewDecoder(r.Body).Decode(&updatedHobby)
	updatedHobby.ID = id
	if !hc.hobbyService.CheckIfHobbyExists(id) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Could not find hobby!")
		return
	}
	err2 := hc.hobbyService.UpdateHobby(updatedHobby)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error in updating hobby")
		return
	}
	json.NewEncoder(w).Encode("Updated hobby")
}

func (hc *HobbyController) addHobby(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Hobby-Count", strconv.Itoa(hc.hobbyService.GetHobbyCount()))
	var updatedHobby model.Hobby
	updatedHobby.ID = uuid.NewV4()
	updatedHobby.CreatedBy = "yogesh"
	updatedHobby.CreatedAt = time.Now()
	json.NewDecoder(r.Body).Decode(&updatedHobby)
	fmt.Println(updatedHobby)
	err2 := hc.hobbyService.AddHobby(updatedHobby)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error in adding hobby")
		return
	}
	json.NewEncoder(w).Encode("Updated hobby")
}

/*
func (uc *HobbyController) doesHobbyExist(id uuid.UUID) bool {
	var hobby model.Hobby
	uc.hobbyService.GetHobbyById(&hobby, id)
	if hobby.ID == uuid.Nil {
		return false
	}
	return true
}*/
