package controller

import (
	"encoding/json"
	"net/http"
	"server/model"
	"strconv"

	"github.com/gorilla/mux"
)

var animals = []model.Animal{
	{Name: "dog", ID: 23, Age: 34},
	{Name: "cow", ID: 29, Age: 52},
}

func GetAnimals(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(animals)
}

func GetAnimal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var animalFound bool

	for _, animal := range animals {
		if animal.ID == id {
			animalFound = true
			json.NewEncoder(w).Encode(animal)
			break
		}
	}
	if !animalFound {
		json.NewEncoder(w).Encode("Unable to find animal with matching ID")
	}
}

func CreateAnimal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var animal model.Animal
	_ = json.NewDecoder(r.Body).Decode(&animal)
	animals = append(animals, animal)
	json.NewEncoder(w).Encode(animal)
}

func UpdateAnimal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var animalFound bool
	var updatedAnimal model.Animal
	_ = json.NewDecoder(r.Body).Decode(&updatedAnimal)
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for i := 0; i < len(animals); i++ {
		if animals[i].ID == id {
			animalFound = true
			animals[i].Name = updatedAnimal.Name
			animals[i].Age = updatedAnimal.Age
			json.NewEncoder(w).Encode(animals[i])
			break
		}
	}
	if !animalFound {
		json.NewEncoder(w).Encode("Unable to find animal with matching ID")
	}

}

func DeleteAnimal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	for index, animal := range animals {
		if animal.ID == id {
			animals = append(animals[:index], animals[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(animals)
}
