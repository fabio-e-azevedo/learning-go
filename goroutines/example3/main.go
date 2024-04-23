package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	for i := 10; i >= 0; i-- {
		go func(i int) {
			time.Sleep(time.Duration(i) * 10 * time.Millisecond)
			if i == 10 {
				time.Sleep(100 * time.Millisecond)
			}
			fmt.Println("Finished:", i)
		}(i)
	}

	time.Sleep(3 * time.Second)
}
