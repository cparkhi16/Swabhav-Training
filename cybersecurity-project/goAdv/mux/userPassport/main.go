package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"userPassport/controller"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//1058963771578-gqagesvt69mmqmv8bk9ecni04pko6frc.apps.googleusercontent.com
//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

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	logger.Info().Msg("server starting...")

	db, _ := ConnectToDB("root", "Panda@19", "localhost", "3306", "userpassport")
	controller.ControllerInit(db, logger)

	fmt.Println(reflect.TypeOf(logger))
	router := mux.NewRouter()
	controller.RegisterRoutes(router)
	router.HandleFunc("/login", controller.GetTokenHandler)
	log.Fatal().Err(http.ListenAndServe(":8000", router))

}

func ConnectToDB(username string, password string, host string, port string, dbName string) (*gorm.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local", username, password, host, port, dbName)
	var err error
	var db *gorm.DB
	db, err = gorm.Open("mysql", dataSourceName)
	/*
		db.AutoMigrate(&model.User{}, &model.Passport{})
		err2 := db.Debug().Model(&model.Passport{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").Error
		if err2 != nil {
			log.Error().Err(err).Msg("")
		}*/
	return db, err
}
