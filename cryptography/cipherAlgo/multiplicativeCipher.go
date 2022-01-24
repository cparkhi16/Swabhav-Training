package cipherAlgo

import (
	"fmt"
)

//var alphaMapping map[string]int
var numMapping = map[int]string{
	1:  "a",
	2:  "b",
	3:  "c",
	4:  "d",
	5:  "e",
	6:  "f",
	7:  "g",
	8:  "h",
	9:  "i",
	10: "j",
	11: "k",
	12: "l",
	13: "m",
	14: "n",
	15: "o",
	16: "p",
	17: "q",
	18: "r",
	19: "s",
	20: "t",
	21: "u",
	22: "v",
	23: "w",
	24: "x",
	25: "y",
	26: "z",
}
var alphaMapping = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
}

// func main() {
// 	plainText := "chinmay parkhi"
// 	plainText = strings.ReplaceAll(plainText, " ", "")
// 	fmt.Println(plainText)
// 	key := 7
// 	alphaMapping = map[string]int{
// 		"a": 1,
// 		"b": 2,
// 		"c": 3,
// 		"d": 4,
// 		"e": 5,
// 		"f": 6,
// 		"g": 7,
// 		"h": 8,
// 		"i": 9,
// 		"j": 10,
// 		"k": 11,
// 		"l": 12,
// 		"m": 13,
// 		"n": 14,
// 		"o": 15,
// 		"p": 16,
// 		"q": 17,
// 		"r": 18,
// 		"s": 19,
// 		"t": 20,
// 		"u": 21,
// 		"v": 22,
// 		"w": 23,
// 		"x": 24,
// 		"y": 25,
// 		"z": 26,
// 	}
// 	numMapping = map[int]string{
// 		1:  "a",
// 		2:  "b",
// 		3:  "c",
// 		4:  "d",
// 		5:  "e",
// 		6:  "f",
// 		7:  "g",
// 		8:  "h",
// 		9:  "i",
// 		10: "j",
// 		11: "k",
// 		12: "l",
// 		13: "m",
// 		14: "n",
// 		15: "o",
// 		16: "p",
// 		17: "q",
// 		18: "r",
// 		19: "s",
// 		20: "t",
// 		21: "u",
// 		22: "v",
// 		23: "w",
// 		24: "x",
// 		25: "y",
// 		26: "z",
// 	}
// 	a := getInverse(key)
// 	fmt.Println("Key val ", a)
// 	cipherText := MultiplicativeEncryption(plainText, key)
// 	decryptedText := MultiplicativeDecryption(cipherText, key, a)
// 	fmt.Println("Text", plainText)
// 	fmt.Println("Cypher Text", cipherText)
// 	fmt.Println("Plain Text ", decryptedText)
// }

func GetInverse(key int) int {
	for i := 1; i < 26; i++ {
		if (key*i)%26 == 1 {
			return i
		}
	}
	return 0
}

func MultiplicativeEncryption(plainText string, key int) string {
	fmt.Println(plainText, key)
	var v int
	ct := ""
	for _, val := range plainText {
		//fmt.Println(val)
		//fmt.Println(alphaMapping[fmt.Sprintf("%c", val)])
		v = (alphaMapping[fmt.Sprintf("%c", val)] * key) % 26
		//fmt.Println(v)
		//fmt.Println(numMapping[v])
		ct += numMapping[v]
	}
	return ct
}

func MultiplicativeDecryption(cypherText string, key int, a int) string {
	//cypherText = "udktmgshgvydk"
	var v int
	pt := ""
	fmt.Println("Here cypher text in decrypt ", cypherText)
	for _, val := range cypherText {
		//fmt.Println("Here ", alphaMapping[fmt.Sprintf("%c", val)]*a)
		v = (alphaMapping[fmt.Sprintf("%c", val)] * a) % 26
		//fmt.Println(alphaMapping[fmt.Sprintf("%c", val)])
		pt += numMapping[v]
		// fmt.Println(v)
		// fmt.Println(numMapping[v])
		// fmt.Println(pt)
	}
	return pt
}
