package main

import (
	"fmt"
	shape "myModuleShape/myPackageShape"
)

func main() {
	s := shape.NewShape("Triangle", 35)
	fmt.Println(s.GetName())
	fmt.Println(s.GetSize())
}
