package guitar

import (
	b "app/builder"
	e "app/guitarType"
	w "app/wood"
	"fmt"
)

type GuitarSpec struct {
	model        *string
	builder      *b.Builder
	typeofGuitar *e.GuitarType
	backWood     *w.Wood
	frontWood    *w.Wood
	noOfStrings  *uint8
	isRefundable *bool
	//isRefundableSetForSearch int
}

func NewGuitarSpec(model *string, builder *b.Builder, typeofGuitar *e.GuitarType, backWood *w.Wood, frontWood *w.Wood, noOfStrings *uint8, i *bool) *GuitarSpec {
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
	if gs.model == nil {
		return ""
	}
	return *gs.model
}
func (gs GuitarSpec) GetBuilder() string {
	if gs.builder == nil {
		return ""
	}
	return gs.builder.String()
}
func (gs GuitarSpec) GetTypeOfGuitar() string {
	if gs.typeofGuitar == nil {
		return ""
	}
	return gs.typeofGuitar.String()
}
func (gs GuitarSpec) GetBackWood() string {
	if gs.backWood == nil {
		return ""
	}
	return gs.backWood.String()
}
func (gs GuitarSpec) GetFrontWood() string {
	if gs.frontWood == nil {
		return ""
	}
	return gs.frontWood.String()
}
func (gs GuitarSpec) GetNumberOfStrings() uint8 {
	if gs.noOfStrings == nil {
		return 100
	}
	return *gs.noOfStrings
}
func (gs *GuitarSpec) SetModel(m *string) {
	gs.model = m
}
func (gs *GuitarSpec) SetBuilder(b *b.Builder) error {
	j := int(*b)
	if j > 3 {
		return fmt.Errorf("please give valid builder type")
	}
	gs.builder = b

	return nil
}
func (gs *GuitarSpec) SetTypeOfGuitar(b *e.GuitarType) error {
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
func (gs *GuitarSpec) SetBackWood(b *w.Wood) error {
	fmt.Println(" String ", b.String())
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
func (gs *GuitarSpec) SetFrontWood(b *w.Wood) error {
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
func (gs *GuitarSpec) SetNoOfStrings(b *uint8) {
	gs.noOfStrings = b
}
func (gs *GuitarSpec) Matches(otherSpec GuitarSpec) bool {
	isGSRefundable, e := gs.GetIsRefundableVal()
	isOtherSpecRefundable, err := otherSpec.GetIsRefundableVal()
	if err != nil && e != nil {
		if gs.GetBackWood() == otherSpec.GetBackWood() || gs.GetBuilder() == otherSpec.GetBuilder() ||
			gs.GetFrontWood() == otherSpec.GetFrontWood() || gs.GetModel() == otherSpec.GetModel() || gs.GetNumberOfStrings() == otherSpec.GetNumberOfStrings() ||
			gs.GetTypeOfGuitar() == otherSpec.GetTypeOfGuitar() {
			return true
		}
	} else {
		if gs.GetBackWood() == otherSpec.GetBackWood() || gs.GetBuilder() == otherSpec.GetBuilder() ||
			gs.GetFrontWood() == otherSpec.GetFrontWood() || gs.GetModel() == otherSpec.GetModel() || gs.GetNumberOfStrings() == otherSpec.GetNumberOfStrings() ||
			gs.GetTypeOfGuitar() == otherSpec.GetTypeOfGuitar() || isGSRefundable == isOtherSpecRefundable {
			return true
		}
	}

	return false
}
func (gs *GuitarSpec) SetRefundable(b *bool) {
	gs.isRefundable = b
	//gs.isRefundableSetForSearch = 1
}
func (gs GuitarSpec) GetIsRefundableVal() (bool, error) {
	if gs.isRefundable == nil {
		return false, fmt.Errorf("nil ptr")
	}
	return *gs.isRefundable, nil
}
