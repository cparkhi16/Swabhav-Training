package main

import "fmt"

func ExerciseOneChapterOne() {
	var price int = 100
	fmt.Println("Price is", price, "dollars.")

	var taxRate float64 = 0.08

	var tax float64 = float64(price) * taxRate
	fmt.Println("Tax is", tax, "dollars.")

	var total float64 = float64(price) + tax
	fmt.Println("Total cost is", total, "dollars.")

	var availableFunds int = 120.

	fmt.Println(availableFunds, "dollars available.")

	fmt.Println("Within budget?", total <= float64(availableFunds))

}
func ExerciseOneChapterTwo() {
	if true {

		fmt.Println("true")
	}
	if false {
		fmt.Println("false")
	}

	if !false {

		fmt.Println("false")
	}

	if true {
		fmt.Println("if true")
	} else {
		fmt.Println("else")
	}

	if false {

		fmt.Println("if false")
	} else if true {

		fmt.Println("else if true")
	}

	if 12 == 12 {

		fmt.Println("12 = 12")
	}
	if 12 != 12 {

		fmt.Println("12 != 12")
	}
	if 12 > 12 {

		fmt.Println("12 > 12")
	}
	if 12 >= 12 {
		fmt.Println("12 >= 12")
	}

	if 12 == 12 && 5.9 == 5.9 {
		fmt.Println("12==12 && 5.9==5.9")
	}
	if 12 == 12 && 5.9 == 6.4 {
		fmt.Println("12==12 && 5.9== 6.4")
	}

	if 12 == 12 || 5.9 == 6.4 {
		fmt.Println("12==12 || 5.9== 6.4")
	}

}
func main() {
	ExerciseOneChapterOne()
	ExerciseOneChapterTwo()
}
