package main

import (
	"app/controller"
	zerologger "app/logger"
	"app/model"
	"app/repository"
	"app/service"
	"log"
	"net/http"

	//"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	/* depends_on vs links , scratch*/
	dbConn := "root:hello@tcp(127.0.0.1:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	//dbConn:="root:hello@tcp(userdb)/userdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dbConn)
	if err != nil {
		//fmt.Println("Err is ",err)
		log.Fatal(err)
	}
	logger := zerologger.GetLogger()
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Hobby{})
	db.AutoMigrate(&model.Course{})
	db.AutoMigrate(&model.Passport{})
	db.Model(&model.Hobby{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&model.Passport{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	repo := repository.NewRepository()
	user := model.NewUser("", "", "test")
	userService := service.NewUserService(repo, db, logger)
	userService.AddUser(user)

	courseService := service.NewCourseService(repo, db, logger)
	//http://localhost:9000/login?email=SP@fp.com&password=sahil
	/* {
	    "Email":"jjk@fp.com",
	    "Password":"jjk"
	}*/
	router := mux.NewRouter()
	authRoute := router.MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		return true
	}).Subrouter()
	nonAuthRoute := router.MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		return true
	}).Subrouter()

	userController := controller.NewUserController(userService)
	courseController := controller.NewCourseController(courseService)
	hobbyController := controller.NewHobbyController(userService)
	passportController := controller.NewPassportController(userService)
	hobbyController.RegisterRoutesForHobby(authRoute, nonAuthRoute)
	userController.RegisterRoutesForUser(authRoute, nonAuthRoute)
	courseController.RegisterRoutesForCourse(authRoute, nonAuthRoute)
	passportController.RegisterRoutesForPassport(authRoute, nonAuthRoute)
	logger.Info().Msgf("Starting server")
	log.Fatal(http.ListenAndServe(":9000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Token"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}), handlers.AllowedOrigins([]string{"http://localhost:4200"}))(router)))

}
