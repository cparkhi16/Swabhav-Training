package model

type Country struct {
	CountryID   string `gorm:"column:COUNTRY_ID"`
	CountryName string `gorm:"column:COUNTRY_NAME"`
	RegionID    int    `gorm:"column:REGION_ID"`
}

func (Country) TableName() string {
	return "countries"
}
