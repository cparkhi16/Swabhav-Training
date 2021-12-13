package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

type Animal struct {
	//gorm.Model
	ID   uuid.UUID `gorm:"type:varchar(36);primary_key"`
	Name string
}

func CreateAnimals(db *gorm.DB) error {
	// Note the use of tx as the database handle once you are within a transaction
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&Animal{ID: uuid.NewV4(), Name: "Giraffe"}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&Animal{ID: uuid.NewV4(), Name: "Lion"}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
func DeleteAnimal(db *gorm.DB) error {
	var a Animal
	a.ID, _ = uuid.FromString("23a1eda2-19ec-489f-9d31-0d91ad76f69f")
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	//Deleting non existing animal with above ID will not throw a error(ERR is nil)
	err := tx.Debug().Delete(&a).Error
	fmt.Println("Value of err ", err)
	if err != nil {
		fmt.Println("Rollback ")
		tx.Rollback()
		return err
	}
	//Deleting a existing value
	/*var ab Animal
	ab.ID, _ = uuid.FromString("82429be6-9297-460d-9b9a-0badddd1f620")
	if err := tx.Debug().Delete(&ab).Error; err != nil {
		fmt.Println("Rollback in 2 nd case will not happen")
		tx.Rollback()
		return err
	}*/

	return tx.Commit().Error
}
func FindAndUpdateAnimal(db *gorm.DB) error {
	var a Animal
	a.ID, _ = uuid.FromString("82429be6-9297-460d-9b9a-0badddd1f620") //Non existing ID
	//a.ID, _ = uuid.FromString("91f525b3-a92a-4740-a2b3-a0379d4266ed")
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	err := db.Debug().Find(&a).Error
	if err != nil {
		fmt.Println("Rollback ")
		tx.Debug().Rollback()
		return err
	} else {
		aMap := make(map[string]interface{})
		aMap["name"] = "Cat"
		tx.Model(&a).Update(aMap)
		fmt.Println("Updating db")
	}
	return tx.Commit().Error
}
func main() {
	dbConn := "root:hello@tcp(127.0.0.1:3306)/newdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dbConn)
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}
	//db.AutoMigrate(&Animal{})
	/*e := CreateAnimals(db)
	if e != nil {
		fmt.Println("Error ", e.Error())
	}*/
	/*e := DeleteAnimal(db)
	if e != nil {
		fmt.Println(e.Error())
	}*/
	e := FindAndUpdateAnimal(db)
	if e != nil {
		fmt.Println(e.Error())
	}
}
