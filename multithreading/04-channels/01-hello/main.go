package main

import "fmt"

// Thread 1
func main() {
	canal := make(chan string) // Canal Vazio

	// Thread 2
	go func() {
		canal <- "Ola Mundo!" // Canal cheio
	}()

	// Thread 1
	msg := <-canal // Canal esvazia
	fmt.Println(msg)
}
