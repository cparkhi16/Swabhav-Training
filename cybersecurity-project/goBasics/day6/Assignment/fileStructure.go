package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	path := "../test1"
	inside := 0
	err := readDir(path, inside)
	if err != nil {
		panic(err)
	}
}

func readDir(path string, inside int) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	names, _ := file.Readdirnames(0)
	for _, name := range names {
		filePath := fmt.Sprintf("%v/%v", path, name)
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()
		fileInfo, err := file.Stat()
		if err != nil {
			return err
		}

		//fmt.Println(filePath, name)
		for i := 0; i < inside; i++ {
			fmt.Printf("   ")
		}
		if fileInfo.IsDir() {
			//for dir
			_, err = file.Readdir(1)
			if err == io.EOF {
				//if dir is empty
				fmt.Printf("|%s\n", name)
			} else {
				fmt.Printf("|%s\n", name)
				inside++
				readDir(filePath, inside)
				inside--
			}
		} else {
			//for normal file

			fmt.Printf("|--%s\n", name)
		}
	}
	return nil
}
