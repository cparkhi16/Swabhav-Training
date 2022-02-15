package main

import "fmt"

//VIOLATION
/*As per dependency Inversion principle, entities must depend on abstractions not on concrete implementations
Here selectQueryOnDb function depends on mariaDb object, which is a concrete imlementation instead it should
depend on abstract interface so that we can switch database object anytime.
*/

type mariaDb struct {
	tableName string
}

func (m mariaDb) query() {
	fmt.Println("query table-", m.tableName)
}

func selectQueryOnDb(dbObj mariaDb) {
	dbObj.query()
}

func main() {
	db := mariaDb{tableName: "users"}
	selectQueryOnDb(db)
}
