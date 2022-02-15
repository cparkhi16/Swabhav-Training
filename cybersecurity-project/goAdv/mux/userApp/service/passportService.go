package service

import (
	"fmt"
	"userPassport/repository"

	"userPassport/model"

	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog"
	uuid "github.com/satori/go.uuid"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type PassportService struct {
	Repo   repository.Repository
	DB     *gorm.DB
	Logger *zerolog.Logger
}

var preloadForPassport = []string{}

func NewPassportService(Repo repository.Repository, db *gorm.DB, logger *zerolog.Logger) *PassportService {
	return &PassportService{
		Repo:   Repo,
		DB:     db,
		Logger: logger,
	}
}

func (p *PassportService) GetPassportByUserId(out *model.Passport, userId uuid.UUID) error {
	unit := repository.NewUnitOfWork(p.DB, true)
	var queryp []repository.QueryProcessor
	queryp = append(queryp, repository.Filter("user_id=?", userId))
	err := p.Repo.GetFirst(unit, out, queryp)
	p.Logger.Info().Interface("userID-", userId).Msg("Get Passport By UserId")
	if err != nil {
		unit.Complete()
		return err
	}
	unit.Commit()
	return nil
}

func (s *PassportService) GetAllPassports(out *[]model.Passport, limit int, offset int, userId uuid.UUID) error {
	unit := repository.NewUnitOfWork(s.DB, true)
	var queryp []repository.QueryProcessor
	var count int
	if limit != 0 {
		queryp = append(queryp, repository.Paginate(limit, offset, &count))
	}
	queryp = append(queryp, repository.Filter("user_id=?", userId))
	fmt.Println(count)
	err := s.Repo.GetAll(unit, out, queryp)
	if err != nil {
		unit.Complete()
		return err
	}
	s.Logger.Info().Msg("Get all passports")
	unit.Commit()
	return nil
}

func (s *PassportService) GetPassportById(out *model.Passport, id uuid.UUID) error {
	unit := repository.NewUnitOfWork(s.DB, true)
	err := s.Repo.Get(unit, out, id, preloadForPassport, "id")
	if err != nil {
		unit.Complete()
		return err
	}
	s.Logger.Info().Interface("ID-", id).Msg("Get passport by ID")
	unit.Commit()
	return nil
}

func (s *PassportService) UpdatePassport(entity model.Passport) error {
	unit := repository.NewUnitOfWork(s.DB, false)
	err := s.Repo.Update(unit, entity)
	if err != nil {
		unit.Complete()
		return err
	}
	s.Logger.Info().Interface("passport-", entity).Msg("Update Passport")
	unit.Commit()
	return nil
}

func (s *PassportService) DeletePassport(passportId uuid.UUID) error {
	unit := repository.NewUnitOfWork(s.DB, false)
	passportToDelete := model.Passport{Base: model.Base{ID: passportId}}
	err := s.Repo.Delete(unit, &passportToDelete)
	if err != nil {
		unit.Complete()
		return err
	}
	s.Logger.Info().Interface("ID-", passportId).Msg("Delete Passport by ID")
	unit.Commit()
	return nil
}

func (s *PassportService) DeletePassportByUserId(userId uuid.UUID) error {
	unit := repository.NewUnitOfWork(s.DB, false)
	fmt.Println(userId)
	passportToDelete := model.Passport{UserId: userId}
	err := s.Repo.Delete(unit, &passportToDelete)
	if err != nil {
		unit.Complete()
		return err
	}
	s.Logger.Info().Interface("user ID-", userId).Msg("Delete Passport by userID")
	unit.Commit()
	return nil
}

func (s *PassportService) GetPassportCount() int {
	unit := repository.NewUnitOfWork(s.DB, true)
	var passports []model.Passport
	var count int
	var queryp []repository.QueryProcessor
	s.Repo.GetCount(unit, passports, &count, queryp)
	s.Logger.Info().Int("count-", count).Msg("Get Passports Count")
	return count
}

func (s *PassportService) CheckIfPassportExists(id uuid.UUID) bool {
	unit := repository.NewUnitOfWork(s.DB, true)
	var passports []model.Passport
	var count int
	var queryp = []repository.QueryProcessor{repository.Filter("id=?", id)}
	s.Repo.GetCount(unit, passports, &count, queryp)
	//db.Model(&users).Count(&count)
	s.Logger.Info().Int("count-", count).Msg("Check if passport exists")
	if count == 0 {
		return false
	}
	return true
}
