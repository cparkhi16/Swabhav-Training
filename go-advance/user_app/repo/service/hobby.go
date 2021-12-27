package service

import (
	"app/model"
	"app/repository"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func (us *UserService) GetHobbyCount(condition string, value interface{}) int {
	uow := repository.NewUnitOfWork(us.DB, true)
	var c int = 0
	err := us.Repo.GetCount(uow, model.Hobby{}, &c, condition, value)
	if err != nil {
		us.Logger.Error().Msgf("error in getting hobby count")
	}
	//fmt.Println("Count --", c)
	return c
}
func (us *UserService) GetHobbyById(h *model.Hobby, ID uuid.UUID) model.Hobby {
	uow := repository.NewUnitOfWork(us.DB, true)
	p := []string{}
	err := us.Repo.GetAllForTenant(uow, h, ID, p)
	if err != nil {
		us.Logger.Error().Msg("Error in getting hobby by id ")
	}
	//o := out.(*m.Hobby)
	return *h
}

func (us *UserService) DeleteHobbyById(h *model.Hobby) error {
	uow := repository.NewUnitOfWork(us.DB, false)
	//ho := *h
	//qp := repository.Filter("id = ?", ho.ID)
	//err := us.Repo.GetFirst(uow, &ho, qp)
	count := us.GetHobbyCount("id = ?", h.ID)
	if count == 0 {
		us.Logger.Error().Msg("Could not find the hobby ")
		err := fmt.Errorf("could not find the hobby")
		return err
	}
	er := us.Repo.Delete(uow, h)
	if er != nil {
		uow.Complete()
		us.Logger.Error().Msg("Error in deleting hobby by ID")
		return er
	}
	uow.Commit()
	return nil
}
func (us *UserService) UpdateHobbyById(h *model.Hobby) error {
	uow := repository.NewUnitOfWork(us.DB, false)
	//ho := *h
	//qp := repository.Filter("id = ?", ho.ID)
	//err := us.Repo.GetFirst(uow, &ho, qp)
	count := us.GetHobbyCount("id = ?", h.ID)
	fmt.Println(count)
	if count == 0 {
		us.Logger.Error().Msg("Could not find the hobby ")
		err := fmt.Errorf("could not find the hobby")
		return err
	}
	er := us.Repo.Update(uow, h)
	if er != nil {
		uow.Complete()
		us.Logger.Error().Msg("Error in updatingg hobby by ID")
		return er
	}
	uow.Commit()
	return nil
}

func (us *UserService) GetAllHobbiesWithPagination(page, limit int, hobbies *[]model.Hobby) []model.Hobby {
	uow := repository.NewUnitOfWork(us.DB, true)
	offset := (page - 1) * limit
	queryLimit := repository.Limit(limit)
	queryOffset := repository.Offset(offset)
	qp := []repository.QueryProcessor{queryLimit, queryOffset}
	result := us.Repo.GetAllWithQueryProcessor(uow, hobbies, qp)
	if result != nil {
		us.Logger.Error().Msgf("Error while get all hobbies with pagination %v", result)
		return nil
	}
	//o := out.(*[]m.Hobby)
	return *hobbies
}
