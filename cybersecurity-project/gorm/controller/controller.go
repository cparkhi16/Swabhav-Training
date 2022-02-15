package controller

import (
	"farmDB/animal"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func InitDB() {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	} else {
		fmt.Println("connected to database")
	}

	//db.Exec("CREATE DATABASE animals")
	db.Exec("USE animals")
	db.AutoMigrate(&animal.Animal{})
}

func GetAnimals() {
	var animals []animal.Animal
	db.Find(&animals)
	fmt.Println(animals)
}

func GetAnimal(ID int) {
	var animal animal.Animal
	db.First(&animal, ID)
	fmt.Println(animal)
}

func CreateAnimal(animal animal.Animal) {
	db.Create(&animal)
	fmt.Println(animal)
}

func UpdateAnimal(updatedAnimal animal.Animal) {
	db.Save(&updatedAnimal)
	fmt.Println(updatedAnimal)
}

func DeleteAnimal(idToDelete int) {
	db.Where("id = ?", idToDelete).Delete(&animal.Animal{})
}
