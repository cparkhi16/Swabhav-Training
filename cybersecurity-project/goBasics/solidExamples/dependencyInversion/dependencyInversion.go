package main

import "fmt"

type database interface {
	query()
}

type mariaDb struct {
	tableName string
}

func (m mariaDb) query() {
	fmt.Println("mariaDB query table-", m.tableName)
}

type sqlDb struct {
	tableName string
}

func (s sqlDb) query() {
	fmt.Println("sqlDB query table-", s.tableName)
}

func selectQueryOnDb(dbObj database) {
	dbObj.query()
}

func main() {
	db1 := mariaDb{tableName: "users"}
	selectQueryOnDb(db1)
	db2 := sqlDb{tableName: "data"}
	selectQueryOnDb(db2)
}
