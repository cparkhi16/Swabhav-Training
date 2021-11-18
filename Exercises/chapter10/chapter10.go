package main

import (
	geo "chapter10/geo"
	"fmt"
	"log"
)

func main() {
	coordinates := geo.Coordinates{}
	err := coordinates.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}
	err = coordinates.SetLongitude(-122.3)
	if err != nil {
		log.Fatal(err)
	}
	coordinates.SetLatitude(37.42)
	coordinates.SetLongitude(-122.3)

	fmt.Println(coordinates.Latitude())
	fmt.Println(coordinates.Longitude())

	location := geo.Landmark{}
	err1 := location.SetName("The Googleaplex")
	if err1 != nil {
		log.Fatal(err)
	}
	err = location.SetLatitude(37.4)
	if err != nil {
		log.Fatal(err)
	}
	err = location.SetLongitude(-87.4)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(location.Name())
	fmt.Println(location.Latitude())
	fmt.Println(location.Longitude())
}
