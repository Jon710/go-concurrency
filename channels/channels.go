// channels are like fifo queues, used to communicate between goroutines
package main

import (
	"fmt"
	"time"
)

func worker(id int, msg chan int) {
	for res := range msg {
		fmt.Println("Worker: ", id, "received: ", res)
		time.Sleep(time.Second)
	}
}

// entry point of the program
func main() {
	msgChannel := make(chan int)

	// go forks a new goroutine from main (parent)
	go worker(1, msgChannel)
	go worker(2, msgChannel)

	for i := 0; i < 10; i++ {
		msgChannel <- i
	}
}

// go uses a fork-join model for concurrency (structure)
// unbuffered channels have a capacity of 1, hence it's synchronous
// buffered channels have a capacity of n, hence it's asynchronous
