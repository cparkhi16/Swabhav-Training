package service

import (
	m "app/model"
	re "app/repository"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type UserService struct {
	uow *re.UnitOfWork
}

func NewUserService(uow *re.UnitOfWork) *UserService {
	return &UserService{uow: uow}
}
func (us *UserService) AddUser(u *m.User) {
	r := re.NewRepository()
	err := r.Add(us.uow, u)
	if err != nil {
		fmt.Println("Error in add user ", err)
	} else {
		us.uow.Commit()
	}
}
func (us *UserService) GetUser() {
	r := re.NewRepository()
	qp := re.Filter("name = ?", "Jay")
	//qps := []re.QueryProcessor{}
	//qps = append(qps, qp)
	var user m.User
	err := r.GetFirst(us.uow, &user, qp)
	if err != nil {
		fmt.Println("Error using quey processor")
	}
	fmt.Println("User object from db --- ", user)
}
func (us *UserService) GetUsers(out interface{}, preloadAssociations []string) {
	r := re.NewRepository()
	err := r.GetAll(us.uow, out, preloadAssociations)
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
}

func (us *UserService) GetUserById(out interface{}, tenantID uuid.UUID, preloadAssociations []string) *m.User {
	r := re.NewRepository()
	err := r.GetAllForTenant(us.uow, out, tenantID, preloadAssociations)
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

func (us *UserService) UpdateUser(entity interface{}) {
	fmt.Println(entity)
	r := re.NewRepository()
	err := r.Update(us.uow, entity)
	if err != nil {
		us.uow.Complete()
		fmt.Println("Error updating user")
	} else {
		us.uow.Commit()
	}
}

func (us *UserService) DeleteUser(entity interface{}) {
	r := re.NewRepository()
	user := entity.(*m.User)
	err := r.Delete(us.uow, entity)
	if err != nil {
		us.uow.Complete()
		fmt.Println("Error deleting user")
	} else {
		defer us.uow.Commit()
		fmt.Println("Deleting hobby entry for this user --")
		var h []m.Hobby
		fmt.Println("User ID ---", user.ID)
		us.uow.DB.Where("user_id = ?", user.ID).Find(&h)
		fmt.Println("Hobby for user", h)
		//ef := uow.DB.Debug().Delete(&h).Error
		for _, val := range h {
			ef := r.Delete(us.uow, &val)
			if ef != nil {
				fmt.Println("Error deleting hobby for this user")
			}
		}
		fmt.Println("Deleting course user entry from person_courses ")
		//var courses []m.Course
		//uow.DB.Where("user_id = ?", user.ID).Find(&courses)
		we := us.uow.DB.Model(&user).Debug().Association("courses").Clear().Error
		if we != nil {
			fmt.Println("Error deleting user course map in person_courses")
		}

	}
}

func (us *UserService) GetUsersWithCourse(id uuid.UUID, preloadAssociations []string) {
	r := re.NewRepository()
	var people []m.User
	a := []string{}
	var c m.Course
	r.Get(us.uow, &c, id, a)
	fmt.Println("Course info ", c)
	f := us.uow.DB.Model(&c).Debug().Related(&people, "Users").Error
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
