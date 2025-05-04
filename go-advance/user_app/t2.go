package main

import "sync"

func main() {
	even := make(chan bool)
	odd := make(chan bool)
	go func() {
		defer close(odd)
		for i := 0; i <= 10; i += 2 {
			<-even
			print("Even ====>")
			println(i)
			odd <- true
		}
	}()
	var wait sync.WaitGroup
	wait.Add(1)
	go func() {
		for i := 1; i <= 10; i += 2 {
			_, ok := <-odd
			if !ok {
				wait.Done()
				return
			}
			print("Odd ====>")
			println(i)
			even <- true
		}
	}()
	even <- true
	wait.Wait()
}
