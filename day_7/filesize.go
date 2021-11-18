package main

import (
	"fmt"
	"log"
	"os"
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
)

func DisplayFileSize(path string, fsize string) {
	//fmt.Println(fsize)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	fi, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	if fsize == "Bytes" {
		fmt.Println("Size of file in bytes", fi.Size())
	} else if fsize == "KB" {
		fmt.Println("Size of file in KB", fi.Size()/KB)
	} else if fsize == "MB" {
		fmt.Println("Size of file in MB", fi.Size()/MB)
	}

}
func main() {
	fmt.Print("Enter file path and your preferred file size (Bytes,KB,MB)")
	var path string
	var size string
	fmt.Scanf("%s %s", &path, &size)
	DisplayFileSize(path, size)

}
