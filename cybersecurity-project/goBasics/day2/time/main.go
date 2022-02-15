package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	currentTime := time.Now()
	time := currentTime.Format("15:04:05")
	time_hour, _ := strconv.Atoi(time[:2])
	time_min, _ := strconv.Atoi(time[3:5])
	fmt.Println(time_hour, time_min)
	//time_hour = 0
	if time_hour >= 8 && time_hour <= 11 {
		fmt.Println("Good Morning")
	} else if time_hour >= 12 && time_hour <= 15 {
		fmt.Println("Good Afternoon")
	} else if time_hour >= 16 && time_hour <= 20 {
		fmt.Println("Good Evening")
	} else {
		fmt.Println("Good Night")
	}

}
