package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Pet struct {
	ID   int `gorm:"primary_key"`
	Name string
	Prof int
}

type Profile struct {
	Name      string
	Pet       Pet `gorm:"foreignkey:Prof"`
	ProfileId int `gorm:"primary_key"`
}

func main() {
	dataSourceName := "root:Panda@19@tcp(127.0.0.1:3306)/goadv?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	} else {
		fmt.Println("connected to database")
	}
	fmt.Println(db)
	/*
		//this is the right way, the struct which has a forign key should be mentioned in Model() in AddForeignKey
			db.AutoMigrate(&Profile{}, &Pet{})

			err2 := db.Debug().Model(&Pet{}).AddForeignKey("prof", "profiles(profile_id)", "CASCADE", "CASCADE").Error
			if err2 != nil {
				fmt.Println(err2)
			}*/

	/*
		err2 := db.Debug().Model(&Profile{}).AddForeignKey("profile_id", "pets(prof)", "CASCADE", "CASCADE").Error
		if err2 != nil {
			fmt.Println(err2)
		}*/

	//db.Debug().Exec("SET FOREIGN_KEY_CHECKS=0;")
	/*
		pet1 := Pet{ID: 11, Name: "B"}
		db.Debug().Create(&pet1)*/

	//pet1 := Pet{ID: 1, Name: "B", Prof: 11}
	//db.Debug().Create(&pet1)
	//db.Debug().Exec("SET FOREIGN_KEY_CHECKS=0;")
	/*
		p := Profile{Name: "C", Pet: Pet{ID: 1, Name: "B", Prof: 11}, ProfileId: 11}
		db.Debug().Create(&p)

		p2 := Profile{Name: "A", Pet: Pet{ID: 2, Name: "B", Prof: 15}, ProfileId: 30}
		db.Debug().Create(&p2)*/

	var profiles []Profile
	err3 := db.Debug().Preload("Pet").Find(&profiles).Error //field name from Profile struct is Pet that should be mentioned in preload
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(profiles)

	/*
		p := Profile{Name: "C", Pet: Pet{ID: 1, Name: "B",Prof:11}, ProfileId: 11}
		db.Debug().Create(&p)

		var profiles []Profile
		db.Debug().Preload("pets").Find(&profiles)
		fmt.Println(profiles)*/
	/*
		db.Set("gorm:auto_preload", true).Find(&profiles)
		fmt.Println(profiles)*/

}
