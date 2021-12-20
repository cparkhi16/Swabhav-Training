package service

import (
	m "app/model"
	re "app/repository"

	uuid "github.com/satori/go.uuid"
)

func (us *UserService) GetPassportIDForUser(ID uuid.UUID) m.Passport {
	uow := re.NewUnitOfWork(us.DB, true)
	qp := re.Filter("user_id = ?", ID)
	var p m.Passport
	err := us.Repo.GetFirst(uow, &p, qp)
	if err != nil {
		us.Logger.Error().Msg("No passport found for user")
	}
	return p
}
func (us *UserService) UpdatePassportDetailForUser(entity interface{}) error {
	uow := re.NewUnitOfWork(us.DB, false)
	err := us.Repo.Update(uow, entity)
	if err != nil {
		uow.Complete()
		//fmt.Println("Error updating user's passport")
		us.Logger.Error().Msgf("Error updating user's passport %v", err)
		return err
	} else {
		uow.Commit()
	}
	return nil
}

func (us *UserService) FindAndDeletePassport(entity interface{}, preloadAssociations []string) error {
	//fmt.Println("----------------")
	uow := re.NewUnitOfWork(us.DB, false)
	pqp := re.PreloadAssociations(preloadAssociations)
	user := entity.(*m.User)
	err := us.Repo.GetFirst(uow, &user, pqp)
	if err != nil {
		us.Logger.Error().Msgf("Error finding user %v", err)
		return err
	} else {
		o := entity.(*m.User)
		var p m.Passport
		p.ID = o.Passport.ID
		//fmt.Println("Passport ID for user", p.ID)
		zeroUUID, _ := uuid.FromString("00000000-0000-0000-0000-000000000000")
		if p.ID != zeroUUID {
			e := us.Repo.Delete(uow, &p)
			if e != nil {
				uow.Complete()
				//fmt.Println("Error deleting passport for user", e)
				us.Logger.Error().Msg("Error deleting passport for user")
				return e
			}
			uow.Commit()
		}
	}
	return nil
}
