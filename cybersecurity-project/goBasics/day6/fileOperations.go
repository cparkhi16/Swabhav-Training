// Golang program to read and write the files
package main

// importing the packages
import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func createFile(fileName string) *os.File {
	_, err := os.Stat(fileName)
	if err == nil {
		fmt.Println(fileName, " File already exists")
		file, _ := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0777) //
		return file
	}

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	return file
}

func writeToFile(file *os.File, text string) {
	len, err := file.WriteString(text)

	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", len)
}

func readFromFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", fileName)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)

}

func writeToFile2(fileName string, text string) {
	err := ioutil.WriteFile(fileName, []byte(text), 0644)
	if err != nil {
		panic(err)
	}
}

func readFromFile2(fileName string) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nData: %s", data)
}

func readFromFile3(file *os.File, numberOfBytes int) {
	byteSlice := make([]byte, numberOfBytes)
	bytesRead, err := file.Read(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", bytesRead)
	log.Printf("Data read: %s\n", byteSlice)
}

func main() {
	fileName := "text.txt"
	file := createFile(fileName)
	defer file.Close()
	//writeToFile(file, "hello((")
	writeToFile(file, "hello")
	//readFromFile(fileName)
	//readFromFile2(fileName)
	readFromFile3(file, 1)
}
