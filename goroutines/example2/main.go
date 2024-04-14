package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	totalTime := time.Now()

	var wg sync.WaitGroup

	wg.Add(3)
	results := make(chan string, 3)
	go getUser(&wg, results)
	go getCard(&wg, results)
	go getBuy(&wg, results)
	wg.Wait()
	close(results)

	for result := range results {
		fmt.Println(result)
	}
	fmt.Println("Total Time: ", time.Since(totalTime))
}

func getUser(wg *sync.WaitGroup, result chan string) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 100)
	result <- "User: Dead Duck"
}

func getCard(wg *sync.WaitGroup, result chan string) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 150)
	result <- "Card: VISA 4556 7124 0953 6117"
}

func getBuy(wg *sync.WaitGroup, result chan string) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 153)
	result <- "Buy: Product XPTO"
}
