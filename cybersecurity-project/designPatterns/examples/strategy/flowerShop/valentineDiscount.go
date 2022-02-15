package flowerShop

import (
	"shop/flower"
)

type ValentineDiscount struct {
	flowerShop *flowerShop
}

func (n ValentineDiscount) GetPrice(f *flower.Flower) float64 {
	//if red flower then 50% discount and for yellow flower 75% discount
	if f.GetColor() == flower.Red {
		return 0.5 * f.GetPrice()
	} else if f.GetColor() == flower.Yellow {
		return 0.75 * f.GetPrice()
	}
	return f.GetPrice()
}
