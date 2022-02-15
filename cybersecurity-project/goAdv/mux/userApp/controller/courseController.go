package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"userPassport/model"
	"userPassport/service"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type CourseController struct {
	courseService *service.CourseService
}

func NewCourseController(courseService *service.CourseService) *CourseController {
	return &CourseController{
		courseService: courseService,
	}
}

func (cc *CourseController) RegisterCourseRoutes(router *mux.Router) {
	cc.courseService.Logger.Info().Msg("Registered Course Routes")
	subRoute := router.PathPrefix("/courses").Subrouter()
	subRoute.HandleFunc("/", cc.getCourses).Methods("GET")
	subRoute.HandleFunc("/{userId}", cc.getCourseByUserId).Methods("GET")
	subRoute.HandleFunc("/", cc.getCourses).Queries("limit", "{limit:[0-9]+}", "pageNo", "{pageNo:[0-9]+}").Methods("GET")
	subRoute.HandleFunc("/", cc.createCourse).Methods("POST")
	subRoute.HandleFunc("/{id}", cc.updateCourse).Methods("PUT")
	//subRoute.HandleFunc("/{id}", cc.getCourse).Methods("GET")
	subRoute.HandleFunc("/{id}", cc.deleteCourse).Methods("DELETE")
}

func (cc *CourseController) getCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Course-Count", strconv.Itoa(cc.courseService.GetCoursesCount()))
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	pageNo, _ := strconv.Atoi(r.FormValue("pageNo"))
	offset := limit * (pageNo - 1)
	cc.courseService.Logger.Info().Int("limit", limit).Int("pageNo", pageNo).Msg("Get All Courses")
	var courses []model.Course
	cc.courseService.GetAllCourses(&courses, limit, offset)
	json.NewEncoder(w).Encode(courses)
}

func (cc *CourseController) getCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Course-Count", strconv.Itoa(cc.courseService.GetCoursesCount()))
	params := mux.Vars(r)
	id, err := uuid.FromString(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("ID invalid!")
		return
	}
	var str []string
	var course model.Course
	cc.courseService.GetCourseById(&course, id, str)
	if course.ID == uuid.Nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Could not get course!")
		return
	}
	json.NewEncoder(w).Encode(course)
}

func (cc *CourseController) getCourseByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Course-Count", strconv.Itoa(cc.courseService.GetCoursesCount()))
	params := mux.Vars(r)
	id, err := uuid.FromString(params["userId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("ID invalid!")
		return
	}
	var str []string
	var course model.Course
	cc.courseService.GetCourseById(&course, id, str)
	if course.ID == uuid.Nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Could not get course!")
		return
	}
	json.NewEncoder(w).Encode(course)
}

func (cc *CourseController) createCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Course-Count", strconv.Itoa(cc.courseService.GetCoursesCount()))
	var course model.Course
	json.NewDecoder(r.Body).Decode(&course)
	err := cc.courseService.CreateCourse(&course)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error in creating course")
		return
	}
	json.NewEncoder(w).Encode(course)
}

func (cc *CourseController) updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Course-Count", strconv.Itoa(cc.courseService.GetCoursesCount()))
	params := mux.Vars(r)
	id, err := uuid.FromString(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("ID invalid!")
		return
	}
	var updatedCourse model.Course
	json.NewDecoder(r.Body).Decode(&updatedCourse)
	updatedCourse.ID = id
	if !cc.courseService.CheckIfCourseExists(id) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Could not get course!")
		return
	}
	fmt.Println(updatedCourse)
	if !cc.courseService.CheckIfCourseExistsByName(updatedCourse.Name, updatedCourse.ID) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("duplicate course error!")
		return
	}
	err2 := cc.courseService.UpdateCourse(updatedCourse)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error in updating course")
		return
	}
	json.NewEncoder(w).Encode("Course updated")
}

func (cc *CourseController) deleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Course-Count", strconv.Itoa(cc.courseService.GetCoursesCount()))
	params := mux.Vars(r)
	idToDelete, err := uuid.FromString(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("ID invalid!")
		return
	}
	if !cc.courseService.CheckIfCourseExists(idToDelete) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Could not get course!")
		return
	}
	err2 := cc.courseService.DeleteCourse(idToDelete)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error in deleting course")
		return
	}
	json.NewEncoder(w).Encode("Course deleted")
}

/*
func (cc *CourseController) doesCourseExist(id uuid.UUID) bool {
	var course model.Course
	var str []string
	cc.courseService.GetCourseById(&course, id, str)
	if course.ID == uuid.Nil {
		return false
	}
	return true
}*/
