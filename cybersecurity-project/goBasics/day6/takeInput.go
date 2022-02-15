package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//1
	var name string
	var age int
	fmt.Print("Enter your name & age: ")
	fmt.Scanf("%s %d", &name, &age)
	fmt.Println("hi")
	//2
	fmt.Printf("Enter Your Address: ")
	var address string
	if _, err := fmt.Scanln(&address); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("hi")

	//3
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your college name: ")
	collegeName, _ := reader.ReadString('\n')

	//4
	fmt.Println("Enter your school name")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println(scanner.Text())

	if scanner.Err() != nil {
		panic("error")
	}

	//print data
	fmt.Printf("Name-%s Age-%d\n Address-%s College name-%s", name, age, address, collegeName)

}
