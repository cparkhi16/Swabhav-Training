package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Square struct {
	side uint
}

/*func (s *Square) Area() { //Violates SRP as Area() is doing two tasks of calculating area and displaying the area.. It should do one task at a time
	fmt.Println("Area of square is ", s.side*s.side)
}*/
type output struct{}

func (o output) JSON(s Square) string {
	res := struct {
		Area uint `json:"area"`
	}{
		Area: s.Area(),
	}

	bs, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	return string(bs)
}

func (s *Square) Area() uint {
	return s.side * s.side
}
func (o output) PrintTextArea(s Square) {
	fmt.Println("Area of sqaure is : ", s.Area())
}
func main() {
	s := Square{side: 4}
	o := output{}
	o.PrintTextArea(s)
	fmt.Println(o.JSON(s))
}
