package main

import (
	"fmt"
	"test/duck"
)

func main() {
	d := duck.New("anni")
	fmt.Println(d.GetName())
}
