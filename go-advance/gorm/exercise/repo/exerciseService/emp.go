package exerciseservice

import (
	r "app/repository"
	"fmt"

	m "app/model"

	"github.com/jinzhu/gorm"
)

type EmployeeService struct {
	//uow *r.UnitOfWork
	Repo r.Repository
	DB   *gorm.DB
}

func NewEmployeeService(r r.Repository, DB *gorm.DB) *EmployeeService {
	return &EmployeeService{Repo: r, DB: DB}
}
func (e *EmployeeService) GetNumberOfEmployees() {
	uow := r.NewUnitOfWork(e.DB, true)
	var emp []m.Employee
	err := e.Repo.GetAllWithQueryProcessor(uow, &emp)
	if err != nil {
		fmt.Println("Error using query processor", err)
	}
	fmt.Println("Count of emp ", len(emp))
}
func (e *EmployeeService) GetSumOfSlaryofEmployees() {
	uow := r.NewUnitOfWork(e.DB, true)
	var emp []m.Employee
	err := e.Repo.GetAllWithQueryProcessor(uow, &emp)
	if err != nil {
		fmt.Println("Error using query processor", err)
	}
	//fmt.Println("Count of emp ", len(emp))
	var sal uint
	for _, val := range emp {
		sal = sal + val.Salary
	}
	fmt.Println("Sum of salary of all employees ", sal)
}
func (e *EmployeeService) GetAverageSalaryofEmployees() {
	uow := r.NewUnitOfWork(e.DB, true)
	var emp []m.Employee
	err := e.Repo.GetAllWithQueryProcessor(uow, &emp)
	var count uint = uint(len(emp))
	if err != nil {
		fmt.Println("Error using query processor", err)
	}
	//fmt.Println("Count of emp ", len(emp))
	var sal uint
	for _, val := range emp {
		sal = sal + val.Salary
	}
	fmt.Println("Average salary of employees ", sal/count)
}

func (e *EmployeeService) GetDeptWiseCount(entity interface{}, statement, group string) {
	//SELECT count(*) as DeptWiseCount,deptNo as DeptNo FROM `emp`   GROUP BY deptNo
	uow := r.NewUnitOfWork(e.DB, true)
	//qr := r.GroupBy(entity, statement, group)
	selectStatement := r.Select(statement)
	entityModel := r.Model()
	qr := r.Group(group)
	//var emp []m.Employee
	//db, _ := qr(e.uow.DB, &emp)
	rows, _ := e.Repo.GetAllRows(uow, entity, entityModel, selectStatement, qr)
	var result []m.Employee
	for rows.Next() {
		emp := m.Employee{}
		rows.Scan(&emp.DeptWiseCount, &emp.DeptNo)
		result = append(result, emp)
	}
	for _, val := range result {
		fmt.Println("DeptWiseCount - ", val.DeptWiseCount)
		fmt.Println("DeptNo -", val.DeptNo)
		fmt.Println("=============")
	}
}

func (e *EmployeeService) GetJobWiseCount(entity interface{}, statement, group string) {
	//SELECT count(*) as JobWiseCount,Job as Job FROM `emp`   GROUP BY JOB
	uow := r.NewUnitOfWork(e.DB, true)
	//qr := r.GroupBy(entity, statement, group)
	selectStatement := r.Select(statement)
	entityModel := r.Model()
	qr := r.Group(group)
	//var emp []m.Employee
	//db, _ := qr(e.uow.DB, &emp)
	rows, _ := e.Repo.GetAllRows(uow, entity, entityModel, selectStatement, qr)
	var result []m.Employee
	for rows.Next() {
		emp := m.Employee{}
		rows.Scan(&emp.JobWiseCount, &emp.Job)
		result = append(result, emp)
	}
	for _, val := range result {
		fmt.Println("JobWiseCount - ", val.JobWiseCount)
		fmt.Println("Job -", val.Job)
		fmt.Println("=============")
	}
}

func (e *EmployeeService) GetDeptAndJobWiseCount(entity interface{}, statement, group string) {
	//SELECT count(*) ,Job as Job,deptNo as DeptNo FROM `emp`   GROUP BY deptNo,JOB
	uow := r.NewUnitOfWork(e.DB, true)
	//qr := r.GroupBy(entity, statement, group)
	//var emp []m.Employee
	//db, _ := qr(e.uow.DB, &emp)
	selectStatement := r.Select(statement)
	entityModel := r.Model()
	qr := r.Group(group)
	rows, _ := e.Repo.GetAllRows(uow, entity, entityModel, selectStatement, qr)
	var result []m.Employee
	for rows.Next() {
		emp := m.Employee{}
		rows.Scan(&emp.JobWiseCount, &emp.Job, &emp.DeptNo)
		result = append(result, emp)
	}
	for _, val := range result {
		fmt.Println("Count - ", val.JobWiseCount)
		fmt.Println("Job -", val.Job)
		fmt.Println("DeptNo -", val.DeptNo)
		fmt.Println("=============")
	}
}

