// Solve this without running the code
func main() {
	s := make([]int, 0, 5)
	//s = append(s, 1, 2, 3, 4)
	s = []int{1, 2, 3, 4}
	a := append(s, 5)
	fmt.Println(a) // O/P :- 1 2 3 4 5
	b := append(s, 6)
	fmt.Println(b) // O/P :- 1 2 3 4 6
	fmt.Println(a) // O/P :- 1 2 3 4 6
	fmt.Println(s) // O/P :- 1 2 3 4

	// What is the expected output and why?
	//when we append 5 to s, it returns new address and assign it to var "a", so under the var name "s" we still have 1 2 3 4
	//and under the var name "a" we have 1 2 3 4 6 .... Now when we append 6 and assign it to var b, we are still making the change
	//at same address as it was for var "a" .....so var "b[0]" and var "a[0]" will have same address
	//hence now a and b are 1 2 3 4 6 ... and s it still 1 2 3 4

}
