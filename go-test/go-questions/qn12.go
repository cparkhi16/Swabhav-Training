package main

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}

	for _, num := range nums {
		// You can add code but can't remove to get the expected output
		// Expected output should be 1 2 3 4 5 6
		wg.Add(1)
		go func() { // go func(){test(num)}() with wg.Wait() inside op -> 1,2,3,4,5,6 if wg.Wait() outside then random 5,6,6,6,6,6 because in that case num value is getting updated till test is getting called
			test(num)
		}()
	}
	wg.Wait()
	
}

func test(i int) {
	//wg.Add(1)
	fmt.Printf("%d ", i)
	wg.Done()
}
