package main

import (
	"fmt"
	"reflect"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
	BRL
)

func main() {
	symbol := []string{USD: "$", EUR: "€", GBP: "£", RMB: "¥", BRL: "R$"}
	fmt.Println(USD, symbol[USD])
	fmt.Println(EUR, symbol[EUR])
	fmt.Println(GBP, symbol[GBP])
	fmt.Println(RMB, symbol[RMB])
	fmt.Println(BRL, symbol[BRL])
	fmt.Printf("Type %s: %T\n", reflect.TypeOf(symbol).Kind(), symbol)
}
