// Solve this without running the code
func main() {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4)
	a := append(s, 5)
	fmt.Println(a) // O/P :- 1,2,3,4,5____________________OP 1
	b := append(s, 6)
	fmt.Println(b) // O/P :- 1,2,3,4,6____________________OP 2
	fmt.Println(a) // O/P :- 1,2,3,4,6____________________OP 3
	fmt.Println(s) // O/P :- 1,2,3,4____________________OP 4

	// What is the expected output and why?
	//FOR OP1 --> Since capacity of slice is 5 , and it contains 4 elements an integer (5) will be added and resulted is assigned to a new slice (a)
	//FOR OP2 ---> Since previous append has not actually changed s and also still it has a capacity to store one integer it will be assigned to a new slice (b)
	//FOR OP3 --> Now Since previous append operation had not actually changed s and still s had a capacity the integer (6) will be replaced with (5) in slice (a) as update is done on same address
	//FOR OP4 --> In the entire append operations the appended results are not assigned to s , so s will not be updated
}
