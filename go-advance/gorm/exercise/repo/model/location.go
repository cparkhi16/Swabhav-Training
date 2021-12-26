package model

type Location struct {
	CountryID     string `gorm:"column:COUNTRY_ID"`
	StreetAddress string `gorm:"column:STREET_ADDRESS"`
	LocationID    int    `gorm:"column:LOCATION_ID"`
	City          string `gorm:"column:CITY"`
	StateProvince string `gorm:"column:STATE_PROVINCE"`
	PostalCode    string `gorm:"column:POSTAL_CODE"`
}

func (Location) TableName() string {
	return "locations"
}

func NewLocation(LocationID int, c, s, city, state, postal string) *Location {
	return &Location{LocationID: LocationID, CountryID: c, StreetAddress: s, City: city, StateProvince: state, PostalCode: postal}
}
