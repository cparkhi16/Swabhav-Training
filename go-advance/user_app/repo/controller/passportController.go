package controller

import (
	m "app/model"
	s "app/service"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type PassportController struct {
	us *s.UserService
}

func NewPassportController(us *s.UserService) *PassportController {
	return &PassportController{us: us}
}
func (pc *PassportController) UpdatePassportDetails(w http.ResponseWriter, r *http.Request) {
	var updatePassport m.Passport
	er := json.NewDecoder(r.Body).Decode(&updatePassport)
	if er != nil {
		w.WriteHeader(http.StatusBadRequest)
		pc.us.Logger.Error().Msgf("Error in decoding passport JSON", er)
	}
	params := mux.Vars(r)
	id, erID := uuid.FromString(params["id"])
	if erID == nil {
		if id != uuid.Nil {
			updatePassport.ID = id
			e := pc.us.UpdatePassport(&updatePassport)
			if e != nil {
				w.WriteHeader(http.StatusBadRequest)
				pc.us.Logger.Error().Msgf("Error updating course detail %v", e)
				fmt.Fprintf(w, e.Error())
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			pc.us.Logger.Error().Msg("Please enter a passport ID in params")
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Incorrect UUID ")
	}
}
