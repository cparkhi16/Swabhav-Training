package service

import (
	m "app/model"
	re "app/repository"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func AddUser(uow *re.UnitOfWork, u *m.User) {
	r := re.NewRepository()
	err := r.Add(uow, u)
	if err != nil {
		fmt.Println("Error in add user ", err)
	} else {
		uow.Commit()
	}
}
func GetUser(uow *re.UnitOfWork) {
	r := re.NewRepository()
	qp := re.Filter("name = ?", "Jay")
	//qps := []re.QueryProcessor{}
	//qps = append(qps, qp)
	var user m.User
	err := r.GetFirst(uow, &user, qp)
	if err != nil {
		fmt.Println("Error using quey processor")
	}
	fmt.Println("User object from db --- ", user)
}
func GetUsers(uow *re.UnitOfWork, out interface{}, preloadAssociations []string) {
	r := re.NewRepository()
	err := r.GetAll(uow, out, preloadAssociations)
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

func GetUserById(uow *re.UnitOfWork, out interface{}, tenantID uuid.UUID, preloadAssociations []string) *m.User {
	r := re.NewRepository()
	err := r.GetAllForTenant(uow, out, tenantID, preloadAssociations)
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

func UpdateUser(uow *re.UnitOfWork, entity interface{}) {
	fmt.Println(entity)
	r := re.NewRepository()
	err := r.Update(uow, entity)
	if err != nil {
		fmt.Println("Error updating user")
	}
}

func DeleteUser(uow *re.UnitOfWork, entity interface{}) {
	r := re.NewRepository()
	user := entity.(*m.User)
	err := r.Delete(uow, entity)
	if err != nil {
		fmt.Println("Error deleting user")
	} else {
		fmt.Println("Deleting hobby entry for this user --")
		var h []m.Hobby
		fmt.Println("User ID ---", user.ID)
		uow.DB.Where("user_id = ?", user.ID).Find(&h)
		fmt.Println("Hobby for user", h)
		//ef := uow.DB.Debug().Delete(&h).Error
		for _, val := range h {
			ef := r.Delete(uow, &val)
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

func GetUsersWithCourse(uow *re.UnitOfWork, id uuid.UUID, preloadAssociations []string) {
	r := re.NewRepository()
	var people []m.User
	a := []string{}
	var c m.Course
	r.Get(uow, &c, id, a)
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
