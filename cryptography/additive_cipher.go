package main

import (
	"fmt"
	"strings"
)

func caesar(r rune, shift int) rune {
	s := int(r) + shift
	if s > 'z' {
		return rune(s - 26)
	} else if s < 'a' {
		return rune(s + 26)
	}
	return rune(s)
}
func decrypt(r rune, shift int) rune {
	plainText := int(r) - shift
	if plainText > 'z' {
		return rune(plainText - 26)
	} else if plainText < 'a' {
		return rune(plainText + 26)
	}
	//fmt.Println(plainText)
	return rune(plainText)
}
func main() {
	value := "chinmay parkhi"
	value = strings.ReplaceAll(value, " ", "")
	fmt.Println(value)
	result := strings.Map(func(r rune) rune {
		return caesar(r, 1)
	}, value)
	fmt.Println(value, result)
	decrypt := strings.Map(func(r rune) rune {
		return decrypt(r, 1)
	}, result)
	fmt.Println(result, decrypt)
}
