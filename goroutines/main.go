package main

import (
	"fmt"
	"time"
)

func working(from string, number int) {
	for i := range number {
		fmt.Println(from, i+1)
		time.Sleep(time.Second)
	}
}

func main() {
	go working("1 - goroutine:", 10)
	go working("2 - goroutine:", 11)
	working("1 - direct:", 12)
}
