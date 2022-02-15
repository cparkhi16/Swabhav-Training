package model

type Location struct {
	LOCATION_ID    int    `gorm:"type:int;column:LOCATION_ID"`
	STREET_ADDRESS string `gorm:"type:varchar(40);column:STREET_ADDRESS"`
	POSTAL_CODE    string `gorm:"type:varchar(12);column:POSTAL_CODE"`
	CITY           string `gorm:"type:varchar(30);column:CITY"`
	STATE_PROVINCE string `gorm:"type:varchar(25);column:STATE_PROVINCE"`
	COUNTRY_ID     string `gorm:"type:char(2);column:COUNTRY_ID"`
}
