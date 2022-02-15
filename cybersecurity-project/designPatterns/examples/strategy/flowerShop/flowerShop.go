package flowerShop

import (
	"shop/flower"
)

type DiscountType interface {
	GetPrice(f *flower.Flower) float64
}

type flowerShop struct {
	strategy     DiscountType
	currentMonth uint8
}

func NewFlowerShop(currentMonth uint8) *flowerShop {
	return &flowerShop{
		currentMonth: currentMonth,
	}
}

func (f *flowerShop) SetStrategy() {
	if f.currentMonth == 2 {
		f.strategy = ValentineDiscount{flowerShop: f}
	} else if f.currentMonth == 12 {
		f.strategy = ChristmasDiscount{flowerShop: f}
	} else {
		f.strategy = NoDiscount{flowerShop: f}
	}
}

func (f *flowerShop) GetPrice(flower *flower.Flower) float64 {
	return f.strategy.GetPrice(flower)
}
