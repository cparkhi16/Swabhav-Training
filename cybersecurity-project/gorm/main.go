package main

import (
	"farmDB/animal"
	"farmDB/controller"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	controller.InitDB()
	dog := animal.Animal{
		ID:   1,
		Name: "joy",
		Age:  34,
	}
	/*
		cat := animal.Animal{
			ID:   2,
			Name: "catt",
			Age:  12,
		}*/
	controller.UpdateAnimal(dog)
	//controller.CreateAnimal(cat)
	controller.GetAnimals()
	controller.GetAnimal(2)

}
