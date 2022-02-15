package rectangle

type rectangle struct {
	length  float64
	breadth float64
}

func New(length float64, breadth float64) *rectangle {
	return &rectangle{
		length:  length,
		breadth: breadth,
	}
}

func (r *rectangle) GetLength() float64 {
	return r.length
}

func (r *rectangle) SetLength(newLength float64) {
	r.length = newLength
}

func (r *rectangle) GetBreadth() float64 {
	return r.breadth
}

func (r *rectangle) SetBreadth(newBreadth float64) {
	r.breadth = newBreadth
}

func (r *rectangle) Area() float64 {
	return r.length * r.breadth
}

func (r *rectangle) Perimeter() float64 {
	return 2 * (r.length + r.breadth)
}
