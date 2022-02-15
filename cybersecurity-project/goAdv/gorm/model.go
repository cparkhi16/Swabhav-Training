package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Planet struct {
	Name    string
	Userid  int
	Address string
	IsMale  bool
	gorm.Model
}

func main() {
	dataSourceName := "root:******@tcp(localhost:3306)/goadv?parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	} else {
		fmt.Println("connected to database")
	}
	fmt.Println(db)

	db.Exec("USE goadv")
	db.AutoMigrate(&Planet{})

	//create

	planet2 := Planet{Name: "user2", Userid: 2, Address: "kan", IsMale: true}
	db.Debug().Create(&planet2)
	fmt.Println(planet2)

	planet3 := Planet{Name: "user3", Userid: 3, Address: "jahs", IsMale: false}
	db.Debug().Create(&planet3)
	fmt.Println(planet3)

	planet4 := Planet{Name: "user4", Userid: 4, Address: "defs", IsMale: true}
	db.Debug().Create(&planet4)
	fmt.Println(planet4)

	//GetAll
	var planets []Planet
	db.Debug().Find(&planets)
	fmt.Println(planets)

	//GetFromId
	var planet1 Planet
	ID := 2
	db.Debug().First(&planet1, ID)
	fmt.Println(planet1)

	//update
	updatedPlanet := Planet{Name: "user5", Userid: 10}
	db.Debug().Save(&updatedPlanet)
	fmt.Println(updatedPlanet)

	//update2
	db.Debug().Table("planets").Where("Userid = ?", 3).Update("Name", "hello")
	//fmt.Println(updatedUser1)

	//delete
	IDdelete := 2
	db.Debug().Where("Userid = ?", IDdelete).Delete(&Planet{})

}
