package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Emp struct {
	EmpNo    int     `gorm:"primary_key;column:EMPNO"`
	EName    string  `gorm:"type:varchar(10);column:EName"`
	Job      string  `gorm:"type:varchar(9);column:JOB`
	Mgr      int     `gorm:"column:MGR"`
	HireDate string  `gorm:"type:date;column:HIREDATE"`
	Sal      float64 `gorm:"type:decimal(19,4);column:SAL"`
	Comm     float64 `gorm:"type:decimal(19,4);column:COMM"`
	DeptNo   int     `gorm:"column:DEPTNO"`
}
