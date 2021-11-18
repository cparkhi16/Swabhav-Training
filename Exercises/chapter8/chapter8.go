package main

import (
	"chapter8/geo"
	"fmt"
)

type student struct {
	name  string
	grade float64
}
type car struct {
	name     string
	topSpeed float64
}
type part struct {
	description string
	count       int
}

func doublePack(p *part) {
	p.count *= 2
}
func nitrBoost(c *car) {
	c.topSpeed += 50
}

func stru() {
	var pet struct {
		name string
		age  int
	}
	pet.name = "max"
	pet.age = 5
	fmt.Println(pet.name)
	fmt.Println(pet.age)
}
func printInfo(s student) {
	fmt.Println("name:", s.name)
	fmt.Printf("grade :%0.1f\n", s.grade)
}
func main() {
	stru()
	var s student
	s.name = "alo"
	s.grade = 98.7
	printInfo(s)
	var mustang car
	mustang.name = "Mustang cobra"
	mustang.topSpeed = 225
	nitrBoost(&mustang)
	fmt.Println(mustang.name)
	fmt.Println(mustang.topSpeed)
	var fuses part
	fuses.description = "fuses"
	fuses.count = 5
	doublePack(&fuses)
	fmt.Println(fuses.description)
	fmt.Println(fuses.count)
	location := geo.Coordinates{Latitude: 37.4, Longitude: -122.8}
	fmt.Println("latitude:", location.Latitude)
	fmt.Println("longitude:", location.Longitude)
	location2 := geo.Landmark{}
	location2.Name = "the googleplex"
	location2.Latitude = 32.4
	location2.Longitude = -123.4
	fmt.Println(location2)
}
