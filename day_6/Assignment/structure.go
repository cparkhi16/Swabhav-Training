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
			fmt.Println("===FILE WITHIN THE CURRENT DIRECTORY ====", file.Name())
		} else if file.IsDir() {
			var fname string = filep + "/"
			fmt.Println(" directory within a directory --------------------", file.Name())
			ReadDir(fname+file.Name()+"/", initialpath, true)
		}
	}
}
func main() {
	var initialpath string = "C:/Users/chinmay.parkhi/OneDrive - Forcepoint/Desktop/Swabhav-Training/"
	files, err := ioutil.ReadDir(initialpath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			fmt.Println("-----------File-------------", file.Name())
		} else if file.IsDir() {
			fmt.Println("---------------------Directory --------------------", file.Name())
			ReadDir(file.Name(), initialpath, false)
		}
	}
}