func (e *EmployeeService) GetDeptWiseCountWithSpecificCondition(entity interface{}, statement, group string) {
	// SELECT count(*) as DeptWiseCount,deptNo as DeptNo FROM `emp`  WHERE (deptNo = 10 or deptNo = 20) GROUP BY deptNo HAVING (count(*)>= 2) ORDER BY count(*) desc
	uow := r.NewUnitOfWork(e.DB, true)
	whereClause := r.Filter("deptNo = ? or deptNo = ?", 10, 20)
	//qr := r.GroupBy(entity, statement, group)
	selectStatement := r.Select(statement)
	entityModel := r.Model()
	qr := r.Group(group)
	havingClause := r.Having("count(*)>= ?", 2)
	orderBy := r.OrderBy("count(*) desc")
	//var emp []m.Employee
	//err := e.Repo.GetAllWithQueryProcessor(uow, &emp, whereClause, qr, havingClause, orderBy)
	rows, _ := e.Repo.GetAllRows(uow, entity, whereClause, entityModel, selectStatement, qr, havingClause, orderBy)
	var result []m.Employee
	for rows.Next() {
		emp := m.Employee{}
		rows.Scan(&emp.DeptWiseCount, &emp.DeptNo)
		result = append(result, emp)
	}
	for _, val := range result {
		fmt.Println("Count --", val.DeptWiseCount)
		fmt.Println("DeptNo-", val.DeptNo)
	}

}

func (e *EmployeeService) GetEmployeeDeptNames(entity interface{}, statement, condition string) {
	//SELECT emp.ename,dept.dname FROM `emp` join dept on emp.deptno = dept.deptno
	uow := r.NewUnitOfWork(e.DB, true)
	//var emp []m.Employee
	//rows, _ := e.DB.Table("emp").Select("emp.ename, dept.dname").Joins("join dept on emp.deptno = dept.deptno").Rows()
	//join := r.Join(entity, statement, condition)
	selectStatement := r.Select(statement)
	entityModel := r.Model()
	joins := r.Joins(condition)
	rows, err := e.Repo.GetAllRows(uow, entity, entityModel, selectStatement, joins)
	if err != nil {
		fmt.Println("Error in join ", err)
		return
	}
	result := make(map[*m.Employee]*m.Department)
	for rows.Next() {
		emp := m.Employee{}
		dep := m.Department{}
		rows.Scan(&emp.EmpName, &dep.DeptName)
		result[&emp] = &dep
	}
	fmt.Println("--- Employee and dept join ----")
	for k, val := range result {
		fmt.Println("Emp Name-", k.EmpName)
		fmt.Println("Dept Name - ", val.DeptName)
		fmt.Println("=========")
	}
}

func (e *EmployeeService) GetDeptNameWiseCount(entity interface{}, statement, group, condition string) {
	uow := r.NewUnitOfWork(e.DB, true)
	//SELECT count(*),dname from emp join dept on emp.DEPTNO = dept.DEPTNO group by dname;
	//join := r.Join(entity, statement, condition)
	//qr := r.GroupBy(entity, statement, group)
	selectStatement := r.Select(statement)
	entityModel := r.Model()
	joins := r.Joins(condition)
	qr := r.Group(group)
	//var emp []m.Employee
	result := make(map[*m.Employee]*m.Department)
	rows, err := e.Repo.GetAllRows(uow, entity, entityModel, selectStatement, joins, qr)
	if err != nil {
		fmt.Println("Error in join ", err)
		return
	}
	for rows.Next() {
		emp := m.Employee{}
		dep := m.Department{}
		rows.Scan(&emp.DeptWiseCount, &dep.DeptName)
		result[&emp] = &dep
	}
	for k, val := range result {
		fmt.Println("DeptWiseCount - ", k.DeptWiseCount)
		fmt.Println("DeptName -", val.DeptName)
		fmt.Println("=============")
	}
}

func (e *EmployeeService) GetDeptNameJobWiseCount(entity interface{}, statement, group, condition string) {
	uow := r.NewUnitOfWork(e.DB, true)
	//SELECT count(*),dname,job from emp join dept on emp.DEPTNO = dept.DEPTNO group by dname,job;
	//join := r.Join(entity, statement, condition)
	selectStatement := r.Select(statement)
	entityModel := r.Model()
	joins := r.Joins(condition)
	//qr := r.GroupBy(entity, statement, group)
	qp := r.Group(group)
	//var emp []m.Employee
	result := make(map[*m.Employee]*m.Department)
	rows, err := e.Repo.GetAllRows(uow, entity, entityModel, selectStatement, joins, qp)
	if err != nil {
		fmt.Println("Error in join ", err)
		return
	}
	for rows.Next() {
		emp := m.Employee{}
		dep := m.Department{}
		rows.Scan(&emp.DeptWiseCount, &dep.DeptName, &emp.Job)
		result[&emp] = &dep
	}
	for k, val := range result {
		fmt.Println("Job name ", k.Job)
		fmt.Println("DeptWiseCount - ", k.DeptWiseCount)
		fmt.Println("DeptName -", val.DeptName)

		fmt.Println("=============")
	}
}

