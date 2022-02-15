package service

import (
	"fmt"
	"userPassport/model"
	"userPassport/repository"

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

func (s *UserService) CreateUser(newUser *model.User) error {
	unit := repository.NewUnitOfWork(s.DB, false)
	err := s.Repo.Add(unit, newUser)
	if err != nil {
		unit.Complete()
		return err
	}
	unit.Commit()
	return nil
}

func (s *UserService) GetAllUsers(out *[]model.User, limit int, offset int) error {
	unit := repository.NewUnitOfWork(s.DB, true)
	var queryp []repository.QueryProcessor
	var count int
	var preload []string
	preload = append(preload, "Passport")
	queryp = append(queryp, repository.PreloadAssociations(preload))
	if limit != 0 {
		queryp = append(queryp, repository.Paginate(limit, offset, &count))
	}
	fmt.Println(count)
	err := s.Repo.GetAll(unit, out, queryp)
	if err != nil {
		unit.Complete()
		return err
	}
	unit.Commit()
	return nil
}

func (s *UserService) GetUsersCount() int {
	unit := repository.NewUnitOfWork(s.DB, true)
	db := unit.DB
	var users []model.User
	var count int
	db.Model(&users).Count(&count)
	return count
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
	userToDelete := model.User{ID: userId}
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
	err := s.Repo.GetFirst(unit, out, queryp)
	if err != nil {
		unit.Complete()
		return err
	}
	unit.Commit()
	return nil
}
