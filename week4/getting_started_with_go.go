package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func F(x float64) float64 {
	return math.Sqrt(1 - x*x)
}

func Quad(left, right, fleft, fright, lrarea, EPSILON float64) float64 {
	mid := (left + right) / 2
	fmid := F(mid)
	larea := (fleft + fmid) * (mid - left) / 2
	rarea := (fmid + fright) * (right - mid) / 2
	if math.Abs((larea+rarea)-lrarea) > EPSILON {
		larea = Quad(left, mid, fleft, fmid, larea, EPSILON)
		rarea = Quad(mid, right, fmid, fright, rarea, EPSILON)
	}
	return larea + rarea
}

func main() {
	EPSILON, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Usage ./getting_started_with_go.exe a b EPSILON")
		return
	}

	a, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("Usage ./getting_started_with_go.exe a b EPSILON")
		return
	}

	b, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Usage ./getting_started_with_go.exe a b EPSILON")
		return
	}
	start := time.Now()
	area := Quad(a, b, F(a), F(b), (F(a)+F(b))*(b-a)/2, EPSILON)
	elapsed := time.Since(start)
	fmt.Println("Area is:", area*4)
	fmt.Println("Time:", elapsed)

	// My testing on this showed that when EPSILON was decreased the accuracy also
	// decreased, b had little impact on my output but a had much impact.
}
