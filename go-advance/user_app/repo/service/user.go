package service

import (
	"app/hash"
	"app/model"
	"app/repository"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog"
	uuid "github.com/satori/go.uuid"
)

type UserService struct {
	Logger *zerolog.Logger
	Repo   repository.Repository
	DB     *gorm.DB
}

func NewUserService(r repository.Repository, DB *gorm.DB, l *zerolog.Logger) *UserService {
	return &UserService{Repo: r, DB: DB, Logger: l}
}
func (us *UserService) AddUser(u *model.User) error {
	fmt.Println("Here in add user service")
	uow := repository.NewUnitOfWork(us.DB, false)
	hashedPassword := hash.CreateHashForPassword(u.Password)
	u.Password = hashedPassword
	if !us.CheckIfPassportIDExists(&u.Passport) {
		err := us.Repo.Add(uow, u)
		if err != nil {
			us.Logger.Error().Msgf("Error in add user %v", err)
			return err
		}
		uow.Commit()
	}
	return nil
}
func (us *UserService) GetUsersCount(condition string, value interface{}) int {
	uow := repository.NewUnitOfWork(us.DB, true)
	var c int = 0
	err := us.Repo.GetCount(uow, model.User{}, &c, condition, value)
	if err != nil {
		us.Logger.Error().Msgf("error in getting user count")
	}
	return c
}

func (us *UserService) GetUsers(users *[]model.User, preloadAssociations []string) []model.User {
	uow := repository.NewUnitOfWork(us.DB, true)
	pre := repository.PreloadAssociations(preloadAssociations)
	qp := []repository.QueryProcessor{pre}
	err := us.Repo.GetAllWithQueryProcessor(uow, users, qp)
	if err != nil {
		us.Logger.Error().Msgf("Error in get all user %v ", err)
	}
	for _, val := range *users {
		fmt.Println(val.FirstName)
		for _, val := range val.Hobbies {
			fmt.Println("Hobby ---", val.HobbyName)
		}
		for _, vl := range val.Courses {
			fmt.Println("Courses ---", vl.Name)
		}
	}
	return *users
}

func (us *UserService) GetUserById(user *model.User, tenantID uuid.UUID, preloadAssociations []string) (model.User, error) {
	uow := repository.NewUnitOfWork(us.DB, true)
	err := us.Repo.GetAllForTenant(uow, user, tenantID, preloadAssociations)
	if err != nil {
		us.Logger.Error().Msg("Error in get user by ID ")
		return model.User{}, err
	}
	return *user, nil
}

func (us *UserService) GetUserHobbies(user *model.User) []model.Hobby {
	p := []string{"Hobbies"}
	var hobbiesForUser []model.Hobby
	userByID, err := us.GetUserById(user, user.ID, p)
	if err == nil {
		for _, val := range userByID.Hobbies {
			hobbiesForUser = append(hobbiesForUser, val)
		}
	}
	return hobbiesForUser
}

func (us *UserService) UpdateUser(user *model.User) error {
	uow := repository.NewUnitOfWork(us.DB, false)
	count := us.GetUsersCount("id = ?", user.ID)
	if count == 0 || user.Password == "" {
		us.Logger.Error().Msg("Could not find the user or user password is empty")
		err := fmt.Errorf("could not find the user or user password is empty")
		return err
	}
	currUser := *user
	dbUser, _ := us.GetUserById(&currUser, currUser.ID, []string{})
	if dbUser.Email != user.Email {
		userWithSameEmail := us.GetUsersCount("email = ?", user.Email)
		if userWithSameEmail > 0 {
			err := fmt.Errorf("user exists with same email id")
			return err
		}
	}
	hashedPassword := hash.CreateHashForPassword(user.Password)
	user.Password = hashedPassword
	if !us.CheckIfPassportIDExists(&user.Passport) {
		er := us.Repo.Update(uow, user)
		if er != nil {
			uow.Complete()
			us.Logger.Error().Msgf("Error while updating user %v", er)
			return er
		}
		uow.Commit()
	}
	fmt.Println("Passport ID ", user.Passport.ID)
	if user.Passport.ID == uuid.Nil {
		us.FindAndDeletePassport(user)
	}
	return nil
}

