package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

var serverAddress string = "localhost"

func aSec(t chan int) {
	timer := time.NewTimer(time.Second)

	<-timer.C
	t <- 0
}

func main() {
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	args := &Args{7, 3}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith␣error:", err)
	}
	fmt.Printf("Arith: %d∗%d=%d\n", args.A, args.B, reply)

	// Asynchronous call

	// We need to send a request several times because it is unreliable and the server
	// freezes (stuck in loop) send it five times with 1 second interval, we also need
	// to process response and check if any error accured, even errors that whould not
	// happen on our local machine, such as too much traffic / slow internet.
	// All this, and to add insult to injury, it is also very slow, even though this is
	// on localhost (same computer) it is much more slow than accessing directly by memory.
	// Btw this does not actually work as intended, when encoutering sertain errors.
	for i := 0; i < 5; i++ {
		t := make(chan int)
		go aSec(t)
		quotient := new(Quotient)
		divCall := client.Go("Arith.Divide", args, quotient, nil)
		replyCall := <-divCall.Done // will be equal to divCall
		// check errors, print, etc.
		if replyCall.Error != nil {
			log.Fatal("Encountered an error: ", replyCall.Error)
		}
		fmt.Printf("Arith: %d/%d=%d\n", args.A, args.B, quotient.Quo)
		fmt.Printf("Arith: %d%%%d=%d\n", args.A, args.B, quotient.Rem)
		break
	}

}
