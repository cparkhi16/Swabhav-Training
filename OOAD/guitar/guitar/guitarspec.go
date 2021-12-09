package guitar

import (
	b "app/builder"
	e "app/guitarType"
	w "app/wood"
	"fmt"
)

type GuitarSpec struct {
	model                    string
	builder                  b.Builder
	typeofGuitar             e.GuitarType
	backWood                 w.Wood
	frontWood                w.Wood
	noOfStrings              uint8
	isRefundable             bool
	isRefundableSetForSearch int
}

func NewGuitarSpec(model string, builder b.Builder, typeofGuitar e.GuitarType, backWood w.Wood, frontWood w.Wood, noOfStrings uint8, i bool) *GuitarSpec {
	return &GuitarSpec{
		model:        model,
		builder:      builder,
		typeofGuitar: typeofGuitar,
		backWood:     backWood,
		frontWood:    frontWood,
		noOfStrings:  noOfStrings,
		isRefundable: i,
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
	j := int(b)
	if j > 3 {
		return fmt.Errorf("please give valid builder type")
	}
	gs.builder = b

	return nil
}
func (gs *GuitarSpec) SetTypeOfGuitar(b e.GuitarType) error {
	s := []string{"Accoustic", "Electric"}
	var found bool = false
	for _, val := range s {
		if b.String() == val {
			found = true
		}
	}
	if !found {
		return fmt.Errorf("please give valid guitar type")
	}
	gs.typeofGuitar = b

	return nil
}
func (gs *GuitarSpec) SetBackWood(b w.Wood) error {
	s := []string{"Mahagony", "Maple", "Cocobolo", "Cedar", "Sitka"}
	var found bool = false
	for _, val := range s {
		if b.String() == val {
			found = true
		}
	}
	if !found {
		return fmt.Errorf("please give valid wood type")
	}
	gs.backWood = b

	return nil
}
func (gs *GuitarSpec) SetFrontWood(b w.Wood) error {
	s := []string{"Mahagony", "Maple", "Cocobolo", "Cedar", "Sitka"}
	var found bool = false
	for _, val := range s {
		if b.String() == val {
			found = true
		}
	}
	if !found {
		return fmt.Errorf("please give valid wood type")
	}
	gs.backWood = b
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
func (gs *GuitarSpec) SetRefundable(b bool) {
	gs.isRefundable = b
	gs.isRefundableSetForSearch = 1
}
func (gs GuitarSpec) GetIsRefundableVal() (bool, int) {
	return gs.isRefundable, gs.isRefundableSetForSearch
}
