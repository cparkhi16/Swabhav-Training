// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
)

func main() {
	//two prime numbers selected
	p := 4.0
	q := 17.0

	//sender
	senderPrivateKey := 3.0
	X := math.Mod(math.Pow(p, senderPrivateKey), q)

	//receiver
	receiverPrivateKey := 6.0
	Y := math.Mod(math.Pow(p, receiverPrivateKey), q)

	fmt.Println(X, Y)
	senderKey := math.Mod(math.Pow(Y, senderPrivateKey), q)
	receiverKey := math.Mod(math.Pow(X, receiverPrivateKey), q)

	fmt.Println(senderKey, receiverKey)

}
