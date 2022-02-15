package main

import (
	"fmt"
	"shoppingApp/customer"
	"shoppingApp/item"
	"shoppingApp/store"
	"sync"
)

func main() {
	var m sync.Mutex
	var wg sync.WaitGroup
	//items
	bathtub := item.New(1, "bathing", 2000, 3)
	rubberDuck := item.New(2, "toy", 400, 3)
	flowerPot := item.New(3, "pot", 500, 3)
	//store
	onestop := store.New("onestop")
	onestop.AddItemToStore(bathtub)
	onestop.AddItemToStore(rubberDuck)
	onestop.AddItemToStore(flowerPot)

	shan := customer.New(23, "shan", 10000)
	sanju := customer.New(45, "sanju", 2000)

	shan.AddToCart(item.New(1, "bathing", 2000, 1))
	shan.AddToCart(item.New(3, "pot", 500, 1))
	sanju.AddToCart(item.New(3, "pot", 500, 3))

	fmt.Println("shan item list-->")
	shan.DisplayItemList()

	fmt.Println("sanju item list-->")
	sanju.DisplayItemList()

	fmt.Println("onestop item list-->")
	onestop.DisplayItemList()

	wg.Add(2)
	go shan.PlaceOrder(onestop, &m, &wg)
	go sanju.PlaceOrder(onestop, &m, &wg)
	//time.Sleep(1 * time.Second)
	wg.Wait()
	fmt.Println("shan balance-->", shan.GetBalance())
	fmt.Println("sanju balance-->", sanju.GetBalance())

	//time.Sleep(1 * time.Second)

	fmt.Println("onestop item list-->")
	onestop.DisplayItemList()
	

}
