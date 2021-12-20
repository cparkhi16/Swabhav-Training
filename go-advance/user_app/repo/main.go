package main

import (
	c "app/controller"
	lr "app/logger"
	m "app/model"
	r "app/repository"
	s "app/service"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	dbConn := "root:hello@tcp(127.0.0.1:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dbConn)
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}
	logger := lr.GetLogger()
	db.AutoMigrate(&m.User{})
	db.AutoMigrate(&m.Hobby{})
	db.AutoMigrate(&m.Course{})
	db.AutoMigrate(&m.Passport{})
	db.Model(&m.Hobby{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&m.Passport{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	repo := r.NewRepository()
	user := m.NewUser("", "", "test")
	userService := s.NewUserService(repo, db, logger)
	userService.AddUser(user)

	courseService := s.NewCourseService(repo, db, logger)
	//http://localhost:9000/users/token?email=rk@fp.com&password=Role23
	router := mux.NewRouter()
	userController := c.NewUserController(userService)
	courseController := c.NewCourseController(courseService)
	hobbyController := c.NewHobbyController(userService)
	c.RegisterRoutesForHobby(hobbyController, router)
	c.RegisterRoutesForUser(userController, router)
	c.RegisterRoutesForCourse(courseController, router)
	logger.Info().Msgf("Starting server")
	log.Fatal(http.ListenAndServe(":9000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"POST", "PUT", "DELETE"}), handlers.AllowedOrigins([]string{"abc.com"}))(router)))

}
