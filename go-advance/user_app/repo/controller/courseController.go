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
		w.WriteHeader(http.StatusBadRequest)
		cc.cs.Logger.Error().Msg("Error in course JSON decoding")
	}
	err := cc.cs.AddCourse(&newCourse)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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
		courses = cc.cs.GetCourses(&courses)
	} else {
		pageInt, _ = strconv.Atoi(page)
		limitInt, _ = strconv.Atoi(limit)
		if limitInt < 0 && pageInt >= 2 {
			return
		}
		courses = cc.cs.GetAllCoursesWithPagination(pageInt, limitInt, &courses)
	}
	//http://localhost:9000/courses?limit=2&page=1
	w.Header().Set("Course-Count", strconv.Itoa(len(courses)))
	fmt.Println("Courses slice len ", len(courses))
	json.NewEncoder(w).Encode(courses)

}
func (cc *CourseController) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	var updateCourse m.Course
	er := json.NewDecoder(r.Body).Decode(&updateCourse)
	if er != nil {
		w.WriteHeader(http.StatusBadRequest)
		cc.cs.Logger.Error().Msgf("Error in decoding course JSON", er)
	}
	params := mux.Vars(r)
	id, erID := uuid.FromString(params["id"])
	if erID == nil {
		if id != uuid.Nil {
			updateCourse.ID = id
			e := cc.cs.UpdateCourse(&updateCourse)
			if e != nil {
				w.WriteHeader(http.StatusBadRequest)
				cc.cs.Logger.Error().Msgf("Error updating course detail %v", e)
				fmt.Fprintf(w, e.Error())
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			cc.cs.Logger.Error().Msg("Please enter a course ID in params")
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Incorrect UUID ")
	}
}

func (cc *CourseController) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	var deleteCourse m.Course
	er := json.NewDecoder(r.Body).Decode(&deleteCourse)
	if er != nil {
		w.WriteHeader(http.StatusBadRequest)
		cc.cs.Logger.Error().Msgf("Error in decoding course JSON", er)
	}
	params := mux.Vars(r)
	hardDelete := r.FormValue("hardDelete")
	id, erID := uuid.FromString(params["id"])
	if erID == nil {
		if id != uuid.Nil {
			deleteCourse.ID = id
			e := cc.cs.DeleteCourse(&deleteCourse, hardDelete)
			if e != nil {
				w.WriteHeader(http.StatusBadRequest)
				cc.cs.Logger.Error().Msg("Error deleting course")
				fmt.Fprintf(w, "Course already deleted")
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			cc.cs.Logger.Error().Msg("Please enter a User ID in params")
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Incorrect UUID ")
	}
}
