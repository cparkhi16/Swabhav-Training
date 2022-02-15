package model

type Foo struct {
	Name string `gorm:"type:varchar(30);column:name"`
	Age  int    `gorm:"type:int;column:age"`
}
