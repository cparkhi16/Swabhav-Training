package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	const (
		_      = iota
		KB int = 1 << (10 * iota)
		MB
		GB
		TB
	)

	//take input
	var path string
	var sizeUnit string
	fmt.Print("Enter your file path & size unit: ")
	fmt.Scanf("%s %s", &path, &sizeUnit)

	//file size
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fileSize := float64(fileInfo.Size())
	switch sizeUnit {
	case "KB":
		fileSize = fileSize / float64(KB)
	case "MB":
		fileSize = fileSize / float64(MB)
	case "GB":
		fileSize = fileSize / float64(GB)
	case "TB":
		fileSize = fileSize / float64(TB)
	}
	fmt.Println("file size of-", path, " is ", fileSize, " ", sizeUnit)
}
