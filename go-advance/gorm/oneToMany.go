package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Individual struct {
	CustomModel
	Name    string
	Address string
	Hobbies []Hobby //`gorm:"ForeignKey:IndividualID"`
}
type Hobby struct {
	CustomModel
	IndividualID uuid.UUID `gorm:"type:varchar(36)"`
	HobbyName    string
}
type CustomModel struct {
	CreatedBy string
	CreatedAt time.Time
	DeletedAt *time.Time
	ID        uuid.UUID `gorm:"type:varchar(36);primary_key"`
}

func (ind Individual) FindAll(db *gorm.DB) ([]Individual, error) {

	var individuals []Individual
	db.Debug().Preload("Hobbies").Find(&individuals)
	return individuals, nil

}
func (i Individual) FindWithoutPreload(db *gorm.DB) {
	var users []Individual
	db.Find(&users)
	for _, val := range users {
		fmt.Println("-----------------------", val.Name)
		var u Individual
		u.ID = val.ID
		db.Debug().Find(&u)
		fmt.Println("Hobbies ", u)
	}

}
func main() {
	dbConn := "root:hello@tcp(127.0.0.1:3306)/newdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dbConn)
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}
	db.Model(&Hobby{}).AddForeignKey("individual_id", "individuals(id)", "RESTRICT", "RESTRICT")
	//db.AutoMigrate(&Individual{})
	//db.AutoMigrate(&Hobby{})
	//var i Individual
	//i = Individual{Name: "C", Address: "Xyz", CustomModel: CustomModel{ID: uuid.NewV4(), CreatedBy: "Chinmay", CreatedAt: time.Now()}}
	//in := Individual{Name: "D", Address: "hyz", CustomModel: CustomModel{ID: uuid.NewV4(), CreatedBy: "Chinmay", CreatedAt: time.Now()}}
	//i = Individual{Name: "E", Address: "yPQR", CustomModel: CustomModel{ID: uuid.NewV4(), CreatedBy: "Chinmay", CreatedAt: time.Now()}}
	//db.Create(&in)
	//3b4268c3-f31b-4a6f-812e-aa9fa9d5d297

	//iId, _ := uuid.FromString("3b4268c3-f31b-4a6f-812e-aa9fa9d5d297")
	//h := Hobby{CustomModel: CustomModel{ID: uuid.NewV4(), CreatedBy: "Chinmay", CreatedAt: time.Now()}, HobbyName: "Cricket", IndividualID: iId}
	//db.Create(&h)
	//newHobby := Hobby{CustomModel: CustomModel{ID: uuid.NewV4(), CreatedBy: "Chinmay", CreatedAt: time.Now()}, HobbyName: "Football", IndividualID: iId}
	//db.Create(&newHobby)

	// ADD HOBBIES FOR a9076f2f-4427-4972-80b5-8583a21282f3
	//iId, _ := uuid.FromString("a9076f2f-4427-4972-80b5-8583a21282f3")
	//h := Hobby{CustomModel: CustomModel{ID: uuid.NewV4(), CreatedBy: "Chinmay", CreatedAt: time.Now()}, HobbyName: "Trekking", IndividualID: iId}
	//db.Create(&h)
	//newHobby := Hobby{CustomModel: CustomModel{ID: uuid.NewV4(), CreatedBy: "Chinmay", CreatedAt: time.Now()}, HobbyName: "Swimming", IndividualID: iId}
	//db.Create(&newHobby)
	//oHobby := Hobby{CustomModel: CustomModel{ID: uuid.NewV4(), CreatedBy: "Chinmay", CreatedAt: time.Now()}, HobbyName: "Video Games", IndividualID: iId}
	//db.Create(&oHobby)

	// ADD HOBBIES FOR 383989d6-0e38-4c4c-89b9-90d85c5a6dd9
	/*	iIdN, _ := uuid.FromString("383989d6-0e38-4c4c-89b9-90d85c5a6dd9")
		hi := Hobby{CustomModel: CustomModel{ID: uuid.NewV4(), CreatedBy: "Chinmay", CreatedAt: time.Now()}, HobbyName: "Trekking", IndividualID: iIdN}
		db.Debug().Create(&hi)*/
	iId, _ := uuid.FromString("21556942-ac09-4c82-b6e3-9606066e0a05")
	h := Hobby{CustomModel: CustomModel{ID: uuid.NewV4(), CreatedBy: "Chinmay", CreatedAt: time.Now()}, HobbyName: "Jogging", IndividualID: iId}
	e := db.Create(&h).Error
	if e != nil {
		fmt.Println("Error assigning a hobby to a user ID which is not present in Individuals", e)
	}
	i := Individual{}
	a, _ := i.FindAll(db)
	for _, val := range a {
		fmt.Println(val.Name)
		//fmt.Println(val.ToString())
		fmt.Println("Hobbies: ", len(val.Hobbies))
		if len(val.Hobbies) > 0 {
			for _, h := range val.Hobbies {
				fmt.Println(h)
				fmt.Println("=============================")
			}
		}
		fmt.Println("--------------------")
	}
	//i.FindWithoutPreload(db)
	//on delete restrict
	//Deleting parent table(Individual) will give error as child table (Hobbies) still has that entry
	// NOT APPLICABLE FOR SOFT DELETE
	/*var u Individual
	u.ID, _ = uuid.FromString("383989d6-0e38-4c4c-89b9-90d85c5a6dd9")
	err = db.Delete(&u).Error // Parent row delete
	var ho Hobby
	ho.IndividualID = u.ID
	db.Where("individual_id = ?", u.ID).First(&ho)
	er := db.Debug().Delete(&ho).Error // Child row delete
	if er != nil {
		fmt.Println("Error deleting child table entry")
	}
	fmt.Println("H", ho)
	if err != nil {
		fmt.Println("Error deleting parent table entry ")
	}*/
}
