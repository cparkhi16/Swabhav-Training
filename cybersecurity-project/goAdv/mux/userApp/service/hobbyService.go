package service

import (
	"userPassport/model"
	"userPassport/repository"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/zerolog"
	uuid "github.com/satori/go.uuid"
)

type HobbyService struct {
	Repo   repository.Repository
	DB     *gorm.DB
	Logger *zerolog.Logger
}

var preloadForHobbies = []string{}

func NewHobbyService(Repo repository.Repository, db *gorm.DB, logger *zerolog.Logger) *HobbyService {
	return &HobbyService{
		Repo:   Repo,
		DB:     db,
		Logger: logger,
	}
}

func (s *HobbyService) GetAllHobbies(out *[]model.Hobby, limit int, offset int) error {
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
	s.Logger.Info().Msg("Get all hobbies")
	unit.Commit()
	return nil
}

func (s *HobbyService) GetHobbyById(out *model.Hobby, id uuid.UUID) error {
	unit := repository.NewUnitOfWork(s.DB, true)
	err := s.Repo.Get(unit, out, id, preloadForHobbies, "id")
	if err != nil {
		unit.Complete()
		return err
	}
	s.Logger.Info().Interface("ID-", id).Msg("Get hobby by ID")
	unit.Commit()
	return nil
}

func (s *HobbyService) GetHobbyByUserId(out *[]model.Hobby, userId uuid.UUID) error {
	unit := repository.NewUnitOfWork(s.DB, true)
	var queryp []repository.QueryProcessor
	queryp = append(queryp, repository.Filter("user_id=?", userId))
	err := s.Repo.GetAll(unit, out, queryp)
	if err != nil {
		unit.Complete()
		return err
	}
	s.Logger.Info().Interface("UserId-", userId).Msg("Get hobby by UserId")
	unit.Commit()
	return nil
}

func (s *HobbyService) UpdateHobby(entity model.Hobby) error {
	unit := repository.NewUnitOfWork(s.DB, false)
	err := s.Repo.Update(unit, entity)
	if err != nil {
		unit.Complete()
		return err
	}
	s.Logger.Info().Interface("hobby-", entity).Msg("Update Hobby")
	unit.Commit()
	return nil
}

func (s *HobbyService) AddHobby(entity model.Hobby) error {
	unit := repository.NewUnitOfWork(s.DB, false)
	err := s.Repo.Add(unit, entity)
	if err != nil {
		unit.Complete()
		return err
	}
	s.Logger.Info().Interface("hobby-", entity).Msg("Add Hobby")
	unit.Commit()
	return nil
}

func (s *HobbyService) GetHobbyCount() int {
	unit := repository.NewUnitOfWork(s.DB, true)
	var hobbies []model.Hobby
	var count int
	var queryp []repository.QueryProcessor
	s.Repo.GetCount(unit, hobbies, &count, queryp)
	s.Logger.Info().Int("count-", count).Msg("Get Hobbies Count")
	return count
}

func (s *HobbyService) DeleteHobby(hobbyId uuid.UUID) error {
	unit := repository.NewUnitOfWork(s.DB, false)
	hobbyToDelete := model.Hobby{Base: model.Base{ID: hobbyId}}
	err := s.Repo.Delete(unit, &hobbyToDelete)
	if err != nil {
		unit.Complete()
		return err
	}
	s.Logger.Info().Interface("ID-", hobbyId).Msg("Delete Hobby by ID")
	unit.Commit()
	return nil
}

func (s *HobbyService) CheckIfHobbyExists(id uuid.UUID) bool {
	unit := repository.NewUnitOfWork(s.DB, true)
	var hobbies []model.Hobby
	var count int
	var queryp = []repository.QueryProcessor{repository.Filter("id=?", id)}
	s.Repo.GetCount(unit, hobbies, &count, queryp)
	//db.Model(&users).Count(&count)
	s.Logger.Info().Int("count-", count).Msg("Check if hobby exists")
	if count == 0 {
		return false
	}
	return true
}
