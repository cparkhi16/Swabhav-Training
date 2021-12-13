package service

import (
	m "app/model"
	re "app/repository"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type CourseService struct {
	uow *re.UnitOfWork
}

func NewCourseService(uow *re.UnitOfWork) *CourseService {
	return &CourseService{uow: uow}
}
func (cs *CourseService) AddCourse(c m.Course) {
	r := re.NewRepository()
	e := r.Add(cs.uow, c)
	if e != nil {
		cs.uow.Complete()
		fmt.Println("Error while adding course")
	} else {
		cs.uow.Commit()
	}
}

func (cs *CourseService) GetCourseById(out interface{}, tenantID uuid.UUID, preloadAssociations []string) *m.Course {
	r := re.NewRepository()
	err := r.GetAllForTenant(cs.uow, out, tenantID, preloadAssociations)
	if err != nil {
		fmt.Println("Error in get all user ", err)
	}
	o := out.(*m.Course)

	fmt.Println(o.Name)
	return o

}

func (cs *CourseService) GetCourses(out interface{}, preloadAssociations []string) {
	r := re.NewRepository()
	err := r.GetAll(cs.uow, out, preloadAssociations)
	if err != nil {
		fmt.Println("Error in get all courses ", err)
	}
	o := out.(*[]m.Course)
	for _, val := range *o {
		fmt.Println("Courses available in db", val.Name)
	}
}

func (cs *CourseService) UpdateCourse(entity interface{}) {
	r := re.NewRepository()
	err := r.Update(cs.uow, entity)
	if err != nil {
		cs.uow.Complete()
		fmt.Println("Error updating course")
	} else {
		cs.uow.Commit()
	}
}

func (cs *CourseService) DeleteCourse(entity interface{}) {
	r := re.NewRepository()
	err := r.Delete(cs.uow, entity)
	if err != nil {
		cs.uow.Complete()
		fmt.Println("Error deleting course")
	} else {
		cs.uow.Commit()
	}

}
func (cs *CourseService) GetDetailsWithCourseID() {
	r := re.NewRepository()
	var c m.Course
	cID, _ := uuid.FromString("3852ce46-3f17-4e51-95ef-979893d31f0a")
	qp := re.Filter("id = ?", cID)
	//qps := []re.QueryProcessor{}
	//qps = append(qps, qp)
	err := r.GetFirst(cs.uow, &c, qp)
	if err != nil {
		fmt.Println("Error using quey processor")
	}
	fmt.Println("Course object from db --- ", c)
}
