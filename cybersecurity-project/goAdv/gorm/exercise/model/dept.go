package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type dept struct {
	DeptNo int    `gorm:"column:DEPTNO"`
	DName  string `gorm:"type:varchar(14);column:DNAME"`
	Loc    string `gorm:"type:varchar(13);column:LOC"`
}
