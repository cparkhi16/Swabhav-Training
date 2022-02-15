package guitar

type GuitarSpec struct {
	model        string
	builder      Builder
	guitarType   GuitarType
	backWood     Wood
	frontWood    Wood
	noOfStrings  uint8
	isRefundable bool
}

var isRefundableArgSearched bool = false

func NewGuitarSpec(model string, builder Builder, guitarType GuitarType, backWood Wood, frontWood Wood, noOfStrings uint8, isRefundable bool) GuitarSpec {
	return GuitarSpec{
		model:        model,
		builder:      builder,
		guitarType:   guitarType,
		backWood:     backWood,
		frontWood:    frontWood,
		noOfStrings:  noOfStrings,
		isRefundable: isRefundable,
	}
}

func (g *GuitarSpec) SetModel(newModel string) {
	g.model = newModel
}

func (g *GuitarSpec) SetBuilder(newBuilder Builder) {
	g.builder = newBuilder
}

func (g *GuitarSpec) SetGuitarType(newGuitarType GuitarType) {
	g.guitarType = newGuitarType
}

func (g *GuitarSpec) SetBackWood(newBackWood Wood) {
	g.backWood = newBackWood
}

func (g *GuitarSpec) SetFrontWood(newFrontWood Wood) {
	g.frontWood = newFrontWood
}

func (g *GuitarSpec) SetNoOfStrings(newNoOfStrings uint8) {
	g.noOfStrings = newNoOfStrings
}

func (g *GuitarSpec) SetIsRefundable(newIsRefundable bool) {
	isRefundableArgSearched = true
	g.isRefundable = newIsRefundable
}

func (gs GuitarSpec) Matches(otherSpec GuitarSpec) bool {
	if otherSpec.model == "" {
		otherSpec.model = gs.model
	}
	if otherSpec.builder == "" {
		otherSpec.builder = gs.builder
	}
	if otherSpec.guitarType == "" {
		otherSpec.guitarType = gs.guitarType
	}
	if otherSpec.backWood == "" {
		otherSpec.backWood = gs.backWood
	}
	if otherSpec.frontWood == "" {
		otherSpec.frontWood = gs.frontWood
	}
	if otherSpec.noOfStrings == 0 {
		otherSpec.noOfStrings = gs.noOfStrings
	}
	if !isRefundableArgSearched {
		otherSpec.isRefundable = gs.isRefundable
	}
	if otherSpec == gs {
		return true
	}
	return false
}
