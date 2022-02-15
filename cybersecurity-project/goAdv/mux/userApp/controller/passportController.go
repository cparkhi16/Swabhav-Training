package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"userPassport/model"
	"userPassport/service"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type PassportController struct {
	passportService *service.PassportService
}

func NewPassportController(passportService *service.PassportService) *PassportController {
	return &PassportController{
		passportService: passportService,
	}
}

func (pc *PassportController) RegisterPassportRoutes(router *mux.Router) {
	pc.passportService.Logger.Info().Msg("Registered Passport Routes")
	subRoute := router.PathPrefix("/passports").Subrouter()
	subRoute.HandleFunc("/{userId}", pc.getPassports).Methods("GET")
	subRoute.HandleFunc("/{userId}", pc.getPassports).Queries("limit", "{limit:[0-9]+}", "pageNo", "{pageNo:[0-9]+}").Methods("GET")
	subRoute.HandleFunc("/{id}", pc.updatePassport).Methods("PUT")
	//subRoute.HandleFunc("/{id}", pc.getPassport).Methods("GET")
	subRoute.HandleFunc("/{id}", pc.deletePassport).Methods("DELETE")
}

func (pc *PassportController) getPassports(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Passport-Count", strconv.Itoa(pc.passportService.GetPassportCount()))
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	pageNo, _ := strconv.Atoi(r.FormValue("pageNo"))
	offset := limit * (pageNo - 1)
	pc.passportService.Logger.Info().Int("limit", limit).Int("pageNo", pageNo).Msg("Get All Passports")
	params := mux.Vars(r)
	userId, err := uuid.FromString(params["userId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("ID invalid!")
		return
	}
	var passports []model.Passport
	pc.passportService.GetAllPassports(&passports, limit, offset, userId)
	json.NewEncoder(w).Encode(passports)
}

func (pc *PassportController) getPassport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Passport-Count", strconv.Itoa(pc.passportService.GetPassportCount()))
	params := mux.Vars(r)
	id, err := uuid.FromString(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("ID invalid!")
		return
	}
	var passport model.Passport
	err2 := pc.passportService.GetPassportById(&passport, id)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Could not find passport")
		return
	}
	json.NewEncoder(w).Encode(passport)
}

func (pc *PassportController) updatePassport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Passport-Count", strconv.Itoa(pc.passportService.GetPassportCount()))
	params := mux.Vars(r)
	id, err := uuid.FromString(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("ID invalid!")
		return
	}
	var updatedPassport model.Passport
	json.NewDecoder(r.Body).Decode(&updatedPassport)
	updatedPassport.ID = id
	if !pc.passportService.CheckIfPassportExists(id) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Could not find passport!")
		return
	}
	err2 := pc.passportService.UpdatePassport(updatedPassport)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error in finding passport")
		return
	}
	json.NewEncoder(w).Encode("Passport updated")
}

func (pc *PassportController) deletePassport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Passport-Count", strconv.Itoa(pc.passportService.GetPassportCount()))
	params := mux.Vars(r)
	idToDelete, err := uuid.FromString(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("ID invalid!")
		return
	}
	if !pc.passportService.CheckIfPassportExists(idToDelete) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Could not find passport!")
		return
	}
	err2 := pc.passportService.DeletePassport(idToDelete)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error in deleting passport")
		return
	}
	json.NewEncoder(w).Encode("Passport deleted")
}

/*
func (pc *PassportController) doesPassportExist(id uuid.UUID) bool {
	var passport model.Passport
	pc.passportService.GetPassportById(&passport, id)
	if passport.ID == uuid.Nil {
		return false
	}
	return true
}*/
