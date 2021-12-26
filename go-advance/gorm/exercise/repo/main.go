package main

import (
	s "app/exerciseService"
	m "app/model"
	r "app/repository"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	dbConn := "root:hello@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dbConn)
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}
	re := r.NewRepository()
	employeeService := s.NewEmployeeService(re, db)
	regionService := s.NewRegionsService(re, db)
	employeeService.GetNumberOfEmployees()
	employeeService.GetSumOfSlaryofEmployees()
	employeeService.GetAverageSalaryofEmployees()
	var ep m.Employee
	var dept m.Department
	var reg m.Region
	var c m.Country
	employeeService.GetDeptWiseCount(&ep, "count(*) as DeptWiseCount,deptNo as DeptNo", "deptNo")
	employeeService.GetJobWiseCount(&ep, "count(*) as JobWiseCount,Job as Job", "JOB")
	employeeService.GetDeptAndJobWiseCount(&ep, "count(*) ,Job as Job,deptNo as DeptNo", "deptNo,JOB")
	employeeService.GetDeptWiseCountWithSpecificCondition(&ep, "count(*) as DeptWiseCount,deptNo as DeptNo", "deptNo")
	employeeService.GetEmployeeDeptNames(&ep, "emp.ename,dept.dname", "join dept on emp.deptno = dept.deptno")
	employeeService.GetDeptNameWiseCount(&ep, "count(*),dname", "dname", "join dept on emp.deptno=dept.deptno")
	employeeService.GetDeptNameJobWiseCount(&ep, "count(*),dname,job", "dname,job", "join dept on emp.deptno=dept.deptno")
	employeeService.GetEmpNameDeptNameWithNoEmployeeAsNull(&dept, "ename as ENAME,dname as DNAME,job as JOB", "left join emp on emp.deptno=dept.deptno")
	employeeService.GetDeptNameWhereEmployeesAreNull(&dept, "dname as DeptName", "left join emp on emp.deptno=dept.deptno")
	employeeService.GetEmployeeAndBossName(&ep, "e.ename,m.ename", "e join emp m on e.mgr=m.empno")
	//If no boss give null -> //SELECT  e.ename 'Employee', m.ename 'Manager' FROM emp e left join emp m ON (e.mgr = m.empno);
	employeeService.GetEmployeeAndBossName(&ep, "e.ename,m.ename", "e left join emp m on e.mgr=m.empno")
	//SELECT  e.ename 'Employee', m.ename 'Manager',dname FROM emp e left join emp m ON (e.mgr = m.empno) join dept d ON e.deptno=d.deptno ;
	employeeService.GetEmpBossAndDeptName(&ep, "e.ename,m.ename,dname", "e left join emp m on e.mgr=m.empno", " join dept d on e.deptno=d.deptno")
	regionService.GetRegionsWithNoCountries(&reg, "region_name", "left join countries on regions.region_id=countries.REGION_ID")
	//select country_name from countries left join locations on countries.COUNTRY_ID=locations.COUNTRY_ID group by countries.country_id having count(state_province)=0
	regionService.GetCountriesWithNoStates(&c, "`countries`.*", "left join locations on countries.COUNTRY_ID = locations.COUNTRY_ID", "countries.country_id")
	//select regions.region_name,countries.country_name,locations.STATE_PROVINCE from regions join countries ON regions.region_id = countries.REGION_ID left join locations on countries.country_id = locations.country_id;
	regionService.GetRegionCountryState(&reg, "regions.region_name,countries.country_name,locations.STATE_PROVINCE", "left join countries ON regions.region_id = countries.REGION_ID", "left join locations on countries.country_id = locations.country_id")

	//newEntry := m.NewLocation(1555, "IN", "Swabhav techlabs", "Mumbai", "Maharashtra", "431202")
	//regionService.CreateNewEntry(newEntry)
	regionService.GetEnitiesWithCity("Mumbai")
}
