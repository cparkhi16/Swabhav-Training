package main

import (
	b "app/builder"
	g "app/guitar"
	e "app/guitarType"
	in "app/inventory"
	w "app/wood"
)

func main() {
	i := in.GetInventory()
	var refundable bool = true
	var nonRefundable bool = false
	models := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	builders := []b.Builder{b.Fender, b.Martin, b.Gibson}
	strings := []uint8{19, 10, 18, 20, 25, 21, 21, 12, 20, 34}
	s := []w.Wood{w.Mahagony, w.Maple, w.Cocobolo, w.Cedar, w.Sitka}
	t := []e.GuitarType{e.Accoustic, e.Electric}
	gs := g.NewGuitarSpec(&models[0], &builders[0], &t[0], &s[1], &s[0], &strings[0], &refundable)
	guitarOne := g.NewGuitar("123", 1200, *gs)
	gst := g.NewGuitarSpec(&models[1], &builders[1], &t[0], &s[3], &s[1], &strings[1], &nonRefundable)
	guitarTwo := g.NewGuitar("211", 1900, *gst)
	gsc := g.NewGuitarSpec(&models[2], &builders[2], &t[1], &s[0], &s[1], &strings[2], &refundable)
	guitarThree := g.NewGuitar("212", 1100, *gsc)
	gsm := g.NewGuitarSpec(&models[3], &builders[2], &t[0], &s[1], &s[2], &strings[3], &nonRefundable)
	guitarFour := g.NewGuitar("213", 1700, *gsm)
	gh := g.NewGuitarSpec(&models[4], &builders[1], &t[0], &s[4], &s[2], &strings[4], &nonRefundable)
	guitarFive := g.NewGuitar("222", 1000, *gh)
	gt := g.NewGuitarSpec(&models[5], &builders[1], &t[1], &s[1], &s[1], &strings[5], &refundable)
	guitarSix := g.NewGuitar("234", 1500, *gt)
	ga := g.NewGuitarSpec(&models[6], &builders[2], &t[0], &s[4], &s[2], &strings[6], &nonRefundable)
	guitarSeven := g.NewGuitar("238", 1990, *ga)
	gb := g.NewGuitarSpec(&models[7], &builders[0], &t[1], &s[2], &s[4], &strings[7], &refundable)
	guitarEight := g.NewGuitar("290", 1590, *gb)
	gd := g.NewGuitarSpec(&models[8], &builders[1], &t[0], &s[0], &s[1], &strings[8], &nonRefundable)
	guitarNine := g.NewGuitar("298", 1510, *gd)
	ge := g.NewGuitarSpec(&models[9], &builders[2], &t[1], &s[3], &s[4], &strings[9], &nonRefundable)
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
	//err := gm.SetBuilder(b.Fender)
	/*if err != nil {
		log.Fatal(err.Error())
	}*/
	//e := gm.SetTypeOfGuitar(60001)
	//gm.SetNoOfStrings(21)
	/*e := gm.SetBackWood(&s[0])
	if e != nil {
		log.Fatal(e.Error())
	}*/
	gm.SetRefundable(&refundable)
	i.SearchGuitar(*gm)
	//fmt.Println("Is there any specs match ? -- ", gs.Matches(*gm))
	//i.GetGuitarsFromInventory()
}
