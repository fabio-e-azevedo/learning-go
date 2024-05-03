package main

import "fmt"

func main() {
	fmt.Println("Dead Duck")
}

// init execute before function main
func init() {
	fmt.Printf("Hello ")
}
