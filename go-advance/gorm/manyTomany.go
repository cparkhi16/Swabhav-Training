package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Entity struct {
	TestModel
	Name    string
	Address string    `gorm:"column:ADDR"`
	Courses []*Course `gorm:"many2many:person_courses;association_autoupdate:false"`
}
type Course struct {
	TestModel
	Name   string
	People []*Entity `gorm:"many2many:person_courses;association_autoupdate:false"`
}
type TestModel struct {
	CreatedBy string
	CreatedAt time.Time
	DeletedAt *time.Time
	ID        int `gorm:"type:varchar(36);primary_key"`
}

/*func (ind Individual) FindAllEntities(db *gorm.DB) ([]Individual, error) {

	var individuals []Individual
	db.Preload("Hobbies").Find(&individuals)
	return individuals, nil

}*/
func main() {
	dbConn := "root:hello@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dbConn)
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}
	db.AutoMigrate(&Entity{})
	db.AutoMigrate(&Course{})
	//Courses
	//d := Course{Name: "Golang", TestModel: TestModel{ID: 10, CreatedBy: "Cjp", CreatedAt: time.Now()}}
	//c := Course{Name: "Java", TestModel: TestModel{ID: 11, CreatedBy: "Cjp", CreatedAt: time.Now()}}
	newCourse := Course{Name: "ML", TestModel: TestModel{ID: 29, CreatedBy: "Cjp", CreatedAt: time.Now()}}
	//Entities
	//ef := Entity{Name: "Chinmay", Address: "NYA", TestModel: TestModel{ID: 101, CreatedBy: "Cjp", CreatedAt: time.Now()},
	//	Courses: []*Course{&d, &c}}
	//db.Create(&ef)
	//e := Entity{Name: "Ram", Address: "Nxy", TestModel: TestModel{ID: 102, CreatedBy: "Cjp", CreatedAt: time.Now()},
	//		Courses: []*Course{&c}}
	//	db.Create(&e)
	//e := Entity{Name: "Raju", Address: "Nxy", TestModel: TestModel{ID: 103, CreatedBy: "Cjp", CreatedAt: time.Now()},
	//	Courses: []*Course{&c}}
	//db.Create(&e)
	eo := Entity{Name: "manish", Address: "NYA", TestModel: TestModel{ID: 199, CreatedBy: "Cjp", CreatedAt: time.Now()},
		Courses: []*Course{&newCourse}}
	db.Debug().Create(&eo)
	var people []Entity
	var co Course
	db.First(&co, "id = ?", 11)
	fmt.Println("Course info ", co)
	f := db.Model(&co).Debug().Related(&people, "People").Error
	if f != nil {
		fmt.Println("Error :", f)
	}
	for i, _ := range people {
		fmt.Println("Entities with java course ", people[i].Name)
	}

	var cs []Course
	var en Entity
	db.First(&en, "id = ?", 101)
	fmt.Println("Entity info ", en)
	fe := db.Model(&en).Debug().Related(&cs, "Courses").Error
	if fe != nil {
		fmt.Println("Error :", f)
	}
	for i, _ := range cs {
		fmt.Println("User ID 101 courses :: ", cs[i].Name)
	}

	var dUser Entity
	dUser.ID = 103
	//db.Debug().Delete(&dUser)

}
