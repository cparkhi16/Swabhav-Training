package main

import "fmt"

func reverseString(s string) string {
	left, right := 0, len(s)-1
	r := []rune(s)
	for left < right {
		r[left], r[right] = r[right], r[left]
		left++
		right--
	}
	return string(r)
}
func main() {
	a := "hello"
	fmt.Println(reverseString(a))
}
