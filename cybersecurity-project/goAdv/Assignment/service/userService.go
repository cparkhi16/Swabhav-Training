package service

import (
	"app/model"
	"app/repository"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserService struct {
	Repo repository.Repository
	DB   *gorm.DB
}

func NewUserService(Repo repository.Repository, db *gorm.DB) *UserService {
	return &UserService{
		Repo: Repo,
		DB:   db,
	}
}

func (s *UserService) CreateUser(name string, address string, courses []model.Course, hobbies []model.Hobby) error {
	unit := repository.NewUnitOfWork(s.DB, false)
	emptyUuid, _ := uuid.FromString("")
	newUser := model.User{
		CustomModel: model.CustomModel{
			ID:        emptyUuid,
			CreatedBy: "yogesh",
			CreatedAt: time.Now(),
		},
		Name:    name,
		Address: address,
		Courses: courses,
		Hobbies: hobbies,
	}
	err := s.Repo.Add(unit, &newUser)
	if err != nil {
		unit.Complete()
		return err
	}
	unit.Commit()
	return nil
}

func (s *UserService) GetAllUsers(out *[]model.User, preloadAssociations []string) error {
	unit := repository.NewUnitOfWork(s.DB, true)
	err := s.Repo.GetAll(unit, out, preloadAssociations)
	if err != nil {
		unit.Complete()
		return err
	}
	unit.Commit()
	return nil
}

func (s *UserService) GetUserById(out *model.User, id uuid.UUID, preloadAssociations []string) error {
	unit := repository.NewUnitOfWork(s.DB, true)
	err := s.Repo.Get(unit, out, id, preloadAssociations, "id")
	if err != nil {
		unit.Complete()
		return err
	}
	unit.Commit()
	return nil
}

func (s *UserService) UpdateUser(entity model.User) error {
	unit := repository.NewUnitOfWork(s.DB, false)
	err := s.Repo.Update(unit, entity)
	if err != nil {
		unit.Complete()
		return err
	}
	unit.Commit()
	return nil
}

func (s *UserService) DeleteUser(userId uuid.UUID) error {
	unit := repository.NewUnitOfWork(s.DB, false)
	userToDelete := model.User{CustomModel: model.CustomModel{ID: userId}}
	err := s.Repo.Delete(unit, &userToDelete)
	if err != nil {
		unit.Complete()
		return err
	}
	unit.Commit()
	return nil
}

func (s *UserService) GetFirstUser(out *model.User, queryp []repository.QueryProcessor) error {
	unit := repository.NewUnitOfWork(s.DB, true)
	//var queryp []repository.QueryProcessor
	//queryp = append(queryp, repository.Wherefunc())
	err := s.Repo.GetFirst(unit, out, queryp)
	if err != nil {
		unit.Complete()
		return err
	}
	unit.Commit()
	return nil
}
