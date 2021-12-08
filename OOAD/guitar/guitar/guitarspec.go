package guitar

import (
	b "app/builder"
	e "app/guitarType"
	w "app/wood"
	"fmt"
)

type GuitarSpec struct {
	model        string
	builder      b.Builder
	typeofGuitar e.GuitarType
	backWood     w.Wood
	frontWood    w.Wood
	noOfStrings  uint8
}

func NewGuitarSpec(model string, builder b.Builder, typeofGuitar e.GuitarType, backWood w.Wood, frontWood w.Wood, noOfStrings uint8) *GuitarSpec {
	return &GuitarSpec{
		model:        model,
		builder:      builder,
		typeofGuitar: typeofGuitar,
		backWood:     backWood,
		frontWood:    frontWood,
		noOfStrings:  noOfStrings,
	}
}
func (gs GuitarSpec) GetModel() string {
	if gs.model == "" {
		return ""
	}
	return gs.model
}
func (gs GuitarSpec) GetBuilder() string {
	return gs.builder.String()
}
func (gs GuitarSpec) GetTypeOfGuitar() string {
	return gs.typeofGuitar.String()
}
func (gs GuitarSpec) GetBackWood() string {
	return gs.backWood.String()
}
func (gs GuitarSpec) GetFrontWood() string {
	return gs.frontWood.String()
}
func (gs GuitarSpec) GetNumberOfStrings() uint8 {
	if gs.noOfStrings == 0 {
		return 0
	}
	return gs.noOfStrings
}
func (gs *GuitarSpec) SetModel(m string) {
	gs.model = m
}
func (gs *GuitarSpec) SetBuilder(b b.Builder) error {
	/*fmt.Println(b.String())
	//gs.builder = b
	s := []string{"Fender", "Martin", "Gibson"}
	var found bool = false
	for _, val := range s {
		//fmt.Println("Val ", val)
		//fmt.Println("String", b.String())
		if b.String() == val {
			//fmt.Println("True")
			found = true
		}
	}*/
	j := int(b)
	if j > 3 {
		return fmt.Errorf("please give valid builder type")
	} else {
		//fmt.Println("Here")
		gs.builder = b
	}
	return nil
}
func (gs *GuitarSpec) SetTypeOfGuitar(b e.GuitarType) error {
	//gs.typeofGuitar = b
	//fmt.Println(b.String())
	s := []string{"Accoustic", "Electric"}
	var found bool = false
	for _, val := range s {
		//fmt.Println("Val ", val)
		//fmt.Println("String", b.String())
		if b.String() == val {
			//fmt.Println("True")
			found = true
		}
	}
	if !found {
		return fmt.Errorf("please give valid guitar type")
	} else {
		//fmt.Println("Here")
		gs.typeofGuitar = b
	}
	return nil
}
func (gs *GuitarSpec) SetBackWood(b w.Wood) error {
	//fmt.Println(b.String())
	s := []string{"Mahagony", "Maple", "Cocobolo", "Cedar", "Sitka"}
	var found bool = false
	for _, val := range s {
		//fmt.Println("Val ", val)
		//fmt.Println("String", b.String())
		if b.String() == val {
			//fmt.Println("True")
			found = true
		}
	}
	if !found {
		return fmt.Errorf("please give valid wood type")
	} else {
		//fmt.Println("Here")
		gs.backWood = b
	}
	return nil
}
func (gs *GuitarSpec) SetFrontWood(b w.Wood) error {
	//gs.frontWood = b
	s := []string{"Mahagony", "Maple", "Cocobolo", "Cedar", "Sitka"}
	var found bool = false
	for _, val := range s {
		//fmt.Println("Val ", val)
		//fmt.Println("String", b.String())
		if b.String() == val {
			//fmt.Println("True")
			found = true
		}
	}
	if !found {
		return fmt.Errorf("please give valid wood type")
	} else {
		//fmt.Println("Here")
		gs.backWood = b
	}
	return nil
}
func (gs *GuitarSpec) SetNoOfStrings(b uint8) {
	gs.noOfStrings = b
}
func (gs *GuitarSpec) Matches(otherSpec GuitarSpec) bool {
	if gs.backWood.String() == otherSpec.backWood.String() || gs.builder.String() == otherSpec.builder.String() ||
		gs.frontWood.String() == otherSpec.frontWood.String() || gs.model == otherSpec.model || gs.noOfStrings == otherSpec.noOfStrings ||
		gs.typeofGuitar == otherSpec.typeofGuitar {
		return true
	}
	return false
}
