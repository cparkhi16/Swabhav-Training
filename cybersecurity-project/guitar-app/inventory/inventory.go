package inventory

import (
	"fmt"
	"guitarManagement/guitar"
)

type inventory struct {
	allGuitars []*guitar.Guitar
}

func New() *inventory {
	return &inventory{}
}

func (i *inventory) AddGuitar(g *guitar.Guitar) {
	i.allGuitars = append(i.allGuitars, g)
}

/*func remove(s []int, i int) []int {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}*/

func (i *inventory) RemoveGuitar(serialNo string) {
	indexToRemove := -1
	for j, v := range i.allGuitars {
		if serialNo == v.GetSerialNo() {
			indexToRemove = j
			break
		}
	}
	if indexToRemove != -1 {
		i.allGuitars[indexToRemove] = i.allGuitars[len(i.allGuitars)-1]
		i.allGuitars = i.allGuitars[:len(i.allGuitars)-1]
	} else {
		fmt.Println("Serial number not found in inventory")
	}
}

func (i *inventory) GetGuitar(serialNo string) *guitar.Guitar {
	for _, v := range i.allGuitars {
		if serialNo == v.GetSerialNo() {
			return v
		}
	}
	return nil
}

func (i *inventory) DisplayInventory() {
	for j, v := range i.allGuitars {
		fmt.Println("guitar", j, "-", *v)
	}
}

func (i *inventory) SearchGuitar(g *guitar.Guitar) []*guitar.Guitar {
	var matchedGuitars []*guitar.Guitar
	isSerialNoEmpty := g.IsSerialNoEmpty()
	isPriceEmpty := g.IsPriceEmpty()
	for _, v := range i.allGuitars {
		if isSerialNoEmpty {
			g.SetSerialNo(v.GetSerialNo())
		}
		if isPriceEmpty {
			g.SetPrice(v.GetPrice())
		}
		if g.GetPrice() == v.GetPrice() && g.GetSerialNo() == v.GetSerialNo() && v.GetSpec().Matches(g.GetSpec()) {
			matchedGuitars = append(matchedGuitars, v)
		}
	}
	return matchedGuitars

}
