package main

import (
	"fmt"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) {
	for {
		// except data from channel
		s := <-ping
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type someting and press Enter (enter Q to quit)")

	for {
		fmt.Print("-> ")
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			break
		}

		ping <- userInput
		response := <-pong

		fmt.Println("Response:", response)
	}

	fmt.Println("Closing channels")
	close(ping)
	close(pong)
}
