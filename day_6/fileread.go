package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func UsingReadFile() {
	//should not be used to read large files
	content, err := ioutil.ReadFile("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}
func ReadFileLineByLine() {
	f, err := os.Open("C:/Users/chinmay.parkhi/OneDrive - Forcepoint/Desktop/Swabhav-Training/change.txt")

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
func ReadCsv() {
	//var records [][]string
	txt_file, err := os.Open("C:/Users/chinmay.parkhi/OneDrive - Forcepoint/Desktop/GOLANG/GOLANGv/user_data.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer txt_file.Close()

	r := csv.NewReader(txt_file)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(records)
}
func main() {

	//UsingReadFile()
	//ReadFileLineByLine()
	ReadCsv()
}
