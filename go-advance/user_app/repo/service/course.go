package service

import (
	lr "app/logger"
	m "app/model"
	re "app/repository"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

var logger = lr.GetLogger()

type CourseService struct {
	Repo re.Repository
	DB   *gorm.DB
}

func NewCourseService(r re.Repository, DB *gorm.DB) *CourseService {
	return &CourseService{Repo: r, DB: DB}
}
func (cs *CourseService) AddCourse(c *m.Course) {
	uow := re.NewUnitOfWork(cs.DB, false)
	e := cs.Repo.Add(uow, c)
	if e != nil {
		uow.Complete()
		//	fmt.Println("Error while adding course")
		logger.Error().Msgf("Error adding course %v", e)
	} else {
		uow.Commit()
	}
}

func (cs *CourseService) GetCourseById(out interface{}, tenantID uuid.UUID, preloadAssociations []string) *m.Course {
	uow := re.NewUnitOfWork(cs.DB, true)
	err := cs.Repo.GetAllForTenant(uow, out, tenantID, preloadAssociations)
	if err != nil {
		fmt.Println("Error in get all user ", err)
	}
	o := out.(*m.Course)

	fmt.Println(o.Name)
	return o

}

func (cs *CourseService) GetCourses(out interface{}, preloadAssociations []string) {
	uow := re.NewUnitOfWork(cs.DB, true)
	err := cs.Repo.GetAll(uow, out, preloadAssociations)
	if err != nil {
		fmt.Println("Error in get all courses ", err)
	}
	o := out.(*[]m.Course)
	for _, val := range *o {
		fmt.Println("Courses available in db", val.Name)
	}
}

func (cs *CourseService) UpdateCourse(entity interface{}) error {
	uow := re.NewUnitOfWork(cs.DB, false)
	err := cs.Repo.Update(uow, entity)
	if err != nil {
		uow.Complete()
		//fmt.Println("Error updating course")
		logger.Error().Msgf("Error updating course %v ", err)
		return err
	} else {
		uow.Commit()
	}
	return nil
}

func (cs *CourseService) DeleteCourse(entity interface{}) error {
	uow := re.NewUnitOfWork(cs.DB, false)
	err := cs.Repo.Delete(uow, entity)
	if err != nil {
		uow.Complete()
		//fmt.Println("Error deleting course")
		logger.Error().Msgf("Error deleting course %v", err)
		return err
	} else {
		uow.Commit()
	}
	return nil
}
func (cs *CourseService) GetDetailsWithCourseID() {
	uow := re.NewUnitOfWork(cs.DB, true)
	var c m.Course
	cID, _ := uuid.FromString("3852ce46-3f17-4e51-95ef-979893d31f0a")
	qp := re.Filter("id = ?", cID)
	//qps := []re.QueryProcessor{}
	//qps = append(qps, qp)
	err := cs.Repo.GetFirst(uow, &c, qp)
	if err != nil {
		fmt.Println("Error using quey processor")
	}
	fmt.Println("Course object from db --- ", c)
}

func (cs *CourseService) GetAllCoursesWithPagination(page, limit int, out interface{}) []m.Course {
	uow := re.NewUnitOfWork(cs.DB, true)
	offset := (page - 1) * limit
	queryLimit := re.Limit(limit)
	queryOffset := re.Offset(offset)
	qp := []re.QueryProcessor{queryLimit, queryOffset}
	result := cs.Repo.GetAllWithQueryProcessor(uow, out, qp)
	//res := result.Debug().Preload("Hobbies").Find(out)
	if result != nil {
		log.Fatal("Error in pagination for courses ")
		return nil
	}
	o := out.(*[]m.Course)
	return *o
}
