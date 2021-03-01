package main

import (
	"fmt"
)

var n int = 0

func P() {
	var temp int
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		temp = n
		n = temp + 1
	}
}

func main() {
	P()
	P()
	
	fmt.Println(n)
}