func (us *UserService) DeleteUser(user *model.User) error {
	uow := repository.NewUnitOfWork(us.DB, false)
	count := us.GetUsersCount("id = ?", user.ID)
	if count == 0 {
		us.Logger.Error().Msg("Could not find the user ")
		err := fmt.Errorf("could not find the user")
		return err
	}
	er := us.Repo.Delete(uow, user)
	if er != nil {
		uow.Complete()
		//fmt.Println("Error deleting user")
		us.Logger.Error().Msgf("Error deleting user %v", er)
	} else {
		var h []model.Hobby
		e := uow.DB.Debug().Where("user_id = ?", user.ID).Find(&h).Error
		if len(h) != 0 {
			if e != nil {
				us.Logger.Error().Msg("Error finding associated hobbies for user")
			} else {
				for _, val := range h {
					ef := us.Repo.Delete(uow, &val)
					if ef != nil {
						uow.Complete()
						us.Logger.Error().Msgf("Error deleting hobby for this user %v", ef)
					}
				}
			}
		}
		we := us.Repo.ClearAssociation(uow, user, "courses")
		if we != nil {
			uow.Complete()
			us.Logger.Error().Msgf("Error while deleting associated courses %v ", we)
		}
		us.FindAndDeletePassport(user)
		uow.Commit()
	}
	return nil
}

func (us *UserService) GetAllUsersWithPagination(page, limit int, hobby []string, users *[]model.User) []model.User {
	uow := repository.NewUnitOfWork(us.DB, true)
	offset := (page - 1) * limit
	queryLimit := repository.Limit(limit)
	queryOffset := repository.Offset(offset)
	preload := []string{"Hobbies", "Passport"}
	pre := repository.PreloadAssociations(preload)
	qps := []repository.QueryProcessor{queryLimit, queryOffset, pre}
	result := us.Repo.GetAllWithQueryProcessor(uow, users, qps)
	if result != nil {
		us.Logger.Error().Msgf("Error while get users with pagination %v", result)
		return nil
	}
	var u []model.User
	fmt.Println(len(hobby))
	if len(hobby) == 1 && hobby[0] != "" {
		for _, user := range *users {
			for _, hob := range user.Hobbies {
				for _, val := range hobby {
					if hob.HobbyName == val {
						u = append(u, user)
					}
				}
			}
		}
		fmt.Println("u val ", len(u))
		return u
	}
	return *users
}

func (us *UserService) AddUserHobbies(user *model.User) error {
	uow := repository.NewUnitOfWork(us.DB, false)
	count := us.GetUsersCount("id = ?", user.ID)
	if count == 0 {
		us.Logger.Error().Msg("Could not find the user ")
		err := fmt.Errorf("could not find the user")
		return err
	}
	er := us.Repo.Update(uow, user)
	if er != nil {
		uow.Complete()
		us.Logger.Error().Msgf("Error adding user hobbies %v", er)
		return er
	}
	uow.Commit()

	return nil
}

func (us *UserService) DeleteUserHobbies(user *model.User) error {
	uow := repository.NewUnitOfWork(us.DB, false)
	count := us.GetUsersCount("id = ?", user.ID)
	if count == 0 {
		us.Logger.Error().Msg("Could not find the user ")
		err := fmt.Errorf("could not find the user")
		return err
	}
	var hobbyToBeDeleted model.Hobby
	for _, val := range user.Hobbies {
		qp := repository.Filter("hobby_name = ? AND user_id = ? ", val.HobbyName, user.ID)
		err := us.Repo.GetFirst(uow, &hobbyToBeDeleted, qp)
		//fmt.Println("Hobby ID in delete user hobby ", hobbyToBeDeleted.ID)
		if err != nil {
			us.Logger.Error().Msg("Error while getting user hobby for deleting")
		} else {
			e := us.Repo.Delete(uow, hobbyToBeDeleted)
			if e != nil {
				uow.Complete()
				us.Logger.Error().Msg("Error while deleting user hobby ")
			}
		}
		uow.Commit()
	}

	return nil
}
