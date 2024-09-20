package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Response struct {
	ID  int64
	Msg string
}

func main() {
	c1 := make(chan Response)
	c2 := make(chan Response)
	var i int64 = 0

	go func() {
		for {
			atomic.AddInt64(&i, 1)
			res := Response{ID: i, Msg: "Hello from RabbitMQ"}
			c1 <- res
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			atomic.AddInt64(&i, 1)
			res := Response{ID: i, Msg: "Hello from Kafka"}
			c2 <- res
			time.Sleep(time.Second)
		}
	}()

	for {
		select {
		case msg := <-c1:
			fmt.Printf("Received from RabbitMQ -> id: %d, msg: %s\n", msg.ID, msg.Msg)
		case msg := <-c2:
			fmt.Printf("Received from Kafka -> id: %d, msg: %s\n", msg.ID, msg.Msg)
		case <-time.After(5 * time.Second):
			fmt.Println("timeout")
		}
	}
}
