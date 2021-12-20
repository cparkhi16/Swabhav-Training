package controller

import (
	m "app/model"
	s "app/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type CourseController struct {
	cs *s.CourseService
}

func NewCourseController(cs *s.CourseService) *CourseController {
	return &CourseController{cs: cs}
}

func (cc *CourseController) CreateCourse(w http.ResponseWriter, r *http.Request) {
	var newCourse m.Course
	er := json.NewDecoder(r.Body).Decode(&newCourse)
	if er != nil {
		logger.Error().Msg("Error in course JSON decoding")
	}
	cc.cs.AddCourse(&newCourse)
}
func (cc *CourseController) GetAllCourses(w http.ResponseWriter, r *http.Request) {
	var courses []m.Course
	w.Header().Set("Content-Type", "application/json")
	// query params
	page := r.FormValue("page")
	limit := r.FormValue("limit")
	var pageInt int
	var limitInt int
	if page == "" || limit == "" {
		pageInt = 1
		limitInt = 2
	} else {
		pageInt, _ = strconv.Atoi(page)
		limitInt, _ = strconv.Atoi(limit)
	}
	//http://localhost:9000/courses?limit=2&page=1
	courses = cc.cs.GetAllCoursesWithPagination(pageInt, limitInt, &courses)
	w.Header().Set("Course-Count", strconv.Itoa(len(courses)))
	fmt.Println("Courses slice len ", len(courses))
	json.NewEncoder(w).Encode(courses)

}
func (cc *CourseController) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	var updateCourse m.Course
	er := json.NewDecoder(r.Body).Decode(&updateCourse)
	if er != nil {
		logger.Error().Msgf("Error in decoding course JSON", er)
	}
	params := mux.Vars(r)
	id, _ := uuid.FromString(params["id"])
	zeroUUID, _ := uuid.FromString("00000000-0000-0000-0000-000000000000")
	if id != zeroUUID {
		updateCourse.ID = id
		e := cc.cs.UpdateCourse(&updateCourse)
		if e != nil {
			logger.Error().Msgf("Error updating user detail %v", e)
		}
	} else {
		logger.Error().Msg("Please enter a User ID in params")
	}
}

func (cc *CourseController) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	var deleteCourse m.Course
	er := json.NewDecoder(r.Body).Decode(&deleteCourse)
	if er != nil {
		logger.Error().Msgf("Error in decoding course JSON", er)
	}
	params := mux.Vars(r)
	id, _ := uuid.FromString(params["id"])
	zeroUUID, _ := uuid.FromString("00000000-0000-0000-0000-000000000000")
	if id != zeroUUID {
		deleteCourse.ID = id
		e := cc.cs.DeleteCourse(&deleteCourse)
		if e != nil {
			logger.Error().Msg("Error deleting course")
		}
	} else {
		logger.Error().Msg("Please enter a User ID in params")
	}
}
