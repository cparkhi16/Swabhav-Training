package service

import (
	"userPassport/model"
	"userPassport/repository"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/zerolog"
	uuid "github.com/satori/go.uuid"
)

type CourseService struct {
	Repo   repository.Repository
	DB     *gorm.DB
	Logger *zerolog.Logger
}

func NewCourseService(Repo repository.Repository, db *gorm.DB, logger *zerolog.Logger) *CourseService {
	return &CourseService{
		Repo:   Repo,
		DB:     db,
		Logger: logger,
	}
}

func (s *CourseService) CreateCourse(newCourse *model.Course) error {
	unit := repository.NewUnitOfWork(s.DB, false)
	err := s.Repo.Add(unit, newCourse)
	if err != nil {
		unit.Complete()
		return err
	}

	s.Logger.Info().Interface("course-", newCourse).Msg("Create New Course")
	unit.Commit()
	return nil
}

func (s *CourseService) GetAllCourses(out *[]model.Course, limit int, offset int) error {
	unit := repository.NewUnitOfWork(s.DB, true)
	var queryp []repository.QueryProcessor
	var count int
	if limit != 0 {
		queryp = append(queryp, repository.Paginate(limit, offset, &count))
	}
	s.Logger.Info().Int("count", count)
	err := s.Repo.GetAll(unit, out, queryp)
	if err != nil {
		unit.Complete()
		return err
	}
	s.Logger.Info().Msg("Get all courses")
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
	s.Logger.Info().Interface("ID-", id).Msg("Get course by ID")
	unit.Commit()
	return nil
}

func (s *CourseService) GetCoursesByUserId(out *model.Course, id uuid.UUID, preloadAssociations []string) error {
	unit := repository.NewUnitOfWork(s.DB, true)
	err := s.Repo.Get(unit, out, id, preloadAssociations, "user_id")
	if err != nil {
		unit.Complete()
		return err
	}
	s.Logger.Info().Interface("ID-", id).Msg("Get course by ID")
	unit.Commit()
	return nil
}

func (s *CourseService) GetCoursesCount() int {
	unit := repository.NewUnitOfWork(s.DB, true)
	var courses []model.Course
	var count int
	var queryp []repository.QueryProcessor
	s.Repo.GetCount(unit, courses, &count, queryp)
	s.Logger.Info().Int("count-", count).Msg("Get Courses Count")
	return count
}

func (s *CourseService) UpdateCourse(entity model.Course) error {
	unit := repository.NewUnitOfWork(s.DB, false)
	err := s.Repo.Update(unit, entity)
	if err != nil {
		unit.Complete()
		return err
	}
	s.Logger.Info().Interface("course-", entity).Msg("Update Course")
	unit.Commit()
	return nil
}

func (s *CourseService) DeleteCourse(courseId uuid.UUID) error {
	unit := repository.NewUnitOfWork(s.DB, false)
	courseToDelete := model.Course{Base: model.Base{ID: courseId}}
	err := s.Repo.Delete(unit, &courseToDelete)
	if err != nil {
		unit.Complete()
		return err
	}
	s.Logger.Info().Interface("ID-", courseId).Msg("Delete Course by ID")
	unit.Commit()
	return nil
}

func (s *CourseService) GetCourseFromName(out *model.Course, name string) error {
	unit := repository.NewUnitOfWork(s.DB, false)
	var queryp []repository.QueryProcessor
	queryp = append(queryp, repository.Table("courses"))
	queryp = append(queryp, repository.Filter("name=?", name))
	err := s.Repo.GetFirst(unit, out, queryp)
	if err != nil {
		unit.Complete()
		return err
	}
	unit.Commit()
	return nil
	//return db.Table("courses").Where("name=?", name).Find(out).Error
}

func (s *CourseService) CheckIfCourseExists(id uuid.UUID) bool {
	unit := repository.NewUnitOfWork(s.DB, true)
	var courses []model.Course
	var count int
	var queryp = []repository.QueryProcessor{repository.Filter("id=?", id)}
	s.Repo.GetCount(unit, courses, &count, queryp)
	//db.Model(&users).Count(&count)
	s.Logger.Info().Int("count-", count).Msg("Check if course exists")
	if count == 0 {
		return false
	}
	return true
}

func (s *CourseService) CheckIfCourseExistsByName(courseName string, courseId uuid.UUID) bool {
	unit := repository.NewUnitOfWork(s.DB, true)
	var courses []model.Course
	var count int
	var queryp = []repository.QueryProcessor{repository.Filter("NOT id=?", courseId), repository.Filter("name=?", courseName)}
	s.Repo.GetCount(unit, courses, &count, queryp)
	//fmt.Println(count)
	//db.Model(&users).Count(&count)
	s.Logger.Info().Int("count-", count).Msg("Check if course exists by name")
	if count > 0 {
		return false
	}
	return true
}
