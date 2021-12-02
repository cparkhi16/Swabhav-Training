package main

import (
	"fmt"
)

type DBConn interface {
	Query() interface{}
}
type UsersRepository struct {
	db DBConn
	//db MySQL //Violates DIP because UsersRepo is directly dependent on concrete implementation rather than interface
}
type MySQL struct{}

/*func (db MySQL) QueryMySQLDB() map[string]string {
	return map[string]string{
		"abcde-1236": "raj",
		"abcdf-1234": "rahul",
	}
}*/
func (db MySQL) Query() interface{} {
	return map[string]string{
		"abcde-1236": "raj",
		"abcdf-1234": "rahul",
	}
}

type PostgreSQL struct{}

/*func (db MySQL) QueryPostgreSQL() []string {
	return []string{"raj", "mahesh", "ram"}
}*/
func (db PostgreSQL) Query() interface{} {
	return []string{"raj", "mahesh", "ram"}
}
func main() {
	/*db := MySQL{}
	users := UsersRepository{db: db}
	fmt.Println(users.db.QueryMySQLDB())*/

	dbPSQL := PostgreSQL{}
	user := UsersRepository{db: dbPSQL}
	fmt.Println(user.db.Query())
}
