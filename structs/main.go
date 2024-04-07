package main

import (
    "fmt"
)

type Shape struct {
	name string
	size int
}

func NewShape(name string, size int) Shape {
	return Shape{
		name: name,
		size: size,
	}
}

func main() {
	s := NewShape("Triangle", 35)

	fmt.Printf("Name: %s ===>>> Size: %d\n", s.name, s.size)
}
