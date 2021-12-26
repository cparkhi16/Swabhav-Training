package model

type Region struct {
	RegionID   int    `gorm:"column:REGION_ID"`
	RegionName string `gorm:"column:REGION_NAME"`
}

func (Region) TableName() string {
	return "regions"
}
