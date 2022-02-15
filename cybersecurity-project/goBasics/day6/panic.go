package main

import "fmt"

func main() {
	var i int
	i = 10
	var j int
	j = 40
	fmt.Println("open resource")
	j = 50

	defer func(i int, j *int) {
		details := recover()
		fmt.Println("recover details-", details)
		closeResource(i, j) //30 100
	}(i, &j)
	defer f1()
	defer closeResource(i, &j) //10 100

	i = 30
	j = 100
	panic("panic here")
	fmt.Println("At the end of main")
}

func closeResource(i int, j *int) {
	fmt.Println("value of i and j from closeResource-", i, *j)
}

func f1() {
	fmt.Println("inside f1")
}
