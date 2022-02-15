package main

import (
	"fmt"
	"guitarManagement/guitar"
	"guitarManagement/inventory"
)

func main() {
	s1 := guitar.NewGuitarSpec("model1", guitar.Fender, guitar.Accoustic, guitar.Mahogony, guitar.Sitka, 6, true)
	g1 := guitar.NewGuitar("123", 20, s1)

	s2 := guitar.NewGuitarSpec("model2", guitar.Martin, guitar.Accoustic, guitar.Maple, guitar.Mahogony, 12, true)
	g2 := guitar.NewGuitar("456", 200, s2)

	s3 := guitar.NewGuitarSpec("model3", guitar.Fender, guitar.Electric, guitar.Sitka, guitar.Cocobolo, 8, true)
	g3 := guitar.NewGuitar("789", 50, s3)

	s4 := guitar.NewGuitarSpec("model4", guitar.Martin, guitar.Electric, guitar.Cocobolo, guitar.Sitka, 12, false)
	g4 := guitar.NewGuitar("1011", 70, s4)

	s5 := guitar.NewGuitarSpec("model5", guitar.Gibson, guitar.Accoustic, guitar.Maple, guitar.Sitka, 6, true)
	g5 := guitar.NewGuitar("1213", 20, s5)

	s6 := guitar.NewGuitarSpec("model6", guitar.Martin, guitar.Accoustic, guitar.Cocobolo, guitar.Maple, 15, false)
	g6 := guitar.NewGuitar("1415", 30, s6)

	s7 := guitar.NewGuitarSpec("model7", guitar.Fender, guitar.Electric, guitar.Maple, guitar.Cocobolo, 20, true)
	g7 := guitar.NewGuitar("1617", 240, s7)

	s8 := guitar.NewGuitarSpec("model8", guitar.Martin, guitar.Electric, guitar.Cocobolo, guitar.Maple, 6, true)
	g8 := guitar.NewGuitar("1819", 320, s8)

	s9 := guitar.NewGuitarSpec("model9", guitar.Gibson, guitar.Electric, guitar.Sitka, guitar.Mahogony, 37, true)
	g9 := guitar.NewGuitar("2021", 1110, s9)

	s10 := guitar.NewGuitarSpec("model10", guitar.Fender, guitar.Accoustic, guitar.Sitka, guitar.Maple, 12, false)
	g10 := guitar.NewGuitar("2223", 100, s10)

	s11 := guitar.NewGuitarSpec("model11", guitar.Gibson, guitar.Accoustic, guitar.Maple, guitar.Maple, 6, false)
	g11 := guitar.NewGuitar("2324", 20, s11)

	s12 := guitar.NewGuitarSpec("model12", guitar.Martin, guitar.Electric, guitar.Maple, guitar.Cocobolo, 60, false)
	g12 := guitar.NewGuitar("2526", 20, s12)

	s13 := guitar.NewGuitarSpec("model13", guitar.Martin, guitar.Accoustic, guitar.Mahogony, guitar.Mahogony, 45, true)
	g13 := guitar.NewGuitar("2728", 2000, s13)

	s14 := guitar.NewGuitarSpec("model14", guitar.Fender, guitar.Electric, guitar.Maple, guitar.Maple, 29, false)
	g14 := guitar.NewGuitar("2930", 1020, s14)

	s15 := guitar.NewGuitarSpec("model15", guitar.Fender, guitar.Accoustic, guitar.Mahogony, guitar.Cocobolo, 6, false)
	g15 := guitar.NewGuitar("3132", 90, s15)

	shop := inventory.New()
	shop.AddGuitar(g1)
	shop.AddGuitar(g2)
	shop.AddGuitar(g3)
	shop.AddGuitar(g4)
	shop.AddGuitar(g5)
	shop.AddGuitar(g6)
	shop.AddGuitar(g7)
	shop.AddGuitar(g8)
	shop.AddGuitar(g9)
	shop.AddGuitar(g10)
	shop.AddGuitar(g11)
	shop.AddGuitar(g12)
	shop.AddGuitar(g13)
	shop.AddGuitar(g14)
	shop.AddGuitar(g15)
	fmt.Println("Inventory-")
	shop.DisplayInventory()
	newg := shop.GetGuitar("123")
	fmt.Println(newg)

	//search guitar try1
	ss := &guitar.GuitarSpec{}
	ss.SetBackWood(guitar.Mahogony)
	//ss.SetIsRefundable(false)
	gs := &guitar.Guitar{}
	gs.SetSpec(*ss)
	fmt.Println("After Search-")
	list := shop.SearchGuitar(gs)
	for _, l := range list {
		fmt.Println(l)
	}
	fmt.Println(list)

	//search guitar try2

}
