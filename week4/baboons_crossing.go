package main

import (
	"fmt"
	"math/rand"
	"time"
)

type semaphore chan int

const maxRopeWeight = 50
const maxQueue = 5

var queueLength = 0
var weightOnRope = 0

// Baboon is a function
func Baboon(id int, mySide, otherSide, rope semaphore) {
	for {
		rope <- 0
		mySide <- 0

		// Set random male or female, weight 20 ro 10 respectively
		var s []int
		s = append(s, 10, 20)
		myWeight := s[rand.Intn(2)]
		queueLength++

		// Go if you can
		for {
			if myWeight+weightOnRope <= maxRopeWeight {
				weightOnRope += myWeight
				break
			}
		}

		// Don't let the other side enter the rope!
		if queueLength > 0 {
			otherSide <- 0
		}

		<-mySide
		<-rope

		// Crossing the canyon (CS)

		// Has crossed
		rope <- 0

		weightOnRope -= myWeight
		queueLength--

		if weightOnRope == 0 || queueLength >= maxQueue {
			queueLength = 0
			mySide <- 0
			<-otherSide
		}
		<-rope

		fmt.Println(id, "crossing with weight:", weightOnRope)
	}
}

func main() {
	b1, b2, r := make(semaphore, 1), make(semaphore, 1), make(semaphore, 1)

	fmt.Println("start")

	go Baboon(1, b1, b2, r)
	go Baboon(2, b2, b1, r)

	time.Sleep(2 * time.Second)
}
