package main

import (
	"fmt"
	"shop/flower"
	"shop/flowerShop"
)

func main() {
	lily := flower.NewFlower("lily", flower.Yellow, 30)
	rose := flower.NewFlower("rose", flower.Red, 100)
	rwsShop := flowerShop.NewFlowerShop(4)
	rwsShop.SetStrategy()
	value := rwsShop.GetPrice(lily)
	value2 := rwsShop.GetPrice(rose)
	fmt.Println(value, value2)
}
