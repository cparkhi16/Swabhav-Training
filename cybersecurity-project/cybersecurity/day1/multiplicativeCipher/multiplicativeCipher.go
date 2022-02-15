package main

import (
	"fmt"
	"math"
	"strings"
)

var myarr []float64
var key int = 3

func main() {
	str := "yogesh"
	result := strings.ReplaceAll(str, " ", "")
	println(result)
	encrypt(result)
	decrypt(myarr)
}
func encrypt(str string) {
	for _, char := range str {
		fmt.Println(int(char))
		myarr = append(myarr, math.Round(math.Mod(float64((int(char)-97)*key), 26)))
	}
	fmt.Println(myarr)
}
func decrypt(myarr []float64) {
	var decr []rune
	for _, i := range myarr {
		fmt.Println(math.Round(math.Mod(i*float64(findInverse(key)), 26)) + 97)
		decr = append(decr, rune(math.Round(math.Mod(i*float64(findInverse(key)), 26))+97))
	}
	for _, i := range decr {
		fmt.Printf("Character corresponding to Ascii Code= %c\n", i)
	}

	fmt.Println(decr)
}

func findInverse(num int) int {
	for i := 0; i <= 26; i++ {
		if math.Mod(float64(num*i), 26) == 1 {
			return i
		}
	}
	return num
}
