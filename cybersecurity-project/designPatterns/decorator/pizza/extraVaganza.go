package pizza

type extraVaganza struct {
	price uint32
}

func NewExtraVaganza(price uint32) *extraVaganza {
	return &extraVaganza{
		price: price,
	}
}

func (e extraVaganza) GetPrice() uint32 {
	return e.price
}
