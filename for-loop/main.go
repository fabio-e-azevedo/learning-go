package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// done := make(chan bool)

	// values := []string{"a", "b", "c"}
	// for _, v := range values {
	// 	go func() {
	// 		fmt.Println(v)
	// 		done <- true
	// 	}()
	// }

	// // wait for all goroutines to complete before exiting
	// for _ = range values {
	// 	<-done
	// }

	// c := 1
	// for range 15 {
	// 	fmt.Printf("%s\n", strings.Repeat("=", c))
	// 	c++
	// }

	for range 6 {
		fmt.Printf("%d, ", rand.N(61))
	}
	fmt.Printf("\n")
}
