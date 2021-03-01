package main

import "fmt"

var f [2]bool
var turn int

func P(id int) {
	for {
		// Remainder of code
		fmt.Println("P:", id, "doing its code")
		f[id] = true
		for turn == 1-id {
			fmt.Println("P:", id, "waiting")
			for f[1-id] {
			} // await !f[1-id]
			turn = id
		}
		fmt.Println("P:", id, "critical section")
		// Critical section
		f[id] = false
	}
}

func main() {
	fmt.Println("Starting main")
	go P(0)
	P(1)
	fmt.Println("Exiting main")
}
