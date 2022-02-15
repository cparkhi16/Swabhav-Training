package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

var db *gorm.DB

type Flower struct {
	Name    string
	UserId  int `gorm:"primary_key;not null"`
	Address string
	IsMale  *bool
	Color
}

type Color struct {
	ID        uuid.UUID `gorm:"type:varchar(50);unique_index;"`
	IsRed     bool
	CreatedBy string
}

func main() {
	dataSourceName := "root:Panda@19@tcp(localhost:3306)/goadv?parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	} else {
		fmt.Println("connected to database")
	}
	fmt.Println(db)

	db.Exec("USE goadv")
	db.AutoMigrate(&Flower{})

	//create

	//redColor := Color{ID: uuid.NewV4(), isRed: true}

	booltrue := true
	boolfalse := false

	planet2 := Flower{Name: "user2", UserId: 2, Address: "kan", IsMale: &booltrue, Color: Color{ID: uuid.NewV4(), IsRed: true, CreatedBy: "xyz"}}
	db.Debug().Create(&planet2)
	fmt.Println(planet2)

	planet3 := Flower{Name: "user3", UserId: 3, Address: "jahs", IsMale: &boolfalse, Color: Color{ID: uuid.NewV4(), IsRed: true, CreatedBy: "pmn"}}
	db.Debug().Create(&planet3)
	fmt.Println(planet3)

	planet4 := Flower{Name: "user4", UserId: 4, Address: "defs", IsMale: &booltrue, Color: Color{ID: uuid.NewV4(), IsRed: true, CreatedBy: "abc"}}
	db.Debug().Create(&planet4)
	fmt.Println(planet4)

	//GetAll
	var planets []Flower
	db.Debug().Find(&planets)
	fmt.Println(planets)

	//GetFromId
	var planet1 Flower
	ID := 2
	db.Debug().First(&planet1, ID)
	fmt.Println(planet1)

	//update
	updatedFlower := Flower{Name: "user5", UserId: 2, Color: Color{ID: uuid.NewV4()}, IsMale: &boolfalse}
	db.Debug().Save(&updatedFlower)
	fmt.Println(updatedFlower)

	updatedFlower2 := Flower{Name: "user4up", UserId: 4, Color: Color{ID: uuid.NewV4()}, IsMale: &boolfalse}
	db.Debug().Model(&Flower{}).Update(&updatedFlower2)
	//update2
	db.Debug().Table("flowers").Where("user_id = ?", 3).Update("is_male", true)

	db.Debug().Model(&Flower{}).Update(map[string]interface{}{"user_id": 3, "is_red": false})
	//fmt.Println(updatedUser1)

	//delete
	//IDdelete := 2
	//db.Debug().Where("Userid = ?", IDdelete).Delete(&Flower{})

	//flowerToBeDeleted := Flower{Userid: 3}
	//db.Debug().Delete(&flowerToBeDeleted)

}
