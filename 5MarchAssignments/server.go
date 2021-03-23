package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}
type Quotient struct {
	Quo, Rem int
}
type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	// This is a intesional bug in the server to show when the server encounters
	// an endless loop, how is the client suppose to handle this??
	// A lot of error handling needed, checking for endless loops, other errors, etc.
	for args.B == 2 {
		quo.Quo = args.B + 1
	}
	if args.B == 0 {
		return errors.New("divide␣by␣zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen␣error:", e)
	}
	c := make(chan int)
	go http.Serve(l, nil)
	c <- 0
}
