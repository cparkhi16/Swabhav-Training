package flowerShop

import (
	"shop/flower"
)

type ChristmasDiscount struct {
	flowerShop *flowerShop
}

func (n ChristmasDiscount) GetPrice(f *flower.Flower) float64 {
	//if blue flower then 50% discount and for green flower 75% discount
	if f.GetColor() == flower.Blue {
		return 0.5 * f.GetPrice()
	} else if f.GetColor() == flower.Green {
		return 0.75 * f.GetPrice()
	}
	return f.GetPrice()
}
