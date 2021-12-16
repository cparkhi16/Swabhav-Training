package service

import (
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
	//r := re.NewRepository()
	uow := re.NewUnitOfWork(us.DB, false)
	fmt.Println("User ID after before create ", u.ID)
	err := us.Repo.Add(uow, u)
	if err != nil {
		fmt.Println("Error in add user ", err)
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

	fmt.Println("--= Combined filter and preload --=")
	var JayUser m.User
	cqp := re.FilterAndPreloadAssociations("name = ?", preloadAssoc, "Jay")
	errt := us.Repo.GetFirst(uow, &JayUser, cqp)
	if errt != nil {
		fmt.Println("Error for combined query")
	}
	fmt.Println("User object from db for combined query --- ", JayUser)
	fmt.Println("User courses for combined query ")
	for _, val := range JayUser.Courses {
		fmt.Println("---", val.Name)
	}
}
func (us *UserService) GetUsers(out interface{}, preloadAssociations []string) []m.User {
	uow := re.NewUnitOfWork(us.DB, true)
	err := us.Repo.GetAll(uow, out, preloadAssociations)
	if err != nil {
		fmt.Println("Error in get all user ", err)
	}
	o := out.(*[]m.User)
	for _, val := range *o {
		fmt.Println(val.Name)
		for _, val := range val.Hobbies {
			fmt.Println("Hobby ---", val.HobbyName)
		}
		for _, vl := range val.Courses {
			fmt.Println("Courses ---", vl.Name)
		}
	}
	return *o
}

func (us *UserService) GetUserById(out interface{}, tenantID uuid.UUID, preloadAssociations []string) *m.User {
	uow := re.NewUnitOfWork(us.DB, true)
	err := us.Repo.GetAllForTenant(uow, out, tenantID, preloadAssociations)
	if err != nil {
		fmt.Println("Error in get all user ", err)
	}
	o := out.(*m.User)

	fmt.Println(o.Name)
	if len(o.Hobbies) > 1 {
		fmt.Println(o.Hobbies[0].HobbyName)
	}
	return o
}

func (us *UserService) UpdateUser(entity interface{}) error {
	uow := re.NewUnitOfWork(us.DB, false)
	err := us.Repo.Update(uow, entity)
	if err != nil {
		uow.Complete()
		fmt.Println("Error updating user")
		return err
	} else {
		uow.Commit()
	}
	return nil
}

/*func (us *UserService) DeletePassportDetails(entity interface{}) error {
	r := re.NewRepository()
	err := r.Delete(us.uow, entity)
	if err != nil {
		return err
	}
	return nil
}*/
func (us *UserService) DeleteUser(entity interface{}) {
	uow := re.NewUnitOfWork(us.DB, false)
	user := entity.(*m.User)
	err := us.Repo.Delete(uow, entity)
	if err != nil {
		uow.Complete()
		fmt.Println("Error deleting user")
	} else {
		defer uow.Commit()
		fmt.Println("Deleting hobby entry for this user --")
		var h []m.Hobby
		fmt.Println("User ID ---", user.ID)
		uow.DB.Where("user_id = ?", user.ID).Find(&h)
		fmt.Println("Hobby for user", h)
		//ef := uow.DB.Debug().Delete(&h).Error
		for _, val := range h {
			ef := us.Repo.Delete(uow, &val)
			if ef != nil {
				fmt.Println("Error deleting hobby for this user")
			}
		}
		fmt.Println("Deleting course user entry from person_courses ")
		//var courses []m.Course
		//uow.DB.Where("user_id = ?", user.ID).Find(&courses)
		we := uow.DB.Model(&user).Debug().Association("courses").Clear().Error
		if we != nil {
			fmt.Println("Error deleting user course map in person_courses")
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
			fmt.Println("Entities with given course ID ", people[i].Name)
		}
	}

}

func (us *UserService) GetPassportIDForUser(ID uuid.UUID) m.Passport {
	uow := re.NewUnitOfWork(us.DB, true)
	qp := re.Filter("user_id = ?", ID)
	var p m.Passport
	err := us.Repo.GetFirst(uow, &p, qp)
	if err != nil {
		log.Fatal("Error in get passport id ", err)
	}
	return p
}

func (us *UserService) GetAllUsersWithPagination(page, limit int, hobby []string, out interface{}) []m.User {
	uow := re.NewUnitOfWork(us.DB, true)
	offset := (page - 1) * limit
	//fmt.Println("h", hobby)
	//queryBuider := us.uow.DB.Debug().Limit(limit).Offset(offset)
	queryLimit := re.Limit(limit)
	queryOffset := re.Offset(offset)
	preload := []string{"Hobbies", "Passport"}
	pqp := re.PreloadAssociations(preload)
	//result := queryBuider.Model(out).Find(out)
	qp := []re.QueryProcessor{queryLimit, queryOffset, pqp}
	result := us.Repo.GetAllUsers(uow, out, qp)
	//res := result.Debug().Preload("Hobbies").Find(out)
	if result != nil {
		log.Fatal("Error in pagination ")
		return nil
	}
	o := out.(*[]m.User)
	var u []m.User
	fmt.Println(len(hobby))
	if len(hobby)-1 != 0 {
		for _, user := range *o {
			for _, hob := range user.Hobbies {
				//fmt.Println(hob.HobbyName)
				/*if hob.HobbyName == hobby {
					u = append(u, user)
				}*/
				for _, val := range hobby {
					//fmt.Println("Val in hobbies slice", val)
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

func (us *UserService) UpdatePassportDetailForUser(entity interface{}) error {
	uow := re.NewUnitOfWork(us.DB, false)
	err := us.Repo.Update(uow, entity)
	if err != nil {
		uow.Complete()
		fmt.Println("Error updating user's passport")
		return err
	} else {
		uow.Commit()
	}
	return nil
}

func (us *UserService) FindAndDeletePassport(entity interface{}, preloadAssociations []string) error {
	uow := re.NewUnitOfWork(us.DB, false)
	pqp := re.PreloadAssociations(preloadAssociations)
	err := us.Repo.GetFirst(uow, entity, pqp)
	if err != nil {
		fmt.Println("error finding user", err)
		return err
	} else {
		o := entity.(*m.User)
		var p m.Passport
		p.ID = o.Passport.ID
		fmt.Println("Passport ID for user", p.ID)
		e := us.Repo.Delete(uow, &p)
		if e != nil {
			fmt.Println("Error deleting passport for user", e)
			return e
		}
	}
	return nil
}
