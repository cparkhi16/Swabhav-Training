package service

import (
	"app/model"
	"app/repository"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

type CourseService struct {
	Repo repository.Repository
	DB   *gorm.DB
}

func NewCourseService(Repo repository.Repository, db *gorm.DB) *CourseService {
	return &CourseService{
		Repo: Repo,
		DB:   db,
	}
}

func (s *CourseService) CreateCourse(name string) error {
	unit := repository.NewUnitOfWork(s.DB, false)
	newCourse := model.Course{
		CustomModel: model.CustomModel{
			ID:        uuid.NewV4(),
			CreatedBy: "yogesh",
			CreatedAt: time.Now(),
		},
		Name: name,
	}
	err := s.Repo.Add(unit, newCourse)
	if err != nil {
		unit.Complete()
		return err
	}
	unit.Commit()
	return nil
}

func (s *CourseService) GetAllCourses(out *[]model.Course, preloadAssociations []string) error {
	unit := repository.NewUnitOfWork(s.DB, true)
	err := s.Repo.GetAll(unit, out, preloadAssociations)
	if err != nil {
		unit.Complete()
		return err
	}
	unit.Commit()
	return nil
}

func (s *CourseService) GetCourseById(out *model.Course, id uuid.UUID, preloadAssociations []string) error {
	unit := repository.NewUnitOfWork(s.DB, true)
	err := s.Repo.Get(unit, out, id, preloadAssociations, "id")
	if err != nil {
		unit.Complete()
		return err
	}
	unit.Commit()
	return nil
}

func (s *CourseService) UpdateCourse(entity model.Course) error {
	unit := repository.NewUnitOfWork(s.DB, false)
	err := s.Repo.Update(unit, entity)
	if err != nil {
		unit.Complete()
		return err
	}
	unit.Commit()
	return nil
}

func (s *CourseService) DeleteCourse(courseId uuid.UUID) error {
	unit := repository.NewUnitOfWork(s.DB, false)
	courseToDelete := model.Course{CustomModel: model.CustomModel{ID: courseId}}
	err := s.Repo.Delete(unit, &courseToDelete)
	if err != nil {
		unit.Complete()
		return err
	}
	unit.Commit()
	return nil
}

func (s *CourseService) GetCourseFromName(out *model.Course, name string) error {
	db := s.DB
	return db.Table("courses").Where("name=?", name).Find(out).Error
}
