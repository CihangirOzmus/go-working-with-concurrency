package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	wg.Add(3)
	go updateMessage("Hello, universe!")
	go updateMessage("Hello, cosmos!")
	go updateMessage("Hello, galaxy!")
	wg.Wait()

	fmt.Println(msg)
}

func updateMessage(s string) {
	defer wg.Done()
	mutex.Lock()
	msg = s
	mutex.Unlock()
}
