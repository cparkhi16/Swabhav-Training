package service

import (
	"app/model"
	"app/repository"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog"
)

type CourseService struct {
	Repo   repository.Repository
	DB     *gorm.DB
	Logger *zerolog.Logger
}

func NewCourseService(r repository.Repository, DB *gorm.DB, l *zerolog.Logger) *CourseService {
	return &CourseService{Repo: r, DB: DB, Logger: l}
}

func (cs *CourseService) GetCourseCount(condition string, value interface{}) int {
	uow := repository.NewUnitOfWork(cs.DB, true)
	var c int = 0
	err := cs.Repo.GetCount(uow, model.Course{}, &c, condition, value)
	if err != nil {
		cs.Logger.Error().Msgf("error in getting course count")
	}
	//fmt.Println("Count --", c)
	return c
}
func (cs *CourseService) AddCourse(c *model.Course) error {
	uow := repository.NewUnitOfWork(cs.DB, false)
	count := cs.GetCourseCount("name = ?", c.Name)
	fmt.Println("Courses  count ", count)
	if count != 0 {
		err := fmt.Errorf("a course with same name already exists")
		return err
	}
	err := cs.Repo.Add(uow, c)
	if err != nil {
		uow.Complete()
		//	fmt.Println("Error while adding course")
		cs.Logger.Error().Msgf("Error adding course %v", err)
		return err
	}
	uow.Commit()
	return nil
}
func (cs *CourseService) UpdateCourse(course *model.Course) error {
	uow := repository.NewUnitOfWork(cs.DB, false)
	coursesWithSameName := cs.GetCourseCount("name = ?", course.Name)
	if coursesWithSameName > 0 {
		cs.Logger.Error().Msg("Course with same name already exists in db")
		err := fmt.Errorf("course with same name already exists in db")
		return err
	}
	count := cs.GetCourseCount("id = ?", course.ID)
	if count == 0 {
		cs.Logger.Error().Msg("Could not find course")
		err := fmt.Errorf("could not find course")
		return err
	}
	er := cs.Repo.Update(uow, course)
	if er != nil {
		uow.Complete()
		cs.Logger.Error().Msgf("Error updating course %v ", er)
		return er
	}
	uow.Commit()
	return nil
}

func (cs *CourseService) DeleteCourse(course *model.Course) error {
	uow := repository.NewUnitOfWork(cs.DB, false)
	count := cs.GetCourseCount("id = ?", course.ID)
	if count == 0 {
		cs.Logger.Error().Msg("Could not find course")
		err := fmt.Errorf("could not find course")
		return err
	}
	err := cs.Repo.Delete(uow, course)
	if err != nil {
		uow.Complete()
		cs.Logger.Error().Msgf("Error deleting course %v", err)
		return err
	}
	we := cs.Repo.ClearAssociation(uow, course, "users")
	if we != nil {
		uow.Complete()
		cs.Logger.Error().Msgf("Error while deleting associated users %v ", we)
		return we
	}
	uow.Commit()
	return nil
}

func (cs *CourseService) GetAllCoursesWithPagination(page, limit int, courses *[]model.Course) []model.Course {
	uow := repository.NewUnitOfWork(cs.DB, true)
	offset := (page - 1) * limit
	queryLimit := repository.Limit(limit)
	queryOffset := repository.Offset(offset)
	qp := []repository.QueryProcessor{queryLimit, queryOffset}
	result := cs.Repo.GetAllWithQueryProcessor(uow, courses, qp)
	if result != nil {
		cs.Logger.Error().Msgf("Error while getting all courses with pagination %v", result)
		return nil
	}
	//o := out.(*[]m.Course)
	return *courses
}
