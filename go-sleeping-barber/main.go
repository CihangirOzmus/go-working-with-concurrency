package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

var seatingCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	// randomize customer arrival
	rand.Seed(time.Now().UnixNano())

	color.Yellow("The Sleeping Barber Problem")
	color.Yellow("---------------------------")

	clientChan := make(chan string, seatingCapacity)
	doneChannel := make(chan bool)

	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan:     clientChan,
		BarbersDoneChan: doneChannel,
		Open:            true,
	}

	color.Green("The shop is open for the day!")

	shop.addBarber("Frank")
	shop.addBarber("Gerard")
	shop.addBarber("Milton")
	shop.addBarber("Susan")
	shop.addBarber("Kelly")

	// start the barber shop
	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.CloseShopForDay()
		closed <- true
	}()

	// add Clients
	i := 1
	go func() {
		for {
			// get a random number with average arrival rate
			randomMilliseconds := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration((randomMilliseconds))):
				shop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()

	// block until the barbershop is closed
	<-closed
}

func (shop *BarberShop) addClient(client string) {
	color.Green("*** %s arrives!", client)
	if shop.Open {
		select {
		case shop.ClientsChan <- client:
			color.Blue("%s takes a seat in the waiting room.", client)
		default:
			color.Red("The waiting room is full, so %s leaves.", client)
		}
	} else {
		color.Red("The shop is already closed, so %s leaves!", client)
	}
}
