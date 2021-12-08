package inventory

import (
	g "app/guitar"
	"fmt"
)

type Inventory struct {
	allGuitars []g.Guitar
}

var instance *Inventory

func GetInventory() *Inventory {
	if instance == nil {
		instance = &Inventory{}
	}
	return instance
}
func (i *Inventory) AddGuitar(g g.Guitar) {
	i.allGuitars = append(i.allGuitars, g)
}
func (i *Inventory) RemoveGuitar(number string) {
	index := 0
	for j, val := range i.allGuitars {
		if val.GetSerialNumber() == number {
			index = j
			break
		}
	}
	i.allGuitars[index] = i.allGuitars[len(i.allGuitars)-1]
	i.allGuitars = i.allGuitars[:len(i.allGuitars)-1]
}

func (i *Inventory) GetGuitar(number string) g.Guitar {
	index := 0
	for j, val := range i.allGuitars {
		if val.GetSerialNumber() == number {
			index = j
			break
		}
	}
	return i.allGuitars[index]
}

func (i *Inventory) SearchGuitar(spec g.GuitarSpec) {
	fmt.Println(spec)
	model := spec.GetModel()
	builder := spec.GetBuilder()
	typeofGuitar := spec.GetTypeOfGuitar()
	backWood := spec.GetBackWood()
	frontWood := spec.GetFrontWood()
	noOfStrings := spec.GetNumberOfStrings()
	recordFound := false
	for _, val := range i.allGuitars {
		if (val.GetSpecs().GetBuilder() == builder) || (val.GetSpecs().GetTypeOfGuitar() == typeofGuitar) || (val.GetSpecs().GetModel() == model) ||
			(val.GetSpecs().GetBackWood() == backWood) || (val.GetSpecs().GetFrontWood() == frontWood) || (val.GetSpecs().GetNumberOfStrings() == noOfStrings) {
			fmt.Println("----------Search Results------------")
			fmt.Println("Matched guitars in inventory based on your search ")
			fmt.Println("Serial no ", val.GetSerialNumber())
			fmt.Println("Builder type ", val.GetSpecs().GetBuilder())
			fmt.Println("Guitar Type ", val.GetSpecs().GetTypeOfGuitar())
			fmt.Println("Guitar model ", val.GetSpecs().GetModel())
			fmt.Println("Type of backWood ", val.GetSpecs().GetBackWood())
			fmt.Println("Type of frontWood ", val.GetSpecs().GetFrontWood())
			fmt.Println("Number of strings ", val.GetSpecs().GetNumberOfStrings())
			fmt.Println("Guitar Price --", val.GetPrice())
			recordFound = true
		}
	}
	if !recordFound {
		fmt.Println("No record found based on your search")
	}
}
func (i *Inventory) GetGuitarsFromInventory() {
	for _, val := range i.allGuitars {
		fmt.Println("-- Guitar --")
		fmt.Println("Serial no ", val.GetSerialNumber())
		fmt.Println("Builder type ", val.GetSpecs().GetBuilder())
		fmt.Println("Guitar Type ", val.GetSpecs().GetTypeOfGuitar())
		fmt.Println("Guitar model ", val.GetSpecs().GetModel())
		fmt.Println("Type of backWood ", val.GetSpecs().GetBackWood())
		fmt.Println("Type of frontWood ", val.GetSpecs().GetFrontWood())
		fmt.Println("Number of strings ", val.GetSpecs().GetNumberOfStrings())
		fmt.Println("Guitar Price --", val.GetPrice())
	}
}
