package service

import (
	"exercise/model"
	"exercise/repository"
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type QueryService struct {
	Repo repository.Repository
	Db   *gorm.DB
}

func (q *QueryService) GetEmpCount() {
	fmt.Println("1)display no of employees")
	//display no of employees
	//q.Db.Debug().Table("emp").Count(count)

	// unit := repository.NewUnitOfWork(q.Db, true)
	// var result model.Emp
	// //var queryp []repository.QueryProcessor
	// q.Repo.GetCount(unit, &result, count) //emps table not found

	unit := repository.NewUnitOfWork(q.Db, true)
	type Result struct {
		COUNT int `gorm:"type:int;column:COUNT"`
	}
	var result Result
	var queryp = []repository.QueryProcessor{repository.Table("emp"), repository.Select("count(empno) as COUNT")}
	q.Repo.GetAll(unit, &result, queryp)
	fmt.Println(result)
}

func (q *QueryService) GetSalSum() {
	fmt.Println("2)display sum of salaries of employees")
	//display sum of salaries of employees
	//row := q.Db.Debug().Table("emp").Select("sum(SAL)").Row()
	//row.Scan(sum)
	unit := repository.NewUnitOfWork(q.Db, true)
	type Result struct {
		SUM float64 `gorm:"type:decimal(19,4);column:SUM"`
	}
	var result Result
	var queryp = []repository.QueryProcessor{repository.Table("emp"), repository.Select("sum(sal) as SUM")}
	q.Repo.GetAll(unit, &result, queryp)
	fmt.Println(result)
}

func (q *QueryService) GetAvgSal() {
	fmt.Println("3)display avg salaries of employees")
	//display avg salaries of employees
	// row := q.Db.Debug().Table("emp").Select("avg(SAL)").Row()
	// row.Scan(avg)

	unit := repository.NewUnitOfWork(q.Db, true)
	type Result struct {
		AVG float64 `gorm:"type:decimal(19,4);column:AVG"`
	}
	var result Result
	var queryp = []repository.QueryProcessor{repository.Table("emp"), repository.Select("avg(sal) as AVG")}
	q.Repo.GetAll(unit, &result, queryp)
	fmt.Println(result)
}

func (q *QueryService) DisplaySumAvgCount() {
	fmt.Println("4)display sum,avg,count of employees")
	//display sum,avg,count of employees
	unit := repository.NewUnitOfWork(q.Db, true)
	type Result struct {
		SUM   float64 `gorm:"type:decimal(19,4);column:SUM"`
		AVG   float64 `gorm:"type:decimal(19,4);column:AVG"`
		COUNT int     `gorm:"type:int;column:COUNT"`
	}
	var result []Result
	var queryp []repository.QueryProcessor
	queryp = append(queryp, repository.Table("emp"))
	queryp = append(queryp, repository.Select("sum(SAL) AS SUM,avg(SAL) AS AVG,count(EMPNO) AS COUNT"))
	q.Repo.GetFirst(unit, &result, queryp)
	fmt.Println(result)
}

func (q *QueryService) DeptWiseHeadCount() {
	fmt.Println("5)display the dept wise , headcount")
	//display the dept wise , headcount
	unit := repository.NewUnitOfWork(q.Db, true)
	type Result struct {
		DEPTNO    int64  `gorm:"type:int;column:DEPTNO"`
		DNAME     string `gorm:"type:varchar(14);column:DNAME"`
		HEADCOUNT int64  `gorm:"column:HEADCOUNT""`
	}
	var result []Result
	var queryp []repository.QueryProcessor
	queryp = append(queryp, repository.Table("dept"))
	queryp = append(queryp, repository.Joins("INNER JOIN emp on emp.DEPTNO=dept.DEPTNO"))
	queryp = append(queryp, repository.Group("dept.DEPTNO"))
	queryp = append(queryp, repository.Select("DEPT.DEPTNO,DEPT.DNAME,COUNT(emp.EMPNO) AS HEADCOUNT"))
	q.Repo.GetAll(unit, &result, queryp)
	fmt.Println(result)
}

func (q *QueryService) JobWiseHeadCount() {
	fmt.Println("6)display the jobwise headcount")
	//display the jobwise headcount
	unit := repository.NewUnitOfWork(q.Db, true)
	type Result struct {
		JOB   string `gorm:"type:varchar(9);column:JOB"`
		COUNT int64  `gorm:"column:COUNT""`
	}
	var result []Result
	var queryp = []repository.QueryProcessor{repository.Table("emp"), repository.Select("JOB,COUNT(EMPNO) AS COUNT"), repository.Group("emp.JOB")}
	q.Repo.GetAll(unit, &result, queryp)
	fmt.Println(result)
}

func (q *QueryService) DeptwiseJobwiseHeadcount() {
	fmt.Println("7)display dept wise ,jobwise head count")
	//display dept wise ,jobwise head count
	unit := repository.NewUnitOfWork(q.Db, true)
	type Result struct {
		DEPTNO    int64  `gorm:"type:int;column:DEPTNO"`
		DNAME     string `gorm:"type:varchar(14);column:DNAME"`
		JOB       string `gorm:"type:varchar(9);column:JOB"`
		HEADCOUNT int64  `gorm:"column:HEADCOUNT""`
	}
	var result []Result
	var queryp = []repository.QueryProcessor{repository.Table("dept"), repository.Joins("INNER JOIN emp on emp.DEPTNO=dept.DEPTNO"), repository.Group("dept.DEPTNO,emp.JOB"), repository.Select("DEPT.DEPTNO,DEPT.DNAME,EMP.JOB,COUNT(emp.EMPNO) AS HEADCOUNT")}
	q.Repo.GetAll(unit, &result, queryp)
	fmt.Println(result)
}

func (q *QueryService) Test() {
	//SELECT ename,dname,job from dept left join emp on emp.DEPTNO = dept.DEPTNO;
	unit := repository.NewUnitOfWork(q.Db, true)
	type Result struct {
		ENAME string `gorm:"type:varchar(10);column:ENAME"`
		DNAME string `gorm:"type:varchar(14);column:DNAME"`
		JOB   string `gorm:"type:varchar(9);column:JOB"`
	}
	var result []Result
	var queryp = []repository.QueryProcessor{repository.Table("DEPT"), repository.Joins("LEFT JOIN EMP ON EMP.DEPTNO=DEPT.DEPTNO"), repository.Select("EMP.ENAME AS ENAME,DEPT.DNAME AS DNAME,EMP.JOB AS JOB")}
	q.Repo.GetAll(unit, &result, queryp)
	fmt.Println(result)
}

func (q *QueryService) DeptwiseEmployees() {
	fmt.Println("8)display the deptwise employees whose count greater than 2 and who are in dept 10 ,20.Sorty the result by descending order of count")
	//display the deptwise employees whose count greater than 2 and who are in dept 10 ,20.Sorty the result by descending order of count
	unit := repository.NewUnitOfWork(q.Db, true)
	type Result struct {
		DEPTNO    int64  `gorm:"type:int;column:DEPTNO"`
		DNAME     string `gorm:"type:varchar(14);column:DNAME"`
		JOB       string `gorm:"type:varchar(9);column:JOB"`
		HEADCOUNT int64  `gorm:"column:HEADCOUNT""`
	}
	var result []Result
	//db.Debug().Table("dept").Joins("INNER JOIN emp ON emp.DEPTNO=dept.DEPTNO").Group("dept.DEPTNO").Having("(DEPTNO=? OR DEPTNO=?) AND COUNT(EMP.EMPNO)>?", 10, 20, 2).Order("HEADCOUNT desc").Select("DEPT.DEPTNO,DEPT.DNAME,EMP.JOB,COUNT(emp.EMPNO) AS HEADCOUNT").Find(&resultVar7)
	var queryp = []repository.QueryProcessor{repository.Table("DEPT"), repository.Joins("INNER JOIN emp ON emp.DEPTNO=dept.DEPTNO"), repository.Group("dept.DEPTNO"), repository.Having("(DEPTNO=10 OR DEPTNO=20) AND COUNT(EMP.EMPNO)>2"), repository.Order("HEADCOUNT desc", true), repository.Select("DEPT.DEPTNO,DEPT.DNAME,EMP.JOB,COUNT(emp.EMPNO) AS HEADCOUNT")}
	q.Repo.GetAll(unit, &result, queryp)
	fmt.Println(result)
}

func (q *QueryService) DisplayEmpDept() {
	fmt.Println("9)display ename,deptname")
	//display ename,deptname
	unit := repository.NewUnitOfWork(q.Db, true)
	type Result struct {
		ENAME string `gorm:"type:varchar(10);column:ENAME"`
		DNAME string `gorm:"type:varchar(14);column:DNAME"`
	}
	var result []Result
	//db.Debug().Table("dept").Joins("LEFT OUTER JOIN emp ON emp.DEPTNO=dept.DEPTNO").Select("emp.ENAME,dept.DNAME").Find(&resultVar8)
	var queryp = []repository.QueryProcessor{repository.Table("dept"), repository.Joins("LEFT OUTER JOIN emp ON emp.DEPTNO=dept.DEPTNO"), repository.Select("emp.ENAME,dept.DNAME")}
	q.Repo.GetAll(unit, &result, queryp)
	fmt.Println(result)
}

func (q *QueryService) DisplayDeptEmp() {
	fmt.Println("12)display all the departments , with employees if any (if no emps then display null)")
	//display all the departments , with employees if any (if no emps then display null)
	unit := repository.NewUnitOfWork(q.Db, true)
	type Result struct {
		DNAME string `gorm:"type:varchar(14);column:DNAME"`
		ENAME string `gorm:"type:varchar(10);column:ENAME"`
	}
	var result []Result
	//db.Debug().Table("dept").Joins("LEFT OUTER JOIN emp ON emp.DEPTNO=dept.DEPTNO").Select("emp.ENAME,dept.DNAME").Find(&resultVar8)
	var queryp = []repository.QueryProcessor{repository.Table("dept"), repository.Joins("LEFT OUTER JOIN emp ON emp.DEPTNO=dept.DEPTNO"), repository.Select("dept.DNAME,emp.ENAME")}
	q.Repo.GetAll(unit, &result, queryp)
	fmt.Println(result)
}

func (q *QueryService) DisplayAllDeptsWithNoEmp() {
	fmt.Println("13)display the departments where there are no employees")
	//display the departments where there are no employees
	unit := repository.NewUnitOfWork(q.Db, true)
	type Result struct {
		DNAME string `gorm:"type:varchar(14);column:DNAME"`
	}
	var result_1 []Result
	var queryp_1 = []repository.QueryProcessor{repository.Table("dept"), repository.Joins("INNER JOIN emp ON emp.DEPTNO=dept.DEPTNO"), repository.Select("dept.DNAME")}
	q.Repo.GetAll(unit, &result_1, queryp_1)
	fmt.Println(result_1)
	var result_2 []Result
	var query string = "DNAME NOT IN ("
	for i, v := range result_1 {
		if i == len(result_1)-1 {
			query = query + "\"" + v.DNAME + "\""
		} else {
			query = query + "\"" + v.DNAME + "\", "
		}
	}
	query = query + ")"
	var queryp_2 = []repository.QueryProcessor{repository.Table("DEPT"), repository.Select("DNAME"), repository.Filter(query)}
	q.Repo.GetAll(unit, &result_2, queryp_2)
	fmt.Println(result_2)
	//q.Db.Debug().Raw(`SELECT DNAME FROM DEPT WHERE DNAME NOT IN (SELECT DEPT.DNAME FROM DEPT INNER JOIN EMP ON EMP.DEPTNO=DEPT.DEPTNO);`).Scan(&result)
}

func (q *QueryService) DisplayEmpBoss() {
	fmt.Println("14)display the empname and their bossnames")
	//display the empname and their bossnames
	unit := repository.NewUnitOfWork(q.Db, true)
	type Result struct {
		EMPNAME  string `gorm:"type:varchar(10);column:EMPNAME"`
		BOSSNAME string `gorm:"type:varchar(10);column:BOSSNAME"`
	}
	var result []Result
	//db.Debug().Table("EMP A, EMP B").Select("A.ENAME AS EMPNAME, B.ENAME AS BOSSNAME").Where("A.MGR=B.EMPNO").Scan(&resultVar10)
	var queryp = []repository.QueryProcessor{repository.Table("EMP A, EMP B"), repository.Select("A.ENAME AS EMPNAME, B.ENAME AS BOSSNAME"), repository.Filter("A.MGR=B.EMPNO")}
	q.Repo.GetAll(unit, &result, queryp)
	fmt.Println(result)
}

func (q *QueryService) DisplayEmpBossWithNulls() {
	fmt.Println("15)display all the empnames and boss names if any (if no boss display null)")
	//display all the empnames and boss names if any (if no boss display null)-SELECT A.ENAME AS EMPNAME, B.ENAME AS BOSSNAME FROM EMP A, EMP B WHERE A.MGR=B.EMPNO;
	unit := repository.NewUnitOfWork(q.Db, true)
	type Result struct {
		EMPNAME  string `gorm:"type:varchar(10);column:EMPNAME"`
		BOSSNAME string `gorm:"type:varchar(10);column:BOSSNAME"`
	}
	var result []Result
	//db.Debug().Table("EMP A, EMP B").Select("A.ENAME AS EMPNAME, B.ENAME AS BOSSNAME").Where("A.MGR=B.EMPNO").Scan(&resultVar10)
	var queryp = []repository.QueryProcessor{repository.Table("EMP A"), repository.Joins("LEFT JOIN EMP B ON A.MGR=B.EMPNO"), repository.Select("A.ENAME AS EMPNAME, B.ENAME AS BOSSNAME")}
	q.Repo.GetAll(unit, &result, queryp)
	fmt.Println(result)
}

func (q *QueryService) DisplayEmpBossDept() {
	fmt.Println("16)display ename,deptname and bossname")
	//display ename,deptname and bossname
	unit := repository.NewUnitOfWork(q.Db, true)
	type Result struct {
		EMPNAME  string `gorm:"type:varchar(10);column:EMPNAME"`
		DEPTNAME string `gorm:"type:varchar(10);column:DEPTNAME"`
		BOSSNAME string `gorm:"type:varchar(10);column:BOSSNAME"`
	}

	var result []Result
	//db.Debug().Table("EMP A,EMP B").Select("A.ENAME AS EMPNAME, DEPT.DNAME AS DEPTNAME, B.ENAME AS BOSSNAME").Joins("INNER JOIN DEPT ON B.DEPTNO=DEPT.DEPTNO").Where("A.MGR=B.EMPNO").Scan(&resultVar11)
	var queryp = []repository.QueryProcessor{repository.Table("EMP A, EMP B"), repository.Select("A.ENAME AS EMPNAME, DEPT.DNAME AS DEPTNAME, B.ENAME AS BOSSNAME"), repository.Joins("INNER JOIN DEPT ON B.DEPTNO=DEPT.DEPTNO"), repository.Filter("A.MGR=B.EMPNO")}
	q.Repo.GetAll(unit, &result, queryp)
	fmt.Println(result)
}

func (q *QueryService) DisplayRegionsWithNoCountry() {
	fmt.Println("17)display the regions there no entry for country")
	//display the regions there no entry for country
	unit := repository.NewUnitOfWork(q.Db, true)
	type Result1 struct {
		REGION_ID int `gorm:"column:REGION_ID"`
	}
	type Result2 struct {
		NAME string `gorm:"type:varchar(25);column:NAME"`
	}
	var result_1 []Result1
	//SELECT REGION_NAME FROM REGIONS WHERE REGION_ID NOT IN (SELECT REGIONS.REGION_ID FROM REGIONS INNER JOIN COUNTRIES ON REGIONS.REGION_ID=COUNTRIES.REGION_ID);
	var queryp_1 = []repository.QueryProcessor{repository.Table("REGIONS"), repository.Select("REGIONS.REGION_ID AS REGION_ID"), repository.Joins("INNER JOIN COUNTRIES ON REGIONS.REGION_ID=COUNTRIES.REGION_ID")}
	q.Repo.GetAll(unit, &result_1, queryp_1)
	var result_2 []Result2
	var query string = "REGION_ID NOT IN ("
	for i, v := range result_1 {
		if i == len(result_1)-1 {
			query = query + "\"" + strconv.Itoa(v.REGION_ID) + "\""
		} else {
			query = query + "\"" + strconv.Itoa(v.REGION_ID) + "\", "
		}
	}
	query = query + ")"
	var queryp_2 = []repository.QueryProcessor{repository.Table("REGIONS"), repository.Select("REGION_NAME AS NAME"), repository.Filter(query)}
	q.Repo.GetAll(unit, &result_2, queryp_2)
	fmt.Println(result_2)
}

func (q *QueryService) DisplayCountriesWithNoState() {
	fmt.Println("18)display the countries there no states")
	//display the countries there no states
	unit := repository.NewUnitOfWork(q.Db, true)
	type Result struct {
		NAME string `gorm:"type:varchar(40);column:NAME"`
	}
	var result_1 []Result
	//db.Debug().Table("COUNTRIES").Select("COUNTRIES.COUNTRY_NAME AS NAME").Joins("RIGHT OUTER JOIN LOCATIONS ON COUNTRIES.COUNTRY_ID=LOCATIONS.COUNTRY_ID").Scan(&resultVar12_1)
	var queryp_1 = []repository.QueryProcessor{repository.Table("COUNTRIES"), repository.Select("COUNTRIES.COUNTRY_NAME AS NAME"), repository.Joins("RIGHT OUTER JOIN LOCATIONS ON COUNTRIES.COUNTRY_ID=LOCATIONS.COUNTRY_ID")}
	q.Repo.GetAll(unit, &result_1, queryp_1)
	var result_2 []Result
	var query string = "COUNTRY_NAME NOT IN ("
	for i, v := range result_1 {
		if i == len(result_1)-1 {
			query = query + "\"" + v.NAME + "\""
		} else {
			query = query + "\"" + v.NAME + "\", "
		}
	}
	query = query + ")"
	//db.Debug().Table("COUNTRIES").Where(query).Select("COUNTRY_NAME AS NAME").Scan(&result_2)
	var queryp_2 = []repository.QueryProcessor{repository.Table("COUNTRIES"), repository.Select("COUNTRY_NAME AS NAME"), repository.Filter(query)}
	q.Repo.GetAll(unit, &result_2, queryp_2)
	fmt.Println(result_2)
}

func (q *QueryService) DisplayRegionCountryState() {
	fmt.Println("19)display region name,country name and state name")
	//display region name,country name and state name
	unit := repository.NewUnitOfWork(q.Db, true)
	type Result struct {
		COUNTRY_NAME   string `gorm:"type:varchar(40);column:COUNTRY_NAME"`
		STATE_PROVINCE string `gorm:"type:varchar(25);column:STATE_PROVINCE"`
		REGION_NAME    string `gorm:"type:varchar(25);column:REGION_NAME"`
	}

	var result []Result
	//db.Debug().Table("COUNTRIES").Select("COUNTRIES.COUNTRY_NAME, LOCATIONS.STATE_PROVINCE, REGIONS.REGION_NAME").Joins("INNER JOIN LOCATIONS ON COUNTRIES.COUNTRY_ID=LOCATIONS.COUNTRY_ID").Joins("INNER JOIN REGIONS ON COUNTRIES.REGION_ID=REGIONS.REGION_ID").Scan(&resultVar13)
	var queryp = []repository.QueryProcessor{repository.Table("COUNTRIES"), repository.Select("COUNTRIES.COUNTRY_NAME, LOCATIONS.STATE_PROVINCE, REGIONS.REGION_NAME"), repository.Joins("INNER JOIN LOCATIONS ON COUNTRIES.COUNTRY_ID=LOCATIONS.COUNTRY_ID"), repository.Joins("INNER JOIN REGIONS ON COUNTRIES.REGION_ID=REGIONS.REGION_ID")}
	q.Repo.GetAll(unit, &result, queryp)
	fmt.Println(result)

}

func (q *QueryService) CreateLocation() {
	fmt.Println("20)Make an insert of swabhav Techlabs in location/state tables map to india and asia.")
	//Make an insert of swabhav Techlabs in location/state tables map to india and asia.
	unit := repository.NewUnitOfWork(q.Db, true)
	var newLocation = model.Location{LOCATION_ID: 222, STREET_ADDRESS: "yyy", POSTAL_CODE: "2222", CITY: "Mumbai", STATE_PROVINCE: "Maharashtra", COUNTRY_ID: "IN"}
	q.Repo.Add(unit, &newLocation)
}

func (q *QueryService) GetMumbaiData() {
	fmt.Println("21)Filter details based on mumbai location")
	//Filter details based on mumbai location
	unit := repository.NewUnitOfWork(q.Db, true)
	var data []model.Location
	var queryp = []repository.QueryProcessor{repository.Table("locations"), repository.Filter("CITY=?", "Mumbai")}
	q.Repo.GetAll(unit, &data, queryp)
	fmt.Println(data)
}

func (q *QueryService) InsertIntoFoo() {
	fmt.Println("22)create a foo table and insert values of different data")
	//create a foo table and insert values of different data
	unit := repository.NewUnitOfWork(q.Db, true)

	q.Db.AutoMigrate(&model.Foo{})
	var newFoo = model.Foo{Name: "xyz", Age: 32}
	q.Repo.Add(unit, newFoo)
	var newFoo2 = model.Foo{Name: "abc", Age: 87}
	q.Repo.Add(unit, newFoo2)
}
