package main

import (
	"azure/token"
	"fmt"
)

func main() {
	fmt.Print(token.Get())
}
