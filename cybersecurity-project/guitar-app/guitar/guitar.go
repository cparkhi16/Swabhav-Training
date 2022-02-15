package guitar

type Builder string

const (
	Fender Builder = "fender"
	Martin Builder = "martin"
	Gibson Builder = "gibson"
)

type Wood string

const (
	Mahogony Wood = "mahogony"
	Maple    Wood = "maple"
	Cocobolo Wood = "cocobolo"
	Cedar    Wood = "cedar"
	Sitka    Wood = "stika"
)

type GuitarType string

const (
	Accoustic GuitarType = "accoustic"
	Electric  GuitarType = "electric"
)

type Guitar struct {
	serialNo string
	price    uint16
	spec     GuitarSpec
}

func NewGuitar(serialNo string, price uint16, spec GuitarSpec) *Guitar {
	return &Guitar{
		serialNo: serialNo,
		price:    price,
		spec:     spec,
	}
}

func (g *Guitar) GetSerialNo() string {
	return g.serialNo
}

func (g *Guitar) SetSerialNo(newSerialNo string) {
	g.serialNo = newSerialNo
}

func (g *Guitar) IsSerialNoEmpty() bool {
	if g.serialNo == "" {
		return true
	}
	return false
}

func (g *Guitar) GetPrice() uint16 {
	return g.price
}

func (g *Guitar) SetPrice(newPrice uint16) {
	g.price = newPrice
}

func (g *Guitar) SetSpec(newSpec GuitarSpec) {
	g.spec = newSpec
}

func (g *Guitar) IsPriceEmpty() bool {
	if g.price == 0 {
		return true
	}
	return false
}

func (g *Guitar) GetSpec() GuitarSpec {
	return g.spec
}
