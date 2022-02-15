package main

import (
	"fmt"

	"exercise/repository"
	"exercise/service"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, _ := ConnectToDB("root", "Panda@19", "localhost", "3306", "test2")
	fmt.Println(db)
	repo1 := repository.NewRepository()
	queryService := service.QueryService{
		Repo: repo1,
		Db:   db,
	}

	//var count int
	queryService.GetEmpCount() //1-display no of employees
	//fmt.Println(count)

	//var sum float64
	queryService.GetSalSum() //2-display sum of salaries of employees
	//fmt.Println(sum)

	//var avg float64
	queryService.GetAvgSal() //3-display avg salaries of employees
	//fmt.Println(avg)

	queryService.DisplaySumAvgCount()       //4-display sum,avg,count of employees
	queryService.DeptWiseHeadCount()        //5-display the dept wise , headcount
	queryService.JobWiseHeadCount()         //6-display the jobwise headcount
	queryService.DeptwiseJobwiseHeadcount() //7-display dept wise ,jobwise head count
	queryService.Test()
	queryService.DeptwiseEmployees()           //8-display the deptwise employees whose count greater than 2 and who are in dept 10 ,20.Sort the result by descending order of count
	queryService.DisplayEmpDept()              //9-display ename,deptname, 10and11 same as 5and6
	queryService.DisplayDeptEmp()              //12-display all the departments , with employees if any (if no emps then display null)
	queryService.DisplayAllDeptsWithNoEmp()    //13-display the departments where there are no employees
	queryService.DisplayEmpBoss()              //14- display the empname and their bossnames    //not showing emp with null boss
	queryService.DisplayEmpBossWithNulls()     //15-display all the empnames and boss names if any (if no boss display null)
	queryService.DisplayEmpBossDept()          //16-display ename,deptname and bossname . //not showing emp with null boss
	queryService.DisplayRegionsWithNoCountry() //17-display the regions there no entry for country
	queryService.DisplayCountriesWithNoState() //18-display the countries there no states
	queryService.DisplayRegionCountryState()   //19-display region name,country name and state name
	queryService.CreateLocation()              //20-Make an insert of swabhav Techlabs in location/state tables map to india and asia.
	queryService.GetMumbaiData()               //21-Filter details based on mumbai location
	queryService.InsertIntoFoo()               //22-create a foo table and insert values of different data
	//23-select data and after that create a primary key,check clustered index
}

func ConnectToDB(username string, password string, host string, port string, dbName string) (*gorm.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local", username, password, host, port, dbName)
	var err error
	var db *gorm.DB
	db, err = gorm.Open("mysql", dataSourceName)
	return db, err
}
