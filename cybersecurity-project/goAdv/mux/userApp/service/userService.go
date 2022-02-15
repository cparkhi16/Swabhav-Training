package service

import (
	"fmt"
	"userPassport/model"
	"userPassport/repository"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/zerolog"
)

type UserService struct {
	Repo   repository.Repository
	DB     *gorm.DB
	Logger *zerolog.Logger
}

var preload = []string{"Passport", "Hobbies", "Courses"}

func NewUserService(Repo repository.Repository, db *gorm.DB, logger *zerolog.Logger) *UserService {
	return &UserService{
		Repo:   Repo,
		DB:     db,
		Logger: logger,
	}
}

func (s *UserService) CreateUser(newUser *model.User) error {
	unit := repository.NewUnitOfWork(s.DB, false)
	var courses []model.Course
	for _, course := range newUser.Courses {
		var c model.Course
		var queryp []repository.QueryProcessor
		queryp = append(queryp, repository.Filter("name=?", course.Name))
		err := s.Repo.GetFirst(unit, &c, queryp)
		if err != nil {
			unit.Complete()
			return err
		}
		//s.DB.Table("courses").Where("name=?", course.Name).Find(&c)
		courses = append(courses, c)
	}
	newUser.Courses = courses
	err := s.Repo.Add(unit, newUser)
	if err != nil {
		unit.Complete()
		return err
	}
	s.Logger.Info().Interface("user-", newUser).Msg("Create New User")
	unit.Commit()
	return nil
}

func (s *UserService) GetAllUsers(out *[]model.User, limit int, offset int) error {
	unit := repository.NewUnitOfWork(s.DB, true)
	var queryp []repository.QueryProcessor
	var count int
	queryp = append(queryp, repository.PreloadAssociations(preload))
	queryp = append(queryp, repository.Paginate(limit, offset, &count))
	//fmt.Println(queryp)
	err := s.Repo.GetAll(unit, out, queryp)
	if err != nil {
		unit.Complete()
		return err
	}
	s.Logger.Info().Msg("Get all users")
	unit.Commit()
	return nil
}

func (s *UserService) GetUsersCount() int {
	unit := repository.NewUnitOfWork(s.DB, true)
	//db := unit.DB
	var users []model.User
	var count int
	var queryp []repository.QueryProcessor
	s.Repo.GetCount(unit, users, &count, queryp)
	//db.Model(&users).Count(&count)
	s.Logger.Info().Int("count-", count).Msg("Get Users Count")
	return count
}

func (s *UserService) CheckIfUserExists(id uuid.UUID) bool {
	unit := repository.NewUnitOfWork(s.DB, true)
	var users []model.User
	var count int
	var queryp = []repository.QueryProcessor{repository.Filter("id=?", id)}
	s.Repo.GetCount(unit, users, &count, queryp)
	//db.Model(&users).Count(&count)
	s.Logger.Info().Int("count-", count).Msg("Check if user exists")
	if count == 0 {
		return false
	}
	return true
}

func (s *UserService) GetUserById(out *model.User, id uuid.UUID) error {
	unit := repository.NewUnitOfWork(s.DB, true)
	err := s.Repo.Get(unit, out, id, preload, "id")
	if err != nil {
		unit.Complete()
		return err
	}
	s.Logger.Info().Interface("ID-", id).Msg("Get user by ID")
	unit.Commit()
	return nil
}

func (s *UserService) UpdateUser(entity model.User) error {
	unit := repository.NewUnitOfWork(s.DB, false)
	var courses []model.Course
	for _, course := range entity.Courses {
		var c model.Course
		var queryp []repository.QueryProcessor
		queryp = append(queryp, repository.Filter("name=?", course.Name))
		err := s.Repo.GetFirst(unit, &c, queryp)
		if err != nil {
			unit.Complete()
			return err
		}
		//s.DB.Table("courses").Where("name=?", course.Name).Find(&c)
		courses = append(courses, c)
	}
	entity.Courses = courses
	err := s.Repo.Update(unit, entity)
	if err != nil {
		unit.Complete()
		return err
	}
	s.Logger.Info().Interface("user-", entity).Msg("Update User")
	unit.Commit()
	return nil
}

func (s *UserService) DeleteUser(userId uuid.UUID) error {
	unit := repository.NewUnitOfWork(s.DB, false)
	//userToDelete := model.User{Base: model.Base{ID: userId}}
	var userToDelete model.User
	s.GetUserById(&userToDelete, userId)
	//fmt.Println(userToDelete)
	//fmt.Println("hereeee")
	// err2 := unit.DB.Debug().Model(&model.User{Base: model.Base{ID: userId}}).Association("Courses").Delete(&userToDelete).Error
	// if err2 != nil {
	// 	fmt.Println("err2", err2)
	// 	return err2
	// }

	//fmt.Println(userToDelete)
	err := s.Repo.Delete(unit, &userToDelete)
	if err != nil {
		unit.Complete()
		return err
	}
	err2 := unit.DB.Model(&userToDelete).Association("Courses").Clear().Error
	if err2 != nil {
		fmt.Println(err2)
		return err2
	}
	s.Logger.Info().Interface("ID-", userId).Msg("Delete User by ID")
	unit.Commit()
	return nil
}

func (s *UserService) GetFirstUser(out *model.User, queryp []repository.QueryProcessor) error {
	unit := repository.NewUnitOfWork(s.DB, true)
	err := s.Repo.GetFirst(unit, out, queryp)
	if err != nil {
		unit.Complete()
		return err
	}
	s.Logger.Info().Msg("Get first user")
	unit.Commit()
	return nil
}

func (s *UserService) GetUserFromEmail(out *model.User, email string) error {
	unit := repository.NewUnitOfWork(s.DB, true)
	var queryp []repository.QueryProcessor
	queryp = append(queryp, repository.PreloadAssociations(preload))
	queryp = append(queryp, repository.Filter("email=?", email))
	err := s.Repo.GetFirst(unit, out, queryp)
	if err != nil {
		unit.Complete()
		return err
	}
	s.Logger.Info().Str("Email-", email).Msg("Get user by Email")
	unit.Commit()
	return nil
}

func (s *UserService) DeleteUserCourse(out *model.User, courseId string) error {
	id, _ := uuid.FromString(courseId)
	course := model.Course{Base: model.Base{ID: id}}
	return s.DB.Debug().Model(out).Association("Courses").Delete(course).Error
}
