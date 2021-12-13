package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Customer struct {
	Model
	Name       string
	Address    string
	Activities []Activity `gorm:"ForeignKey:CustomerRefer"`
}
type Activity struct {
	Model
	CustomerRefer uuid.UUID `gorm:"type:varchar(36);primary_key"`
	HobbyName     string
}
type Model struct {
	CreatedBy string
	CreatedAt time.Time
	DeletedAt *time.Time
	ID        uuid.UUID `gorm:"type:varchar(36);primary_key"`
}

func (c Customer) FindAllCustomers(db *gorm.DB) ([]Customer, error) {

	var customers []Customer
	db.Preload("Activities").Find(&customers)
	return customers, nil

}
func main() {
	dbConn := "root:hello@tcp(127.0.0.1:3306)/newdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dbConn)
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}
	db.AutoMigrate(&Customer{})
	db.AutoMigrate(&Activity{})
	f := db.Model(&Activity{}).Debug().AddForeignKey("customer_refer", "customers(id)", "RESTRICT", "RESTRICT").Error
	if f != nil {
		fmt.Println(f)
	}
	hobby := make([]Activity, 1)
	hobby[0] = Activity{
		Model: Model{ID: uuid.NewV4(), CreatedBy: "Chinmay", CreatedAt: time.Now()}, HobbyName: "Cricket"}
	cust := Customer{Name: "New", Address: "YHS", Model: Model{ID: uuid.NewV4(), CreatedBy: "Chin", CreatedAt: time.Now()}, Activities: []Activity{hobby[0]}}
	db.Create(&cust)
	/*var i Customer
	i = Customer{Name: "CH", Address: "Xyz", Model: Model{ID: uuid.NewV4(), CreatedBy: "Chinmay", CreatedAt: time.Now()}}
	db.Create(&i)
	in := Customer{Name: "DH", Address: "hyz", Model: Model{ID: uuid.NewV4(), CreatedBy: "Chinmay", CreatedAt: time.Now()}}
	j := Customer{Name: "EY", Address: "yPQR", Model: Model{ID: uuid.NewV4(), CreatedBy: "Chinmay", CreatedAt: time.Now()}}
	db.Create(&j)
	db.Create(&in)*/

	/*	iId, _ := uuid.FromString("33f4120a-22dc-4e78-9166-1b31fcea0fd3")
		h := Activity{Model: Model{ID: uuid.NewV4(), CreatedBy: "Chinmay", CreatedAt: time.Now()}, HobbyName: "Cricket", CustomerRefer: iId}
		db.Create(&h)
		newAct := Activity{Model: Model{ID: uuid.NewV4(), CreatedBy: "Chinmay", CreatedAt: time.Now()}, HobbyName: "Football", CustomerRefer: iId}
		db.Create(&newAct)*/

	// ADD HOBBIES FOR a9076f2f-4427-4972-80b5-8583a21282f3
	/*	iId, _ := uuid.FromString("b6f9341a-9bd9-4f2a-b27b-92a9c9538208")
		h := Activity{Model: Model{ID: uuid.NewV4(), CreatedBy: "Chinmay", CreatedAt: time.Now()}, HobbyName: "Trekking", CustomerRefer: iId}
		db.Create(&h)
		newHobby := Activity{Model: Model{ID: uuid.NewV4(), CreatedBy: "Chinmay", CreatedAt: time.Now()}, HobbyName: "Swimming", CustomerRefer: iId}
		db.Create(&newHobby)
		oHobby := Activity{Model: Model{ID: uuid.NewV4(), CreatedBy: "Chinmay", CreatedAt: time.Now()}, HobbyName: "Video Games", CustomerRefer: iId}
		db.Create(&oHobby)*/

	// ADD HOBBIES FOR 383989d6-0e38-4c4c-89b9-90d85c5a6dd9
	/*	iIdN, _ := uuid.FromString("383989d6-0e38-4c4c-89b9-90d85c5a6dd9")
		hi := Hobby{CustomModel: CustomModel{ID: uuid.NewV4(), CreatedBy: "Chinmay", CreatedAt: time.Now()}, HobbyName: "Trekking", IndividualID: iIdN}
		db.Debug().Create(&hi)*/
	/*iId, _ := uuid.FromString("ddc3ca9b-0ceb-482e-800a-03d97b4b865f")
	h := Activity{Model: Model{ID: uuid.NewV4(), CreatedBy: "Chinmay", CreatedAt: time.Now()}, HobbyName: "Watching TV", CustomerRefer: iId}
	e := db.Debug().Create(&h).Error
	if e != nil {
		fmt.Println("Error assigning a activity to a user ID which is not present in Customers")
	}*/
	c := Customer{}
	a, _ := c.FindAllCustomers(db)
	for _, val := range a {
		fmt.Println(val.Name)
		//fmt.Println(val.ToString())
		fmt.Println("Hobbies: ", len(val.Activities))
		if len(val.Activities) > 0 {
			for _, h := range val.Activities {
				fmt.Println(h)
				fmt.Println("=============================")
			}
		}
		fmt.Println("--------------------")
	}

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
