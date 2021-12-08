package main

import (
	b "app/builder"
	g "app/guitar"
	e "app/guitarType"
	in "app/inventory"
	w "app/wood"
	"log"
)

func main() {
	i := in.GetInventory()
	gs := g.NewGuitarSpec("A", b.Fender, e.Accoustic, w.Maple, w.Mahagony, 19)
	guitarOne := g.NewGuitar("123", 1200, *gs)
	gst := g.NewGuitarSpec("B", b.Martin, e.Accoustic, w.Cedar, w.Maple, 10)
	guitarTwo := g.NewGuitar("211", 1900, *gst)
	gsc := g.NewGuitarSpec("C", b.Gibson, e.Electric, w.Mahagony, w.Maple, 18)
	guitarThree := g.NewGuitar("212", 1100, *gsc)
	gsm := g.NewGuitarSpec("D", b.Gibson, e.Accoustic, w.Maple, w.Cocobolo, 20)
	guitarFour := g.NewGuitar("213", 1700, *gsm)
	gh := g.NewGuitarSpec("E", b.Martin, e.Accoustic, w.Sitka, w.Cocobolo, 25)
	guitarFive := g.NewGuitar("222", 1000, *gh)
	gt := g.NewGuitarSpec("F", b.Martin, e.Electric, w.Maple, w.Maple, 21)
	guitarSix := g.NewGuitar("234", 1500, *gt)
	ga := g.NewGuitarSpec("G", b.Gibson, e.Accoustic, w.Sitka, w.Cocobolo, 21)
	guitarSeven := g.NewGuitar("238", 1990, *ga)
	gb := g.NewGuitarSpec("H", b.Fender, e.Electric, w.Cocobolo, w.Sitka, 21)
	guitarEight := g.NewGuitar("290", 1590, *gb)
	gd := g.NewGuitarSpec("I", b.Martin, e.Accoustic, w.Mahagony, w.Maple, 20)
	guitarNine := g.NewGuitar("298", 1510, *gd)
	ge := g.NewGuitarSpec("J", b.Gibson, e.Electric, w.Cedar, w.Sitka, 34)
	guitarTen := g.NewGuitar("208", 1515, *ge)
	var guitars []g.Guitar
	guitars = append(guitars, *guitarOne, *guitarTwo, *guitarThree, *guitarFour, *guitarFive, *guitarSix, *guitarSeven, *guitarEight,
		*guitarNine, *guitarTen)
	for _, val := range guitars {
		i.AddGuitar(val)
	}
	//i.RemoveGuitar("123")
	//j := in.GetInventory()
	//j.GetGuitarsFromInventory()
	//i.GetGuitarsFromInventory()
	//fmt.Println(i.GetGuitar("123"))
	gm := &g.GuitarSpec{}
	//gm.SetModel("A")
	err := gm.SetBuilder(b.Fender)
	if err != nil {
		log.Fatal(err.Error())
	}
	e := gm.SetTypeOfGuitar(e.Accoustic)
	gm.SetNoOfStrings(21)
	//e := gm.SetBackWood(w.Maple)
	if e != nil {
		log.Fatal(e.Error())
	}
	i.SearchGuitar(*gm)
	//fmt.Println("Is there any specs match ? -- ", gs.Matches(*gm))
	//i.GetGuitarsFromInventory()
}
