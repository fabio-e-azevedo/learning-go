package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	yearDay := currentTime.YearDay()

	if yearDay == 256 {
		fmt.Println("Programmer Congratulations!!!")
	} else if yearDay < 256 {
		fmt.Printf("%d days left\n", 256-yearDay)
	} else {
		currentYear := currentTime.Year()
		var yearDaysTotal int
		if currentYear%4 == 0 && (currentYear%100 != 0 || currentYear%400 == 0) {
			yearDaysTotal = 366
		} else {
			yearDaysTotal = 365
		}
		fmt.Printf("%d days left\n", (yearDaysTotal-yearDay)+256)
	}
}
