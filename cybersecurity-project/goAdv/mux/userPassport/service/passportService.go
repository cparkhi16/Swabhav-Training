package service

import (
	"fmt"
	"userPassport/repository"

	"userPassport/model"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type PassportService struct {
	Repo repository.Repository
	DB   *gorm.DB
}

func NewPassportService(Repo repository.Repository, db *gorm.DB) *PassportService {
	return &PassportService{
		Repo: Repo,
		DB:   db,
	}
}

func (p *PassportService) GetPassportByUserId(out *model.Passport, userId uuid.UUID) error {
	unit := repository.NewUnitOfWork(p.DB, true)
	var queryp []repository.QueryProcessor
	queryp = append(queryp, repository.Filter("user_id=?", userId))
	err := p.Repo.GetFirst(unit, out, queryp)
	fmt.Println(out)
	if err != nil {
		unit.Complete()
		return err
	}
	unit.Commit()
	return nil
}

func (s *UserService) GetAllPassports(out *[]model.Passport, limit int, offset int) error {
	unit := repository.NewUnitOfWork(s.DB, true)
	var queryp []repository.QueryProcessor
	var count int
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
