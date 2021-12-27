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
	//fmt.Fprintf(w, "Here in create course")
	var newCourse m.Course
	er := json.NewDecoder(r.Body).Decode(&newCourse)
	if er != nil {
		cc.cs.Logger.Error().Msg("Error in course JSON decoding")
	}
	err := cc.cs.AddCourse(&newCourse)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
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
		cc.cs.Logger.Error().Msgf("Error in decoding course JSON", er)
	}
	params := mux.Vars(r)
	id, erID := uuid.FromString(params["id"])
	if erID == nil {
		if id != uuid.Nil {
			updateCourse.ID = id
			e := cc.cs.UpdateCourse(&updateCourse)
			if e != nil {
				cc.cs.Logger.Error().Msgf("Error updating course detail %v", e)
				fmt.Fprintf(w, e.Error())
			}
		} else {
			cc.cs.Logger.Error().Msg("Please enter a course ID in params")
		}
	} else {
		fmt.Fprintf(w, "Incorrect UUID ")
	}
}

func (cc *CourseController) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	var deleteCourse m.Course
	er := json.NewDecoder(r.Body).Decode(&deleteCourse)
	if er != nil {
		cc.cs.Logger.Error().Msgf("Error in decoding course JSON", er)
	}
	params := mux.Vars(r)
	id, erID := uuid.FromString(params["id"])
	if erID == nil {
		if id != uuid.Nil {
			deleteCourse.ID = id
			e := cc.cs.DeleteCourse(&deleteCourse)
			if e != nil {
				cc.cs.Logger.Error().Msg("Error deleting course")
				fmt.Fprintf(w, "Course already deleted")
			}
		} else {
			cc.cs.Logger.Error().Msg("Please enter a User ID in params")
		}
	} else {
		fmt.Fprintf(w, "Incorrect UUID ")
	}
}
