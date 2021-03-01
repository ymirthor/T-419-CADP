package main

import "fmt"

var bobs_flag [2]bool
var alices_flag [2]bool

func Bob() {
	for {
		// Remainder of code
		bobs_flag[0] = true
		bobs_flag[1] = !alices_flag[1]
		loop := true
		local := bobs_flag[1]
		for loop {
			if !alices_flag[0] {
				loop = false
			}
			if local == alices_flag[1] {
				loop = false
			}
		}
		fmt.Println("Bob is entering critical section")
		// Critical section
		fmt.Println("Bob has exited critical section")
		bobs_flag[0] = false
	}
}

func Alice() {
	for {
		// Remainder of code
		alices_flag[0] = true
		alices_flag[1] = bobs_flag[1]
		loop := true
		local := alices_flag[1]
		for loop {
			if !bobs_flag[0] {
				loop = false
			}
			if local != bobs_flag[1] {
				loop = false
			}
		}
		fmt.Println("Alice is entering critical section")
		// Critical section
		fmt.Println("Alice has exited critical section")
		alices_flag[0] = false
	}
}

func main() {
	go Alice()
	Bob()
}
