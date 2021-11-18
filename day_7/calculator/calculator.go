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
func ipow(base int, exp int) (int, string) {
	output := 1
	if base < 0 && exp < 0 {
		return -1, "NaN"
	}
	for exp != 0 {
		output *= base
		exp -= 1
	}
	//fmt.Println("Output", output)
	return output, ""
}
