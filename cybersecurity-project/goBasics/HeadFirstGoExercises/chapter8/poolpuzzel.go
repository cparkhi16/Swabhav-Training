package main

import (
	"fmt"
)

type coordinates struct {
	Latitude  float64
	Longitude float64
}

func main() {
	location := coordinates{Latitude: 37.42, Longitude: -122.08}
	fmt.Println("Latitude:", location.Latitude)
	fmt.Println("Longitude:", location.Longitude)
}
