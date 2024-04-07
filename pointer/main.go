package main

import (
	"fmt"
)

func returnPointer() *int {
	v := 123
	return &v
}

func incr(p *int) int {
	*p++ // incrementa o valor para o qual p aponta
	return *p
}

func main() {
	x := 1
	p := &x         // p contém o endereço de x
	fmt.Println(*p) // output: 1
	*p = 2          // equivalente a x = 2
	fmt.Println(x)  // output: 2
	var i = returnPointer()
	fmt.Println(*i)
	count := 0
	incr(&count) // 1
	incr(&count) // 2
	fmt.Println(incr(&count))
}
