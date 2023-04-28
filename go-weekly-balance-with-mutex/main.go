package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var balanceMutex sync.Mutex

type Income struct {
	Source string
	Amount int
}

func main() {
	var bankBalance int

	incomes := []Income{
		{Source: "Main job", Amount: 500},
		{Source: "Part time job", Amount: 50},
		{Source: "Investments", Amount: 100},
		{Source: "Gifts", Amount: 10},
	}

	wg.Add(len(incomes))
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balanceMutex.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp

				fmt.Printf("On week %d, earned $%d.00 from %s\n", week, income.Amount, income.Source)
				balanceMutex.Unlock()
			}
		}(i, income)
	}
	wg.Wait()

	fmt.Printf("Final bank balance: $%d.00", bankBalance)
}
