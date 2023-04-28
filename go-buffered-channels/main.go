package main

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {
	for {
		i := <-ch
		fmt.Println("Got", i, "from channel")

		time.Sleep(1 * time.Second)
	}
}

func main() {
	// buffered channel
	// first 10 executes so fast
	// then channel will take one by one
	ch := make(chan int, 10)

	go listenToChan(ch)

	for i := 0; i < 100; i++ {
		fmt.Println("sending", i, "to channel...")
		ch <- i
		fmt.Println("sent", i, "to channel!")
	}

	fmt.Println("Done")
	close(ch)
}
