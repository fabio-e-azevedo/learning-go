package main

import (
	"fmt"
)

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func consume(ch chan int) {
	for x := range ch {
		fmt.Println("Mensagem processada: ", x)
	}
}

func main() {
  fmt.Println("Test Channel")
	ch := make(chan int)
	go publish(ch)
	consume(ch)
}
