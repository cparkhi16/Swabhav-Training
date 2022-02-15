package automobilefactory

import (
	"carfactory/automobile"
)

type Car string

const (
	BMW      Car = "bmw"
	Mercedes Car = "mercedes"
	Tesla    Car = "tesla"
)

func MakeCar(brand Car, modelName string) automobile.Automobile {
	switch brand {
	case BMW:
		return automobile.NewBmw(modelName)
	case Mercedes:
		return automobile.NewMercedes(modelName)
	case Tesla:
		return automobile.NewTesla(modelName)
	}
	return nil
}
