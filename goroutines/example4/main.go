package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for j := range jobs {
		fmt.Println("Worker", id, "processing job", j)
	}
}

func main() {
	// Número máximo de goroutines
	maxWorkers := 5

	// Número de jobs
	numJobs := 20

	// Canal para enviar os jobs
	jobs := make(chan int, numJobs)

	// WaitGroup para esperar todas as goroutines terminarem
	var wg sync.WaitGroup

	// Iniciar as goroutines
	for i := 1; i <= maxWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Enviar os jobs para as goroutines
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	// Fechar o canal de jobs
	close(jobs)

	// Esperar todas as goroutines terminarem
	wg.Wait()
}
