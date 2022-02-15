package main

import (
	"fmt"

	"app/model"
	"app/repository"
	"app/service"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func main() {
	db, _ := ConnectToDB("root", "Panda@19", "localhost", "3306", "goApp")
	/*
		err := db.AutoMigrate(&model.User{}, &model.Course{}, &model.Hobby{}).Error
		if err != nil {
			fmt.Println(err)
		}
		err2 := db.Debug().Model(&model.Hobby{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").Error
		if err2 != nil {
			fmt.Println(err2)
		}*/
	repo1 := repository.NewRepository()
	fmt.Println(repo1)
	userService := service.NewUserService(repo1, db)
	fmt.Println(userService)
	courseService := service.NewCourseService(repo1, db)
	fmt.Println(userService, courseService)

	var user model.User
	var queryp []repository.QueryProcessor
	var str1 = []string{"Courses", "Hobbies"}
	queryp = append(queryp, repository.PreloadAssociations(str1))
	queryp = append(queryp, repository.Wherefunc())
	userService.GetFirstUser(&user, queryp)
	fmt.Println(user)
	/*
		var courses []model.Course
		var pythonCourse model.Course
		courseService.GetCourseFromName(&pythonCourse, "go")
		courses = append(courses, pythonCourse)
		h2 := model.Hobby{HobbyId: uuid.NewV4(), Name: "brushing"}
		h3 := model.Hobby{HobbyId: uuid.NewV4(), Name: "bathing"}
		var hobbies = []model.Hobby{h2, h3}
		userService.CreateUser("monster", "hell", courses, hobbies)

		var users []model.User
		var str1 = []string{"Courses", "Hobbies"}
		userService.GetAllUsers(&users, str1)
		fmt.Println(users)

		var courses2 []model.Course
		var str2 []string
		courseService.GetAllCourses(&courses2, str2)
		fmt.Println(courses2)

		var user model.User
		var queryp []repository.QueryProcessor
		queryp = append(queryp, repository.PreloadAssociations(str1))
		queryp = append(queryp, repository.Wherefunc())
		userService.GetFirstUser(&user, queryp)
		fmt.Println(user)*/

	/*
			Create Table and Adding Foreign Key
			err := db.Debug().CreateTable(&model.User{}, &model.Course{}, &model.Hobby{}).Error
			if err != nil {
				fmt.Println(err)
			}
			err2 := db.Debug().Model(&model.Hobby{}).AddForeignKey("user_id", "users(user_id)", "CASCADE", "CASCADE").Error
			if err2 != nil {
				fmt.Println(err2)
			}

		unit1 := repository.NewUnitOfWork(db, true)
		repo1 := repository.NewRepository()
		fmt.Println(unit1, repo1)

		userService := service.NewUserService(repo1, unit1)
		fmt.Println(userService)
		courseService := service.NewCourseService(repo1, unit1)
		fmt.Println(userService, courseService)*/

	/*User-Service Functions-->

	1)CreateUser-

	var courses []model.Course
	var pythonCourse model.Course
	courseService.GetCourseFromName(&pythonCourse, "python")
	courses = append(courses, pythonCourse)
	h2 := model.Hobby{HobbyId: uuid.NewV4(), Name: "eating"}
	h3 := model.Hobby{HobbyId: uuid.NewV4(), Name: "painting"}
	var hobbies = []model.Hobby{h2, h3}
	err := userService.CreateUser(uuid.NewV4(), "sam", "sanfran", courses, hobbies)
	if err != nil {
		fmt.Println(err)
	}

		2)UpdateUser-
		uuid1, _ := uuid.FromString("d6907509-507d-4a81-9689-3fa551c29d91")
		var userToBeUpdated model.User
		var str = []string{"Courses", "Hobbies"}
		userService.GetUserById(&userToBeUpdated, uuid1, str)
		userToBeUpdated.Name = "shae"
		err := userService.UpdateUser(userToBeUpdated)
		if err != nil {
			fmt.Println(err)
		}

		3)GetAllUsers and GetUserById-
		var str1 = []string{"Courses", "Hobbies"}
		var users []model.User
		userService.GetAllUsers(&users, str1)
		fmt.Println(users)

		uuid1, _ := uuid.FromString("d6907509-507d-4a81-9689-3fa551c29d91")

		var user1 model.User
		userService.GetUserById(&user1, uuid1, str1)
		fmt.Println(user1)

		4)DeleteUser-
		uuid1, _ := uuid.FromString("277b0d40-1445-427f-9f84-00d8a5c99625")
		userToDelete := model.User{UserId: uuid1}
		err := userService.DeleteUser(&userToDelete)
		if err != nil {
			fmt.Println(err)
		}

	*/

	/*Course-Service functions-->
	1)CreateCourse-
	err=courseService.CreateCourse(uuid.NewV4(),"courseName")
	if err!=nil{
		fmt.Println(err)
	}

	2)GetAllCourses and GetCourseById-
	var courses []model.Course
	var str2 = []string{}
	courseService.GetAllCourses(&courses, str2)
	fmt.Println(courses)

	var course model.Course
	var str3 []string
	uuidToFind, _ := uuid.FromString("af4a6102-13a8-4c99-bdf4-4620e14aab02")
	courseService.GetCourseById(&course, uuidToFind, str3)
	fmt.Println(course)

	3)UpdateCourse-
	uuid1,_:=uuid.FromString("760425c7-8eab-449f-98a1-fed23e6514ac")
	courseToUpdate:=Course{CourseId:uuid1,Name:"newName"}
	err=courseService.UpdateCourse(courseToUpdate)
	if err != nil {
		fmt.Println(err)
	}

	4)DeleteCourse-
	uuid1,_:=uuid.FromString("760425c7-8eab-449f-98a1-fed23e6514ac")
	err=courseService.DeleteCourse(uuid1)
	if err != nil {
		fmt.Println(err)
	}
	*/

}

func ConnectToDB(username string, password string, host string, port string, dbName string) (*gorm.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local", username, password, host, port, dbName)
	var err error
	db, err = gorm.Open("mysql", dataSourceName)
	return db, err
}
