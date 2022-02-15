package main

import "fmt"

type coordinates struct {
	Latitude  float64
	Longitude float64
}

type landmark struct {
	name string
	coordinates
}

func main() {
	location := landmark{}
	location.name = "The Googleplex"
	location.Latitude = 37.42
	location.Longitude = -122.08
	fmt.Println(location)
}
