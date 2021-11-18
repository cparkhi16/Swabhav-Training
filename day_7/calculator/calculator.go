package calculator

func Add(n1, n2 int) int {
	return n1 + n2
}

func Subtract(n1, n2 int) int {
	return n1 - n2
}
func SquareRoot(number float64) float64 {

	var sr float64 = number / 2
	var temp float64
	for {
		temp = sr
		sr = (temp + (number / temp)) / 2
		if (temp - sr) == 0 {
			break
		}
	}
	return sr

}
func ipow(base int, exp int) int {
	var result int = 1
	for {
		c := exp & 1
		//fmt.Println("C value ", c)
		if c == 1 {
			result *= base
		}
		if c == 0 {
			break
		}
		exp >>= 1
		//fmt.Println("Exp value", exp)
		base *= base
		//fmt.Println(base)
	}

	return result
}
