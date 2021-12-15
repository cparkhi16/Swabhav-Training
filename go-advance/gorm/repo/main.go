package main

import (
	c "app/controller"
	m "app/model"
	r "app/repository"
	s "app/service"
	"fmt"
	"log"

	"github.com/gorilla/mux"
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
	db.AutoMigrate(&m.Passport{})
	db.Model(&m.Hobby{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&m.Passport{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	//uow := r.NewUnitOfWork(db, true)
	/*user := m.NewUser("Parth", "JKO")
	h := m.NewHobby("TV")
	user.AddHobbies(*h)
	j := m.NewHobby("Games")
	user.AddHobbies(*j)
	s.AddUser(uow, *user)*/

	//Using complete and commit
	//uw := r.NewUnitOfWork(db, false)
	//ud := m.NewUser("Ram", "KLJ")
	//s.AddUser(uw, ud)

	//Delete user ram with ID 124f7dba-a07b-45fc-8978-b2b885a20a37
	var um m.User
	ramId, _ := uuid.FromString("124f7dba-a07b-45fc-8978-b2b885a20a37")
	um.ID = ramId
	//s.DeleteUser(uw, &um)

	//Testing before create hook
	uow := r.NewUnitOfWork(db, true)

	user := m.NewUser("", "test")
	userService := s.NewUserService(uow)
	userService.AddUser(user)

	//uj := m.NewUser("Ritesh", "JIH")
	//userService.AddUser(uj)
	var Ritesh m.User
	ritID, _ := uuid.FromString("9d294f19-9332-4774-9889-a02c025a2424")
	Ritesh.ID = ritID
	p := m.NewPassport(100)
	Ritesh.Passport = *p
	ujHob := m.NewHobby("Gymming")
	Ritesh.Hobbies = append(Ritesh.Hobbies, *ujHob)

	//Assigning hobby to Rohit 532611a1-58cf-4e91-b276-aa4cd6338f8b
	var rohit m.User
	rId, _ := uuid.FromString("532611a1-58cf-4e91-b276-aa4cd6338f8b")
	rohit.ID = rId
	roH := m.NewHobby("Table tennis")
	rohit.Hobbies = append(rohit.Hobbies, *roH)
	//rHobby := m.Hobby{HobbyName: "Table tennis", TestModel: m.TestModel{ID: uuid.NewV4()}}
	//uj.Hobbies = append(uj.Hobbies, rHobby)

	//Testing after create hook
	//us := m.NewUser("Jay", "MHJ") //08607252-7917-4323-aa83-3a4f3ab30cbb
	//s.AddUser(uow, us)

	var jay m.User
	jayID, _ := uuid.FromString("08607252-7917-4323-aa83-3a4f3ab30cbb")
	jay.ID = jayID
	jay.Hobbies = append(jay.Hobbies, m.Hobby{HobbyName: "Cycling", TestModel: m.TestModel{ID: uuid.NewV4()}})

	var users []m.User
	association := []string{"Hobbies", "Courses"}
	userService.GetUsers(&users, association)
	var u m.User
	id, _ := uuid.FromString("0a27768d-9759-48c7-8b85-89152783bf76")
	userService.GetUserById(&u, id, association)

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
	courseService := s.NewCourseService(uow)

	//Getting course Java by ID 3852ce46-3f17-4e51-95ef-979893d31f0a
	var java m.Course
	javaId, _ := uuid.FromString("3852ce46-3f17-4e51-95ef-979893d31f0a")
	a := []string{}
	j := courseService.GetCourseById(&java, javaId, a)
	//fmt.Println("Java obj ", *j)

	var golang m.Course
	goId, _ := uuid.FromString("f9e9d37e-7d14-405a-9e48-39aa9d291ab6")
	b := []string{}
	gl := courseService.GetCourseById(&golang, goId, b)
	fmt.Println("Golang obj ", *gl)

	//jay.Courses = append(jay.Courses, j, gl)
	//rohit.Courses = append(rohit.Courses, j)
	Ritesh.Courses = append(Ritesh.Courses, j, gl)

	//userService.UpdateUser(&Ritesh)
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
	courseService.GetCourses(&courses, v)

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
	userService.GetUsersWithCourse(javaId, v)

	//Golang Course
	userService.GetUsersWithCourse(goId, v)

	//Getting user with name "Jay"
	fmt.Println("-------- Query processor -----")
	//userService.GetUser()
	//courseService.GetDetailsWithCourseID()

	router := mux.NewRouter()
	c.RegisterRoutesForUser(userService, router)

}
