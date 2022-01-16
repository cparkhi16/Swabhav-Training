package service

import (
	"app/model"
	"app/repository"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func (us *UserService) GetPassportIDForUser(ID uuid.UUID) model.Passport {
	uow := repository.NewUnitOfWork(us.DB, true)
	qp := repository.Filter("user_id = ?", ID)
	var p model.Passport
	err := us.Repo.GetFirst(uow, &p, qp)
	if err != nil {
		us.Logger.Error().Msg("No passport found for user")
	}
	return p
}
func (us *UserService) UpdatePassportDetailForUser(u *model.User) error {
	uow := repository.NewUnitOfWork(us.DB, false)
	//uc := *u
	//qp := repository.Filter("id = ?", uc.ID)
	//er := us.Repo.GetFirst(uow, &uc, qp)
	count := us.GetUsersCount("id = ?", u.ID)
	if count == 0 {
		us.Logger.Error().Msg("Could not find the user ")
		err := fmt.Errorf("could not find the user")
		return err
	}
	err := us.Repo.Update(uow, u)
	if err != nil {
		if !us.CheckIfPassportIDExists(&u.Passport) {
			uow.Complete()
			//fmt.Println("Error updating user's passport")
			us.Logger.Error().Msgf("Error updating user's passport %v", err)
			return err
		}
	} else {
		uow.Commit()
	}
	return nil
}
func (us *UserService) UpdatePassport(passport *model.Passport) error {
	uow := repository.NewUnitOfWork(us.DB, false)
	fmt.Println("Count of passport ", us.GetPassportCount(passport.ID))
	if us.GetPassportCount(passport.ID) == 0 {
		uow.Complete()
		err := fmt.Errorf("no passport id found ")
		return err
	}
	e := us.Repo.Update(uow, passport)
	if e != nil {
		return e
	}
	uow.Commit()
	return nil
}
func (us *UserService) FindAndDeletePassport(user *model.User) error {
	uow := repository.NewUnitOfWork(us.DB, false)
	p := us.GetPassportIDForUser(user.ID)
	fmt.Println("Id passport ", p.ID)
	if p.ID != uuid.Nil {
		e := us.Repo.Delete(uow, &p)
		if e != nil {
			uow.Complete()
			us.Logger.Error().Msg("Error deleting passport for user")
			return e
		}
		uow.Commit()
	}

	return nil
}
func (us *UserService) GetPassportCount(ID uuid.UUID) int {
	uow := repository.NewUnitOfWork(us.DB, true)
	var c int = 0
	err := us.Repo.GetCount(uow, model.Passport{}, &c, "id = ?", ID)
	if err != nil {
		us.Logger.Error().Msgf("error in getting passport count")
	}
	return c
}
func (us *UserService) CheckIfPassportIDExists(passport *model.Passport) bool {
	uow := repository.NewUnitOfWork(us.DB, true)
	p := *passport
	qp := repository.Filter("passport_id = ?", p.PassportID)
	err := us.Repo.GetFirst(uow, &p, qp)
	return err == nil
}
