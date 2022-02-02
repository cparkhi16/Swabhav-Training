package main

import (
	"app/aes"
	"app/model"
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var files []model.File

func employeeCSV() {
	empData := [][]string{
		{"Name", "City", "Skills"},
		{"Smith", "Newyork", "Java"},
		{"William", "Paris", "Golang"},
		{"Rose", "London", "PHP"},
	}

	csvFile, err := os.Create("employee.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)

	for _, empRow := range empData {
		_ = csvwriter.Write(empRow)
	}
	csvwriter.Flush()
	csvFile.Close()
}
func itemsCSV() {
	empData := [][]string{
		{"Name", "Price"},
		{"A", "100"},
		{"B", "200"},
		{"C", "300"},
	}

	csvFile, err := os.Create("items.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)

	for _, empRow := range empData {
		_ = csvwriter.Write(empRow)
	}
	csvwriter.Flush()
	csvFile.Close()
}
func GetEmployeeData() (string, string, string) {
	var Name string
	var City string
	var Skills string
	fmt.Println("Enter name")
	fmt.Scan(&Name)
	fmt.Println("Enter city")
	fmt.Scan(&City)
	fmt.Println("Enter skills")
	fmt.Scan(&Skills)
	return Name, City, Skills
}
func GetItemData() (string, string) {
	var Name string
	var Price string
	fmt.Println("Enter name")
	fmt.Scan(&Name)
	fmt.Println("Enter price")
	fmt.Scan(&Price)
	return Name, Price
}
func WriteToCSVFile(fileName string) {
	var data interface{}
	if fileName == "employee.csv" {
		Name, City, Skills := GetEmployeeData()
		//encryptedName := aes.Encryption([]byte(Name), "chinmay")
		//fmt.Println("Decrypting name ", string(aes.Decryption([]byte(encryptedName), "chinmay")))
		data = model.NewEmployee(Name, City, Skills)
	} else if fileName == "items.csv" {
		name, price := GetItemData()
		data = model.NewItem(name, price)
	}
	csvfile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer csvfile.Close()

	writer := csv.NewWriter(csvfile)
	defer writer.Flush()
	records := [][]string{{}}
	switch data.(type) {
	case *model.Employee:
		emp := data.(*model.Employee)
		records = [][]string{{emp.Name, emp.City, emp.Skills}}
	case *model.Item:
		item := data.(*model.Item)
		records = [][]string{{item.Name, item.Price}}
	}
	for _, record := range records {
		err := writer.Write(record)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
}

func AddFiles(file model.File) {
	files = append(files, file)
}

func CheckWriteAccessForFiles(user *model.User) {
	userBIBALevel, _ := strconv.Atoi(user.LevelBIBA)
	userBellLevel, _ := strconv.Atoi(user.LevelBELL)
	for _, val := range files {
		//fmt.Println(val)
		fileBellLevel, _ := strconv.Atoi(val.LevelBell)
		fileBIBALevel, _ := strconv.Atoi(val.LevelBIBA)
		if userBellLevel > fileBellLevel || userBIBALevel < fileBIBALevel {
			fmt.Println("You do not have write permissions for  ", val.FileName)
		} else {
			fmt.Println("Do you want to write to the file ", val.FileName)
			var response string
			fmt.Scan(&response)
			if response == "Y" {
				WriteToCSVFile(val.FileName)
			}
		}
	}
	// } else if mode == "BIBA" {
	// 	userBIBALevel, _ := strconv.Atoi(user.LevelBIBA)
	// 	for _, val := range files {
	// 		//fmt.Println(val)
	// 		fileLevel, _ := strconv.Atoi(val.Level)
	// 		if userBIBALevel < fileLevel {
	// 			fmt.Println("You do not have write permissions for  ", val.FileName)
	// 		} else {
	// 			fmt.Println("Do you want to write to the file ", val.FileName)
	// 			var response string
	// 			fmt.Scan(&response)
	// 			//fmt.Println("My response ", response)
	// 			if response == "Y" {
	// 				WriteToCSVFile(val.FileName)
	// 			}
	// 		}
	// 	}
	// }
}
func ReadCSVFile(fileName string) {
	// f, err := os.Open(fileName)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // remember to close the file at the end of the program
	// defer f.Close()

	// // read csv values using csv.Reader
	// csvReader := csv.NewReader(f)
	// for {
	// 	rec, err := csvReader.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	// do something with read line
	// 	fmt.Printf("%+v\n", rec)
	// }
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	// for _, eachline := range txtlines {
	// 	fmt.Println(eachline)
	// }
	for _, val := range txtlines {
		//fmt.Println(val[1])
		tmp := strings.Split(val, ",")
		fmt.Println(tmp)
	}
}
func CheckReadAccessForFiles(user *model.User) {
	userBellLevel, _ := strconv.Atoi(user.LevelBELL)
	userBIBALevel, _ := strconv.Atoi(user.LevelBIBA)
	for _, val := range files {
		//fmt.Println(val)
		fileBellLevel, _ := strconv.Atoi(val.LevelBell)
		fileBIBALevel, _ := strconv.Atoi(val.LevelBIBA)
		if userBellLevel < fileBellLevel || userBIBALevel > fileBIBALevel {
			fmt.Println("You do not have read permissions for  ", val.FileName)
		} else {
			fmt.Println("Contents from ", val.FileName, "----")
			ReadCSVFile(val.FileName)
		}
	}
	// } else if mode == "BIBA" {

	// 	for _, val := range files {
	// 		//fmt.Println(val)
	// 		fileLevel, _ := strconv.Atoi(val.Level)
	// 		if userBIBALevel > fileLevel {
	// 			fmt.Println("You do not have read permissions for  ", val.FileName)
	// 		} else {
	// 			fmt.Println("Contents from ", val.FileName, "----")
	// 			ReadCSVFile(val.FileName)
	// 		}
	// 	}
	// }
}
func WriteFile(filename string, fileMode os.FileMode, data string) error {
	file, err := os.OpenFile(filename, int(fileMode), 0777)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.Write([]byte(aes.Encryption([]byte(data), "hello"))); err != nil {
		return err
	}
	return nil
}

func ReadFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	fmt.Println("Data here ", data)
	//fmt.Println("---", aes.Decryption(data, "hello"))
	if err != nil {
		return "", err
	}
	return string(aes.Decrypt(data, "hello")), nil
}
func main() {
	file := model.NewFile("employee.csv", "2", "3") // no write // BELL
	AddFiles(*file)
	fileTwo := model.NewFile("items.csv", "3", "2") // no read //BELL
	AddFiles(*fileTwo)
	fmt.Println("1.Register 2.Login ")
	var response string
	fmt.Scanln(&response)
	if response == "1" {
		fmt.Println("Enter name password biba and bell levels")
		var name string
		var password string
		var biba string
		var bell string
		fmt.Scanf("%s %s %s %s", &name, &password, &biba, &bell)
		user := model.NewUser(name, password, biba, bell)
		user.AddUser()
		CheckWriteAccessForFiles(user)
		CheckReadAccessForFiles(user)
	} else {
		//user := model.NewUser("Chinmay", "test", "3", "3")
		user := model.Login()
		if user == nil {
			fmt.Println("Invalid username or password ")
		} else {
			CheckWriteAccessForFiles(user)
			CheckReadAccessForFiles(user)
		}
	}
	// employeeCSV()
	//itemsCSV()
	// user := model.NewUser("Chinmay", "test", "3", "3")
	// CheckWriteAccessForFiles(user)
	// CheckReadAccessForFiles(user)
	// var data string
	// fmt.Println("Enter data")
	// fmt.Scan(&data)
	// err := WriteFile("test.txt", fs.FileMode(os.O_APPEND), data)
	// if err != nil {
	// 	fmt.Println("-= Error while writing to file -=", err)
	// }
	// d, err := ReadFile("test.txt")
	// if err != nil {
	// 	fmt.Println("-= Error reading file -=", err)
	// }
	// fmt.Println("Contents of  file ", d)
}
