package main

import (
	m "app/model"
	r "app/repository"
	s "app/service"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

func main() {
	dbConn := "root:hello@tcp(127.0.0.1:3306)/repo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dbConn)
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}
	db.AutoMigrate(&m.User{})
	db.AutoMigrate(&m.Hobby{})
	db.AutoMigrate(&m.Course{})
	db.Model(&m.Hobby{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

	uow := r.NewUnitOfWork(db, true)
	/*user := m.NewUser("Parth", "JKO")
	h := m.NewHobby("TV")
	user.AddHobbies(*h)
	j := m.NewHobby("Games")
	user.AddHobbies(*j)
	s.AddUser(uow, *user)*/

	//Testing before create hook
	user := m.NewUser("", "test")
	s.AddUser(uow, user)

	//Testing after create hook
	//us := m.NewUser("Jay", "MHJ") //08607252-7917-4323-aa83-3a4f3ab30cbb
	//s.AddUser(uow, us)

	var jay m.User
	jayID, _ := uuid.FromString("08607252-7917-4323-aa83-3a4f3ab30cbb")
	jay.ID = jayID
	jay.Hobbies = append(jay.Hobbies, m.Hobby{HobbyName: "Cycling", TestModel: m.TestModel{ID: uuid.NewV4()}})

	var users []m.User
	association := []string{"Hobbies", "Courses"}
	s.GetUsers(uow, &users, association)
	var u m.User
	id, _ := uuid.FromString("0a27768d-9759-48c7-8b85-89152783bf76")
	s.GetUserById(uow, &u, id, association)

	/*userMap := make(map[string]interface{})
	userMap["id"] = id
	userMap["name"] = "Raju"*/

	//fmt.Println("id", id)
	//u.ID = id
	//u.Name = "Chinmay"
	//hId, _ := uuid.FromString("492427b4-b4e8-4b7f-969d-de2632fca33d")
	//u.Hobbies[0] = m.Hobby{HobbyName: "Trekking", TestModel: m.TestModel{ID: hId}}
	//s.UpdateUser(uow, &u)

	//Courses
	//c := m.NewCourse("Java")
	//s.AddCourse(uow, *c)
	//cd := m.NewCourse("Python")
	//s.AddCourse(uow, *cd)
	//g := m.NewCourse("Golang")
	//s.AddCourse(uow, *g)

	//Getting course Java by ID 3852ce46-3f17-4e51-95ef-979893d31f0a
	var java m.Course
	javaId, _ := uuid.FromString("3852ce46-3f17-4e51-95ef-979893d31f0a")
	a := []string{}
	j := s.GetCourseById(uow, &java, javaId, a)
	fmt.Println("Java obj ", *j)

	var golang m.Course
	goId, _ := uuid.FromString("f9e9d37e-7d14-405a-9e48-39aa9d291ab6")
	b := []string{}
	gl := s.GetCourseById(uow, &golang, goId, b)
	fmt.Println("Golang obj ", *gl)

	jay.Courses = append(jay.Courses, j, gl)
	//s.UpdateUser(uow, &jay) Checking before and update hooks

	//Assigning java to user with ID 0a27768d-9759-48c7-8b85-89152783bf76
	//u.Courses = append(u.Courses, j)
	//s.UpdateUser(uow, &u)

	//Deleting a user and its associated hobby 9ceeeafe-6050-42e0-a712-908db71c54cc
	var deleteUser m.User
	deleteUserID, _ := uuid.FromString("9ceeeafe-6050-42e0-a712-908db71c54cc")
	deleteUser.ID = deleteUserID
	//s.DeleteUser(uow, &deleteUser)

	// user got assigned a course in person_courses but course table didn't updated with this course
	/*h := m.NewCourse("Hadoop")
	fmt.Println(h)
	u.Courses = append(u.Courses, h)
	s.UpdateUser(uow, &u)*/

	//Get all courses
	var courses []m.Course
	v := []string{}
	s.GetCourses(uow, &courses, v)

	// Assigning two courses to parth
	var ut m.User
	utId, _ := uuid.FromString("9ceeeafe-6050-42e0-a712-908db71c54cc")
	ut.ID = utId
	ut.Courses = append(ut.Courses, j, gl)
	//s.UpdateUser(uow, &ut)

	//Deleting parth with his hobby and course
	//s.DeleteUser(uow, &ut)

	//Updating course name (golang to Golang_course)
	golang.Name = "Golang_course"
	//s.UpdateCourse(uow, &golang)

	//New course ML
	//ml := m.NewCourse("ML")
	//s.AddCourse(uow, *ml)

	//Delete course ml
	mlID, _ := uuid.FromString("509aee8c-ea8c-45c8-99fd-dada7ca7e17e")
	var deleteML m.Course
	deleteML.ID = mlID
	//s.DeleteCourse(uow, &deleteML)

	//Java course
	s.GetUsersWithCourse(uow, javaId, v)

	//Golang Course
	s.GetUsersWithCourse(uow, goId, v)

	//Getting user with name "Jay"
	fmt.Println("-------- Query processor -----")
	s.GetUser(uow)
	s.GetDetailsWithCourseID(uow)
}
