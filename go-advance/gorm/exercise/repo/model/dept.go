package model

type Department struct {
	DeptName string `gorm:"column:DNAME"`
	DeptNo   string `gorm:"column:DEPTNO"`
	Location string `gorm:"column:LOC"`
}

func (Department) TableName() string {
	return "dept"
}
