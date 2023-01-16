package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	// simple if statement
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// if short form
	// variable declared by the statement are only in scope until the end of the if
	// if you are using v in the last return you will get error
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 2, 20),
	)
}
