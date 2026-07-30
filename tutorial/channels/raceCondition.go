package main

import (
	"fmt"
	"time"
)

var balance = 0 // Shared memory

func depositRace(amount int) {
	current := balance
	time.Sleep(time.Millisecond) // Simulating processing time
	balance = current + amount   // ⚠️ RACE CONDITION: Direct modification
}

func depositeSafe(amount int, ch chan int) {
	ch <- amount
}

func main() {
	// // Start two workers modifying the same variable at once
	// go depositRace(100)
	// go depositRace(200)

	// time.Sleep(time.Second) // Wait for them to finish
	// fmt.Printf("Final Balance: %d\n", balance)

	var channel chan int = make(chan int, 2) // channel will only hold 2 data items

	go depositeSafe(100, channel)
	go depositeSafe(200, channel)
	balance += <-channel
	balance += <-channel

	fmt.Printf("Final Balance: %d\n", balance)
}