func (e *EmployeeService) GetEmpNameDeptNameWithNoEmployeeAsNull(entity interface{}, statement, condition string) {
	uow := r.NewUnitOfWork(e.DB, true)
	//SELECT ename,dname,job from dept left join emp on emp.DEPTNO = dept.DEPTNO ;
	//join := r.Join(entity, statement, condition)
	selectStatement := r.Select(statement)
	entityModel := r.Model()
	joins := r.Joins(condition)
	//var dep []m.Department
	result := make(map[*m.Department]*m.Employee)
	//where := r.Filter("emp.deptno = ?", "IS Null")
	rows, err := e.Repo.GetAllRows(uow, entity, entityModel, selectStatement, joins)
	if err != nil {
		fmt.Println("Error in join ", err)
		return
	}
	//rows,_=e.DB.Table("dept").Select("ename as ENAME,dname as DNAME,job as JOB").Joins("left join emp on emp.DEPTNO = dept.DEPTNO").Rows()
	for rows.Next() {
		emp := m.Employee{}
		dep := m.Department{}
		rows.Scan(&emp.EmpName, &dep.DeptName, &emp.Job)
		//fmt.Println("dept name ", dep.DeptName)
		result[&dep] = &emp
	}
	for k, val := range result {
		//fmt.Println(k, val)
		fmt.Println("Dept name ", k.DeptName) // Dept Operations coming as blank value
		fmt.Println("Emp Name - ", val.EmpName)
		fmt.Println("JOB -", val.Job)
		fmt.Println("=============")
	}
}
func (e *EmployeeService) GetDeptNameWhereEmployeesAreNull(entity interface{}, statement, condition string) {
	uow := r.NewUnitOfWork(e.DB, true)
	//SELECT dname FROM dept left join emp on emp.deptno=dept.deptno where emp.DEPTNO is NUll;
	//join := r.Join(entity, statement, condition)
	selectStatement := r.Select(statement)
	entityModel := r.Model()
	joins := r.Joins(condition)
	//var emp []m.Employee
	result := make(map[*m.Department]*m.Employee)
	where := r.Filter("emp.deptno is null")
	rows, err := e.Repo.GetAllRows(uow, entity, entityModel, selectStatement, joins, where)
	if err != nil {
		fmt.Println("Error in join ", err)
		return
	}
	for rows.Next() {
		emp := m.Employee{}
		dep := m.Department{}
		rows.Scan(&dep.DeptName)
		//fmt.Println("dept name ", dep.DeptName)
		result[&dep] = &emp
	}
	for k, _ := range result {
		fmt.Println("Dept name ", k.DeptName)
	}
}

func (e *EmployeeService) GetEmployeeAndBossName(entity interface{}, statement, condition string) {
	// SELECT e.empno 'Emp_Id', e.ename 'Employee', m.empno 'Mgr_Id', m.ename 'Manager' FROM emp e join emp m ON (e.mgr = m.empno);
	// SELECT  e.ename 'Employee', m.ename 'Manager' FROM emp e join emp m ON (e.mgr = m.empno);
	uow := r.NewUnitOfWork(e.DB, true)
	//join := r.Join(entity, statement, condition)
	selectStatement := r.Select(statement)
	entityModel := r.Model()
	joins := r.Joins(condition)
	//var emp []m.Employee
	rows, err := e.Repo.GetAllRows(uow, entity, entityModel, selectStatement, joins)
	if err != nil {
		fmt.Println("Error in join ", err)
		return
	}
	var result []m.Employee
	for rows.Next() {
		emp := m.Employee{}
		rows.Scan(&emp.EmpName, &emp.ManagerName)
		result = append(result, emp)
	}
	for _, val := range result {
		fmt.Println("Employee - ", val.EmpName)
		fmt.Println("Manager -", val.ManagerName)
		fmt.Println("=============")
	}

}
func (e *EmployeeService) GetEmpBossAndDeptName(entity interface{}, statement string, condition ...string) {
	//SELECT e.ename,m.ename,dname FROM `emp` e left join emp m on e.mgr=m.empno  join dept d on e.deptno=d.deptno
	uow := r.NewUnitOfWork(e.DB, true)
	//var emp []m.Employee
	//join := r.Join(entity, statement, condition[0])
	//joinTwo := r.Join(entity, statement, condition[1])
	selectStatement := r.Select(statement)
	entityModel := r.Model()
	join := r.Joins(condition[0])
	joinTwo := r.Joins(condition[1])
	rows, err := e.Repo.GetAllRows(uow, entity, entityModel, selectStatement, join, joinTwo)
	if err != nil {
		fmt.Println("Error in join ", err)
		return
	}
	result := make(map[*m.Employee]*m.Department)
	for rows.Next() {
		emp := m.Employee{}
		dep := m.Department{}
		rows.Scan(&emp.EmpName, &emp.ManagerName, &dep.DeptName)
		result[&emp] = &dep
	}
	fmt.Println("--- Employee Manager and dept join ----")
	for k, val := range result {
		fmt.Println("Emp Name-", k.EmpName)
		fmt.Println("Manager name -", k.ManagerName) //Getting empty dept name for king
		fmt.Println("Dept Name - ", val.DeptName)
		fmt.Println("=========")
	}
}
