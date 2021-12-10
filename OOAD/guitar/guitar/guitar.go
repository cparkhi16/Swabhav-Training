package guitar

type Guitar struct {
	serialNumber string
	price        uint16
	spec         GuitarSpec
}

func NewGuitar(serialNumber string, price uint16, spec GuitarSpec) *Guitar {
	return &Guitar{serialNumber: serialNumber, price: price, spec: spec}
}
func (g Guitar) GetSerialNumber() string {
	return g.serialNumber
}
func (g Guitar) GetPrice() uint16 {
	return g.price
}
func (g Guitar) GetSpecs() *GuitarSpec {
	return &g.spec
}
