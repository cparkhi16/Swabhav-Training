package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Rock struct {
	Name string
	Id   int
}

func CreateRocks(db *gorm.DB) error {
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

	if err := tx.Debug().Create(&Rock{Name: "rock1", Id: 1}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Debug().Create(&Rock{Name: "rock2", Id: 2}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
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

	//db.Exec("USE goadv")
	db.AutoMigrate(&Rock{})

	CreateRocks(db)

}
