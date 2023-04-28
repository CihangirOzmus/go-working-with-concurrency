package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	wg.Add(len(words)) // how many GoRoutines are needed
	for i, word := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, word), &wg)
	}
	wg.Wait() // wait for all to be done

	println("GoRoutines are completed.")
}

func printSomething(word string, wg *sync.WaitGroup) {
	defer wg.Done() // decrements wait group
	fmt.Println(word)
}
