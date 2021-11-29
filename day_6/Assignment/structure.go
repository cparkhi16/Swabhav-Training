package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
)

func ReadDir(file string, initialpath string, innerFile bool) {
	var filep string
	var files []fs.FileInfo
	var err error
	if innerFile {
		filep = file
	} else {
		filep = initialpath + file
	}
	files, err = ioutil.ReadDir(filep)
	if err != nil {
		log.Fatal("Invalid path", file, innerFile)
		return
	}
	for _, file := range files {
		if !file.IsDir() {
			if innerFile {
				fmt.Println("   |--", file.Name())
			} else {
				fmt.Println("  |--", file.Name())
			}
		} else if file.IsDir() {
			var fname string = filep + "/"
			if innerFile {
				fmt.Println("  |---", file.Name())
			} else {
				fmt.Println("|---", file.Name())
			}
			ReadDir(fname+file.Name()+"/", initialpath, true)
		}
	}
}
func main() {
	var initialpath string = "C:/Users/chinmay.parkhi/OneDrive - Forcepoint/Desktop/swabhav_training/go-basics/day_8/"
	files, err := ioutil.ReadDir(initialpath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			fmt.Println("|-", file.Name())
		} else if file.IsDir() {
			fmt.Println("|", file.Name())
			ReadDir(file.Name(), initialpath, false)
		}
	}
}
