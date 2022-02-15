package flowerShop

import (
	"shop/flower"
)

type NoDiscount struct {
	flowerShop *flowerShop
}

func (n NoDiscount) GetPrice(f *flower.Flower) float64 {
	return f.GetPrice()
}
