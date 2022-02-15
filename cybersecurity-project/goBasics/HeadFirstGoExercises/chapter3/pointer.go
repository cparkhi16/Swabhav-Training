//Exercise 3 pg no.170
package main

import "fmt"

func negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
}

func main() {
	truth := true
	negate(&truth)
	fmt.Println(truth)
	lies := false
	negate(&lies)
	fmt.Println(lies)
}
