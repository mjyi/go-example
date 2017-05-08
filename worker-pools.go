package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 0; w < 4; w++ {
		go worker(w, jobs, results)
	}

	for j := 0; j < 10; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 0; a < 10; a++ {
		<-results
	}
}
