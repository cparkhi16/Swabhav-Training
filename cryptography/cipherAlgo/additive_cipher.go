package cipherAlgo

import (
	"fmt"
)

func Caesar(r rune, shift int) rune {
	s := int(r) + shift
	if s > 'z' {
		return rune(s - 26)
	} else if s < 'a' {
		return rune(s + 26)
	}
	return rune(s)
}
func Decrypt(r rune, shift int) rune {
	//fmt.Println("rune val shift ", r, string(r), shift)
	plainText := int(r) - shift
	//fmt.Println("In decrypt", string(rune(plainText)))
	//fmt.Println(r, shift, plainText)
	if plainText > 'z' {
		fmt.Println(rune(plainText - 26))
		return rune(plainText - 26)
	} else if plainText < 'a' {
		fmt.Println(rune(plainText + 26))
		return rune(plainText + 26)
	}
	//fmt.Println(plainText)
	return rune(plainText)
}

// func main() {
// 	value := "chinmay parkhi"
// 	value = strings.ReplaceAll(value, " ", "")
// 	fmt.Println(value)
// 	result := strings.Map(func(r rune) rune {
// 		return caesar(r, 1)
// 	}, value)
// 	fmt.Println(value, result)
// 	decrypt := strings.Map(func(r rune) rune {
// 		return decrypt(r, 1)
// 	}, result)
// 	fmt.Println(result, decrypt)
// }
