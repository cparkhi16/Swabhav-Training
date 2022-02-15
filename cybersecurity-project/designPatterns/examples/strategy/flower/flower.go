package flower

type Color string

const (
	Red    Color = "red"
	Green  Color = "green"
	Blue   Color = "purple"
	Yellow Color = "yellow"
)

type Flower struct {
	name  string
	color Color
	price float64
}

func NewFlower(name string, color Color, price float64) *Flower {
	return &Flower{
		name:  name,
		color: color,
		price: price,
	}
}

func (f *Flower) GetColor() Color {
	return f.color
}

func (f *Flower) GetPrice() float64 {
	return f.price
}
