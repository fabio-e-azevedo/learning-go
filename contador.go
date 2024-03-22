package main

import (
	"fmt"
	"time"
)

func contador(count int) {
	for i := 1; i <= count; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}

func main() {
	go contador(10)
	go contador(10)
	contador(10)
}
