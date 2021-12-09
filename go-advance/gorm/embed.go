package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

type Human struct {
	//gorm.Model
	MyModel
	Name    string
	Address string
	Gender  string
	IsMale  bool
}
type MyModel struct {
	CreatedBy string
	CreatedAt time.Time
	DeletedAt *time.Time
	ID        uuid.UUID `gorm:"type:varchar(36);primary_key"`
}

func GetHuman(db *gorm.DB) {
	var p []Human
	db.Debug().Find(&p)
	fmt.Println("DB ALL PEOPLE ", p)
}
func CreateHuman(db *gorm.DB, h *Human) {
	db.Debug().Create(&h)
}
func DeleteHuman(db *gorm.DB) {
	//t := time.Now()
	//db.Model(&h).Debug().Update("deleted_at", t)
	//db.Debug().Delete(&h)
	db.Where("Name = ?", "Keyur").Debug().Delete(&Human{})
}
func main() {
	dbConn := "root:hello@tcp(127.0.0.1:3306)/newdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dbConn)
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}
	//db.AutoMigrate(&Human{})

	/*hid := uuid.NewV4()
	h := Human{Name: "Chinmay", Address: "XYZ", Gender: "Male", IsMale: true, MyModel: MyModel{ID: hid, CreatedBy: "Chinmay", CreatedAt: time.Now()}}
	CreateHuman(db, &h)
	jid := uuid.NewV4()
	j := Human{Name: "Keyur", Address: "PQR", Gender: "Male", IsMale: true, MyModel: MyModel{ID: jid, CreatedBy: "Chinmay", CreatedAt: time.Now()}}
	CreateHuman(db, &j)
	GetHuman(db)
	//var g Human
	//db.Model(&g).Where("Address = ?", "PQR").Debug().Update("Address", "YUV")
	//db.Model(&Human{}).Debug().Update("CreatedBy", "Manager")
	//var g Human
	//g.Name = "Keyur"
	DeleteHuman(db)*/

	/*var updatedHuman Human
	sID, _ := uuid.FromString("23a1eda2-19ec-489f-9d31-0d91ad76f69f")
	updatedHuman.ID = sID
	var b bool = false
	updatedHuman.IsMale = &b //if ptr is not used here then its value is not updated in DB
	db.Model(&updatedHuman).Debug().Update(&updatedHuman)*/
	//Query -  UPDATE `humen` SET `is_male` = false  WHERE `humen`.`deleted_at` IS NULL AND `humen`.
	//`id` = '23a1eda2-19ec-489f-9d31-0d91ad76f69f'

	/*var user Human
	user.ID, _ = uuid.FromString("23a1eda2-19ec-489f-9d31-0d91ad76f69f")
	user.IsMale = true
	db.Debug().Save(&user)*/
	//Query
	// UPDATE `humen` SET `created_by` = '', `deleted_at` = NULL, `name` = '', `address` = '
	//', `gender` = '', `is_male` = true  WHERE `humen`.`deleted_at` IS NULL AND `humen`.`id` = '23a1eda2-19ec-489f-9d31-0d91ad76f69f'
	GetHuman(db)

	// USE MAPS TO UPDATE FIELDS

	var updatedHuman Human
	sID, _ := uuid.FromString("23a1eda2-19ec-489f-9d31-0d91ad76f69f")
	userMap := make(map[string]interface{})
	userMap["id"] = sID
	userMap["is_male"] = true
	db.Model(&updatedHuman).Debug().Update(userMap)

}
