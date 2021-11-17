package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func thirdway() {
	names := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter name: ")

		scanner.Scan()

		text := scanner.Text()
		fmt.Println("Enter age ")
		scanner.Scan()
		age := scanner.Text()
		fmt.Println("Type of age ", reflect.TypeOf(age))
		nAge, _ := strconv.Atoi(age)
		fmt.Println("After conv type of age ", nAge)
		if len(text) != 0 {

			fmt.Println(text)
			names = append(names, text)
		} else {
			break
		}
	}

	fmt.Println(names)

}
func secondway() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")

	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello %s\n", name)
	fmt.Println("Enter your age")
	n, _ := reader.ReadString('\n')
	fmt.Println("Type of age before conv ", reflect.TypeOf(n))
	val, _ := strconv.Atoi(n)
	fmt.Println("Type of age after conv ", reflect.TypeOf(val))
}
func firstway() {
	fmt.Println("Enter Your First Name: ")
	var first string
	fmt.Scanln(&first)
	fmt.Println("Enter Second Last Name: ")
	var second string
	fmt.Scanln(&second)
	fmt.Print("Your Full Name is: ")
	fmt.Print(first + " " + second)
	fmt.Println("Enter int")
	var i int
	fmt.Scanln(&i)
	fmt.Println("Val of i", i, "Type of i", reflect.TypeOf(i))
}
func fourthway() {
	fmt.Println("Enter name and age")
	var name string
	var age int
	fmt.Scan(&name, &age)
	fmt.Println("type of name ", reflect.TypeOf(name))
	fmt.Println("Type of age", reflect.TypeOf(age))
}
func main() {
	//firstway()
	//secondway()
	//thirdway()
	fourthway()
}
