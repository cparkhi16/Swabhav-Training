package main

import (
	"fmt"
	"math"
)

func main() {
	p := 9.0
	q := 23.0

	alicePrivateKey := 4.0
	X := math.Mod(math.Pow(p, alicePrivateKey), q)
	fmt.Println("X ", X)

	bobPrivateKey := 3.0
	Y := math.Mod(math.Pow(p, bobPrivateKey), q)
	fmt.Println("Y ", Y)

	aliceKey := math.Mod(math.Pow(Y, alicePrivateKey), q)
	bobKey := math.Mod(math.Pow(X, bobPrivateKey), q)
	fmt.Println(aliceKey, bobKey)
}
