package main

import "fmt"

var bobs_flag [3]bool
var alices_flag [3]bool
var charlies_flag [3]bool

func Alice() {
	for {
		// Remainder of code
		alices_flag[0] = true
		alices_flag[1] = !charlies_flag[2]
		alices_flag[2] = bobs_flag[1]
		loop := true
		for loop {
			// Both of the other flags are up
			if bobs_flag[0] == true && charlies_flag[0] == true {
				if alices_flag[1] == charlies_flag[2] && alices_flag[2] != bobs_flag[1] {
					loop = false
				}
			}
			// Bobs flag is up
			if bobs_flag[0] == true && charlies_flag[0] == false {
				if alices_flag[2] != bobs_flag[1] {
					loop = false
				}
			}
			// Charlies flag is up
			if charlies_flag[0] == true && bobs_flag[0] == false {
				if alices_flag[1] == charlies_flag[2] {
					loop = false
				}
			}
			// No flags are up
			if charlies_flag[0] == false && bobs_flag[0] == false {
				loop = false
			}
		}
		fmt.Println("Alice is entering critical section")
		// Critical section
		fmt.Println("Alice has exited critical section")
		alices_flag[0] = false
	}
}

func Bob() {
	for {
		// Remainder of code
		bobs_flag[0] = true
		bobs_flag[1] = !alices_flag[2]
		bobs_flag[2] = charlies_flag[1]
		loop := true
		for loop {
			// Both of the other flags are up
			if charlies_flag[0] == true && alices_flag[0] == true {
				if bobs_flag[1] == alices_flag[2] && bobs_flag[2] != charlies_flag[1] {
					loop = false
				}
			}
			// Charlies flag is up
			if charlies_flag[0] == true && alices_flag[0] == false {
				if bobs_flag[2] != charlies_flag[1] {
					loop = false
				}
			}
			// Alices flag is up
			if alices_flag[0] == true && charlies_flag[0] == false {
				if bobs_flag[1] == alices_flag[2] {
					loop = false
				}
			}
			// No flags are up
			if bobs_flag[0] == false && alices_flag[0] == false {
				loop = false
			}
		}
		fmt.Println("Bob is entering critical section")
		// Critical section
		fmt.Println("Bob has exited critical section")
		bobs_flag[0] = false
	}
}

func Charlie() {
	for {
		// Remainder of code
		charlies_flag[0] = true
		charlies_flag[1] = !bobs_flag[2]
		charlies_flag[2] = alices_flag[1]
		loop := true
		for loop {
			// Both of the other flags are up
			if bobs_flag[0] == true && alices_flag[0] == true {
				if charlies_flag[1] == bobs_flag[2] && charlies_flag[2] != alices_flag[1] {
					loop = false
				}
			}
			// Bobs flag is up
			if bobs_flag[0] == true && alices_flag[0] == false {
				if charlies_flag[1] == bobs_flag[2] {
					loop = false
				}
			}
			// Alices flag is up
			if alices_flag[0] == true && bobs_flag[0] == false {
				if charlies_flag[2] != alices_flag[1] {
					loop = false
				}
			}
			// No flags are up
			if bobs_flag[0] == false && alices_flag[0] == false {
				loop = false
			}
		}
		fmt.Println("Charlie is entering critical section")
		// Critical section
		fmt.Println("Charlie has exited critical section")
		charlies_flag[0] = false
	}
}

func main() {
	go Alice()
	go Bob()
	Charlie()
}
