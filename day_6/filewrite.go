package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func WriteFileWithWriteString() {
	//if file exists then it will update the file with new contents or else if it doesn't exist it will create it
	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	fmt.Println("enter some text")
	var text string
	fmt.Scanln(&text)
	_, err2 := f.WriteString(text)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
}
func WriteIntoFileWithIOutil() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")

	name, _ := reader.ReadString('\n')
	data := []byte(name)

	//err := ioutil.WriteFile("datanew.txt", data, 0)
	err := ioutil.WriteFile("dnew.txt", data, 0777)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("done")
}
func readFileLineByLine(path string) {
	f, err := os.Open(path)
	//f, err := os.Open("C:/Users/chinmay.parkhi/OneDrive - Forcepoint/Desktop/Swabhav-Training/change.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		fmt.Println("Hi--------", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func testReadWrite() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")

	name, _ := reader.ReadString('\n')
	data := []byte(name)

	//err := ioutil.WriteFile("datanew.txt", data, 0)
	err := ioutil.WriteFile("new.txt", data, 0777)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Writing done, reading file with os")
	readFileLineByLine("new.txt")
}
func main() {
	//WriteFileWithWriteString()
	//WriteIntoFileWithIOutil()
	testReadWrite()
}
