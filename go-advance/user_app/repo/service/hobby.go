package service

import (
	m "app/model"
	re "app/repository"
	"log"

	uuid "github.com/satori/go.uuid"
)

func (us *UserService) GetHobbyById(out interface{}, ID uuid.UUID) m.Hobby {
	uow := re.NewUnitOfWork(us.DB, true)
	p := []string{}
	err := us.Repo.GetAllForTenant(uow, out, ID, p)
	if err != nil {
		logger.Error().Msg("Error in getting hobby by id ")
	}
	o := out.(*m.Hobby)
	return *o
}

func (us *UserService) DeleteHobbyById(out interface{}) error {
	uow := re.NewUnitOfWork(us.DB, false)
	err := us.Repo.Delete(uow, out)
	if err != nil {
		uow.Complete()
		logger.Error().Msg("Error in deleting hobby by ID")
		return err
	}
	uow.Commit()
	return nil
}
func (us *UserService) UpdateHobbyById(out interface{}) error {
	uow := re.NewUnitOfWork(us.DB, false)
	err := us.Repo.Update(uow, out)
	if err != nil {
		uow.Complete()
		logger.Error().Msg("Error in updatingg hobby by ID")
		return err
	}
	uow.Commit()
	return nil
}

func (us *UserService) GetAllHobbiesWithPagination(page, limit int, out interface{}) []m.Hobby {
	uow := re.NewUnitOfWork(us.DB, true)
	offset := (page - 1) * limit
	queryLimit := re.Limit(limit)
	queryOffset := re.Offset(offset)
	qp := []re.QueryProcessor{queryLimit, queryOffset}
	result := us.Repo.GetAllWithQueryProcessor(uow, out, qp)
	if result != nil {
		log.Fatal("Error in pagination for courses ")
		return nil
	}
	o := out.(*[]m.Hobby)
	return *o
}
