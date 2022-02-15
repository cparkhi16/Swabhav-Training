package main

import (
	"diwali/gift"
	"fmt"
)

func main() {
	box1 := gift.NewAlmond("shan")
	fmt.Println(box1.WhatIsThis())
	box2 := gift.NewGlitterGiftWrapper(box1, "shan")
	fmt.Println(box2.WhatIsThis())
}
