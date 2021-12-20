package service

import (
	h "app/hash"
	lr "app/logger"
	m "app/model"
	re "app/repository"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type UserService struct {
	//uow *re.UnitOfWork
	Repo re.Repository
	DB   *gorm.DB
}

func NewUserService(r re.Repository, DB *gorm.DB) *UserService {
	return &UserService{Repo: r, DB: DB}
}
func (us *UserService) AddUser(u *m.User) {
	uow := re.NewUnitOfWork(us.DB, false)
	hashedPassword := h.CreateHashForPassword(u.Password)
	u.Password = hashedPassword
	err := us.Repo.Add(uow, u)
	if err != nil {
		logger := lr.GetLogger()
		logger.Error().Msgf("Error in add user %v", err)
	} else {
		uow.Commit()
	}
}
func (us *UserService) GetUser() {
	uow := re.NewUnitOfWork(us.DB, true)
	qp := re.Filter("name = ?", "Jay")
	//qps := []re.QueryProcessor{}
	//qps = append(qps, qp)
	preloadAssoc := []string{"Hobbies", "Courses"}
	pqp := re.PreloadAssociations(preloadAssoc)
	var user m.User
	err := us.Repo.GetFirst(uow, &user, qp, pqp)
	if err != nil {
		fmt.Println("Error using quey processor")
	}
	fmt.Println("User object from db --- ", user)
	fmt.Println("User courses ")
	for _, val := range user.Courses {
		fmt.Println("---", val.Name)
	}

}
func (us *UserService) GetUsers(out interface{}, preloadAssociations []string) []m.User {
	uow := re.NewUnitOfWork(us.DB, true)
	err := us.Repo.GetAll(uow, out, preloadAssociations)
	if err != nil {
		//fmt.Println("Error in get all user ", err)
		logger.Error().Msgf("Error in get all user %v ", err)
	}
	o := out.(*[]m.User)
	for _, val := range *o {
		fmt.Println(val.FirstName)
		for _, val := range val.Hobbies {
			fmt.Println("Hobby ---", val.HobbyName)
		}
		for _, vl := range val.Courses {
			fmt.Println("Courses ---", vl.Name)
		}
	}
	return *o
}

func (us *UserService) GetUserById(out interface{}, tenantID uuid.UUID, preloadAssociations []string) m.User {
	uow := re.NewUnitOfWork(us.DB, true)
	err := us.Repo.GetAllForTenant(uow, out, tenantID, preloadAssociations)
	if err != nil {
		//fmt.Println("Error in get all user ", err)
		logger.Error().Msg("Error in get all user ")
	}
	o := out.(*m.User)
	return *o
}

func (us *UserService) GetUserHobbies(user *m.User) []m.Hobby {
	p := []string{"Hobbies"}
	userByID := us.GetUserById(user, user.ID, p)
	var hobbiesForUser []m.Hobby
	for _, val := range userByID.Hobbies {
		hobbiesForUser = append(hobbiesForUser, val)
	}
	return hobbiesForUser
}

func (us *UserService) UpdateUser(entity interface{}) error {
	uow := re.NewUnitOfWork(us.DB, false)
	u := entity.(*m.User)
	hashedPassword := h.CreateHashForPassword(u.Password)
	u.Password = hashedPassword
	err := us.Repo.Update(uow, u)
	if err != nil {
		uow.Complete()
		fmt.Println("Error updating user")
		return err
	} else {
		uow.Commit()
	}
	user := entity.(*m.User)
	fmt.Println("Passport ID ", user.Passport.ID)
	zeroUUID, _ := uuid.FromString("00000000-0000-0000-0000-000000000000")
	if user.Passport.ID == zeroUUID {
		us.FindAndDeletePassport(user, []string{"Passport"})
	}
	return nil
}

func (us *UserService) DeleteUser(entity interface{}) {
	uow := re.NewUnitOfWork(us.DB, false)
	user := entity.(*m.User)
	err := us.Repo.GetFirst(uow, &user)
	if err != nil {
		logger.Error().Msgf("Could not find the user %v", err)
	} else {
		er := us.Repo.Delete(uow, entity)
		if er != nil {
			uow.Complete()
			//fmt.Println("Error deleting user")
			logger.Error().Msgf("Error deleting user %v", er)
		} else {
			var h []m.Hobby
			e := uow.DB.Debug().Where("user_id = ?", user.ID).Find(&h).Error
			if len(h) != 0 {
				if e != nil {
					logger.Error().Msg("Error finding associated hobbies for user")
				} else {
					for _, val := range h {
						ef := us.Repo.Delete(uow, &val)
						if ef != nil {
							uow.Complete()
							logger.Error().Msgf("Error deleting hobby for this user %v", ef)
						}
					}
				}
			}
			we := uow.DB.Model(&user).Debug().Association("courses").Clear().Error
			if we != nil {
				uow.Complete()
				logger.Error().Msgf("Error while deleting associated courses %v ", we)
			}
			us.FindAndDeletePassport(user, []string{"Passport"})
			uow.Commit()
		}
	}
}

func (us *UserService) GetUsersWithCourse(id uuid.UUID, preloadAssociations []string) {
	uow := re.NewUnitOfWork(us.DB, true)
	var people []m.User
	a := []string{}
	var c m.Course
	us.Repo.Get(uow, &c, id, a)
	fmt.Println("Course info ", c)
	f := uow.DB.Model(&c).Debug().Related(&people, "Users").Error
	if f != nil {
		fmt.Println("Error :", f)
	}
	if len(people) == 0 {
		fmt.Println("No user with this course ID")
	} else {
		for i, _ := range people {
			fmt.Println("Entities with given course ID ", people[i].FirstName)
		}
	}

}

func (us *UserService) GetAllUsersWithPagination(page, limit int, hobby []string, out interface{}) []m.User {
	uow := re.NewUnitOfWork(us.DB, true)
	offset := (page - 1) * limit
	queryLimit := re.Limit(limit)
	queryOffset := re.Offset(offset)
	preload := []string{"Hobbies", "Passport"}
	pqp := re.PreloadAssociations(preload)
	qp := []re.QueryProcessor{queryLimit, queryOffset, pqp}
	result := us.Repo.GetAllWithQueryProcessor(uow, out, qp)
	if result != nil {
		log.Fatal("Error in pagination ")
		return nil
	}
	o := out.(*[]m.User)
	var u []m.User
	fmt.Println(len(hobby))
	if len(hobby) == 1 && hobby[0] != "" {
		for _, user := range *o {
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
	return *o
}

func (us *UserService) AddUserHobbies(user *m.User) error {
	uow := re.NewUnitOfWork(us.DB, false)
	err := us.Repo.Update(uow, user)
	if err != nil {
		uow.Complete()
		logger.Error().Msgf("Error adding user hobbies %v", err)
		return err
	} else {
		uow.Commit()
	}
	return nil
}

func (us *UserService) DeleteUserHobbies(user *m.User) error {
	uow := re.NewUnitOfWork(us.DB, false)
	//fmt.Println("User hobbies to be deleted ", user.Hobbies)
	var hobbyToBeDeleted m.Hobby
	for _, val := range user.Hobbies {
		qp := re.Filter("hobby_name = ? AND user_id = ? ", val.HobbyName, user.ID)
		err := us.Repo.GetFirst(uow, &hobbyToBeDeleted, qp)
		//fmt.Println("Hobby ID in delete user hobby ", hobbyToBeDeleted.ID)
		if err != nil {
			logger.Error().Msg("Error while getting user hobby for deleting")
		} else {
			e := us.Repo.Delete(uow, hobbyToBeDeleted)
			if e != nil {
				uow.Complete()
				logger.Error().Msg("Error while deleting user hobby ")
			}
		}
		uow.Commit()
	}
	return nil
}
