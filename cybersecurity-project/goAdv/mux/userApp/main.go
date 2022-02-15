package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"userPassport/controller"
	"userPassport/customlogger"
	"userPassport/model"
	"userPassport/repository"
	"userPassport/service"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	tempFile, err := ioutil.TempFile("./data/logs", "deleteme")
	if err != nil {
		log.Error().Err(err).Msg("error in creating a temporary file")
	}
	logger := zerolog.New(tempFile).With().Logger()
	customlogger.SetLoggerInstance(&logger)
	fmt.Printf("The log file is allocated at %s\n", tempFile.Name())
	logger.Info().Msg("server starting...")
	db, _ := ConnectToDB("root", "root", "localhost", "3306", "userapp")
	fmt.Println(db)
	if err != nil {
		logger.Error().Err(err)
	}
	logger.Info().Msg("connected to DB")
	repo1 := repository.NewRepository()
	userService := service.NewUserService(repo1, db, &logger)
	passportService := service.NewPassportService(repo1, db, &logger)
	hobbyService := service.NewHobbyService(repo1, db, &logger)
	courseService := service.NewCourseService(repo1, db, &logger)
	fileService := service.NewFileService(repo1, db, &logger)

	router := mux.NewRouter()
	router.Use(controller.CheckAuthentication)
	userController := controller.NewUserController(userService, passportService, courseService, hobbyService)
	userController.RegisterUserRoutes(router)
	hobbyController := controller.NewHobbyController(hobbyService)
	hobbyController.RegisterHobbyRoutes(router)
	courseController := controller.NewCourseController(courseService)
	courseController.RegisterCourseRoutes(router)
	passportController := controller.NewPassportController(passportService)
	passportController.RegisterPassportRoutes(router)
	fileController := controller.NewFileController(fileService, userService)
	fileController.RegisterFileRoutes(router)

	//log.Fatal().Err(http.ListenAndServe(":8000", router))
	http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "access_token"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}), handlers.AllowedOrigins([]string{"*"}))(router))

}

func ConnectToDB(username string, password string, host string, port string, dbName string) (*gorm.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local", username, password, host, port, dbName)
	var err error
	var db *gorm.DB
	db, err = gorm.Open("mysql", dataSourceName)

	db.AutoMigrate(&model.User{}, &model.Passport{}, &model.Hobby{}, &model.Course{}, &model.File{})
	err2 := db.Debug().Model(&model.Passport{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").Error
	if err2 != nil {
		log.Error().Err(err2).Msg("")
	}
	err4 := db.Debug().Model(&model.Hobby{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").Error
	if err4 != nil {
		fmt.Println(err2)
	}
	// err5 := db.Debug().Model(&model.Publickey{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").Error
	// if err5 != nil {
	// 	log.Error().Err(err5).Msg("")
	// }
	return db, err
}
