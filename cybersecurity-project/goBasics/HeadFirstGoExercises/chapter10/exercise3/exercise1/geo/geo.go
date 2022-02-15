package geo

import "errors"

type Coordinates struct {
	latitude  float64
	longitude float64
}

type Landmark struct {
	name string
	Coordinates
}

func (l *Landmark) Name() string {
	return l.name
}

func (l *Landmark) SetName(name string) error {
	if name == "" {
		return errors.New("invalid name")
	}
	l.name = name
	return nil

}

func (c *Coordinates) Latitude() float64 {
	return c.latitude
}

func (c *Coordinates) Longitude() float64 {
	return c.longitude
}

func (c *Coordinates) SetLatitude(latitude float64) error {
	if latitude < -90.0 || latitude > 90.0 {
		return errors.New("Invalide latitude")
	}
	c.latitude = latitude
	return nil
}

func (c *Coordinates) SetLongitude(longitude float64) error {
	if longitude < -180.0 || longitude > 180.0 {
		return errors.New("Invalid longitude")
	}
	c.longitude = longitude
	return nil
}
