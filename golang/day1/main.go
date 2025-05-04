package main

import (
	add "day1/math"
	"fmt"
	"sync"
)

//https://medium.com/@mail2rajeevshukla/unlocking-the-power-of-goroutines-understanding-gos-lightweight-concurrency-model-3775f8e696b0
// type Message struct {
// 	chats   []string
// 	friends []string
// }

func main() {
	fmt.Println("Hello World!")
	ans := add.AddTwoNumbers(1, 2)
	fmt.Println(" val after add is ", ans)
	an := add.VariadicArg(1, 2, 3)
	fmt.Println(" val for varidaic func ", an)
	m := add.Closure(1)
	m(1)
	k := add.Sum()
	fmt.Println(" sum in closure is ", k(5))
	fmt.Println(" again after adding 5 ans in closure ", k(10))

	a := add.Switch(1)
	fmt.Println(" switch bool res is ", a)
	add.ForLoop(10)

	add.WhileLoop(9)

	answ := add.DoMathOps(1, 2, add.Addition)
	fmt.Println(" addition ans is ", answ)

	add.ArrayTest()
	add.SliceTest()
	add.ConvertArrToSlice()

	var i int = 17
	var p *int = &i
	add.PassByRef(p)
	fmt.Println(" after pass by ref ", i)

	add.MapTest()
	add.ComplexMap()
	add.DeepCopySlice()
	var wg sync.WaitGroup
	for k := 1; k <= 3; k++ {
		wg.Add(1)
		go add.Worker(k, &wg)
	}
	wg.Wait()
	fmt.Println(" Worker go routine test")

	wh := &sync.WaitGroup{}
	ch := make(chan *add.Message, 2)
	wh.Add(2)
	go add.GetFriendsChat(ch, wh)
	go add.GetFriendsList(ch, wh)
	wh.Wait()
	close(ch)
	for msg := range ch {
		fmt.Println(msg)
	}
	add.SelectDemo()
	add.Temp()
	//fmt.Println(" Friends List ", f)
	//fmt.Println("Chat list ", c)
}
