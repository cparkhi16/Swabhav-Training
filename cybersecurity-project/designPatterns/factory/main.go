package main

import (
	"carfactory/automobilefactory"
)

func main() {
	newCar := automobilefactory.MakeCar(automobilefactory.Mercedes, "Sedan")
	newCar.Start()
	newCar.Stop()
}
