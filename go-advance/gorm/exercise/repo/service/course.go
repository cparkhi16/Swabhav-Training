package service

import (
	m "app/model"
	r "app/repository"
	re "app/repository"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func AddCourse(uow *r.UnitOfWork, c m.Course) {
	r := re.NewRepository()
	e := r.Add(uow, c)
	if e != nil {
		fmt.Println("Error while adding course")
	}
}

func GetCourseById(uow *r.UnitOfWork, out interface{}, tenantID uuid.UUID, preloadAssociations []string) *m.Course {
	r := re.NewRepository()
	err := r.GetAllForTenant(uow, out, tenantID, preloadAssociations)
	if err != nil {
		fmt.Println("Error in get all user ", err)
	}
	o := out.(*m.Course)

	fmt.Println(o.Name)
	return o

}

func GetCourses(uow *r.UnitOfWork, out interface{}, preloadAssociations []string) {
	r := re.NewRepository()
	err := r.GetAll(uow, out, preloadAssociations)
	if err != nil {
		fmt.Println("Error in get all courses ", err)
	}
	o := out.(*[]m.Course)
	for _, val := range *o {
		fmt.Println("Courses available in db", val.Name)
	}
}

func UpdateCourse(uow *r.UnitOfWork, entity interface{}) {
	r := re.NewRepository()
	err := r.Update(uow, entity)
	if err != nil {
		fmt.Println("Error updating course")
	}
}

func DeleteCourse(uow *r.UnitOfWork, entity interface{}) {
	r := re.NewRepository()
	err := r.Delete(uow, entity)
	if err != nil {
		fmt.Println("Error deleting course")
	}

}
func GetDetailsWithCourseID(uow *r.UnitOfWork) {
	r := re.NewRepository()
	var c m.Course
	cID, _ := uuid.FromString("3852ce46-3f17-4e51-95ef-979893d31f0a")
	qp := re.Filter("id = ?", cID)
	//qps := []re.QueryProcessor{}
	//qps = append(qps, qp)
	err := r.GetFirst(uow, &c, qp)
	if err != nil {
		fmt.Println("Error using quey processor")
	}
	fmt.Println("Course object from db --- ", c)
}
