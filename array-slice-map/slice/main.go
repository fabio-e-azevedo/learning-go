package main

import (
	"fmt"
	"reflect"
	"slices"
	"strings"
)

type Currency int

// enumerated type
const (
	USD Currency = iota + 1
	EUR
	GBP
	RMB
	BRL
)

func main() {
	// slice representam sequencias de tamanho variável cujos elementos tem o mesmo tipo
	// um tipo slice é escrito como "[]T", em que os elementos tem o tipo T
	// ele se parece com um tipo array mas sem o tamanho definido
	symbol := []string{USD: "$", EUR: "€", GBP: "£", RMB: "¥", BRL: "R$"}
	fmt.Printf("Index: %d, Value: %s\n", USD, symbol[USD])
	fmt.Printf("Index: %d, Value: %s\n", EUR, symbol[EUR])
	fmt.Printf("Index: %d, Value: %s\n", GBP, symbol[GBP])
	fmt.Printf("Index: %d, Value: %s\n", RMB, symbol[RMB])
	fmt.Printf("Index: %d, Value: %s\n", BRL, symbol[BRL])
	fmt.Printf("Type %s: %T\n", reflect.TypeOf(symbol).Kind(), symbol)
	fmt.Printf("Total Elements: %d\n", len(symbol))
	fmt.Printf("Last Element: %s\n", symbol[len(symbol)-1])
	fmt.Printf("Values: %s\n", symbol)
	fmt.Printf("%s\n\n", strings.Repeat("<>", 40))

	var numbers []int
	printSlice(numbers)

	numbers = append(numbers, 10)
	printSlice(numbers)

	numbers = append(numbers, 20)
	printSlice(numbers)

	numbers = append(numbers, 30, 40)
	printSlice(numbers)

	numbers = append(numbers, 50, 60)
	printSlice(numbers)

	a := []string{"Dead ", "Duck "}
	b := []string{"is ", "death"}
	c := slices.Concat(a, b)
	fmt.Println(c)

	// func make([]T, len, cap) []T
	mySlice1 := make([]int, 5, 10)
	printSlice(mySlice1)

	// the capacity argument is usually omitted and defaults to the length
	mySlice2 := make([]int, 5)
	printSlice(mySlice2)
}

func printSlice(s []int) {
	fmt.Printf("\n%v\n", strings.Repeat("=", 79))
	fmt.Printf("len=%d cap=%d values=%v\n", len(s), cap(s), s)
	fmt.Println(strings.Repeat("=", 79))
}
