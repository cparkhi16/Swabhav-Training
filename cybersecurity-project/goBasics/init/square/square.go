package square

import "fmt"

type Square struct {
	Side int
}

var check int

func init() {
	fmt.Println("init called")
	check = 33
}

func (s Square) Area() int {
	return check
}
