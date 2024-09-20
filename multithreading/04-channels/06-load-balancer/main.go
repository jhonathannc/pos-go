package main

import (
	"fmt"
)

func worker(workerid int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerid, x)
		// time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)

	go worker(1, data)
	go worker(2, data)

	for i := 0; i < 10; i++ {
		data <- i
	}
}
