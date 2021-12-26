package model

import (
	"time"
)

type Employee struct {
	ManagerName   string
	DeptWiseCount int
	JobWiseCount  int
	DeptName      string
	EmpName       string    `gorm:"column:ENAME"`
	EmpNo         uint      `gorm:"column:EMPNO"`
	Job           string    `gorm:"column:JOB"`
	MGR           uint      `gorm:"column:MGR"`
	DeptNo        uint      `gorm:"column:DEPTNO"`
	Comm          uint      `gorm:"column:COMM"`
	Salary        uint      `gorm:"column:SAL"`
	HiredDate     time.Time `gorm:"column:HIREDATE"`
}

func (Employee) TableName() string {
	return "emp"
}
