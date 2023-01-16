package main

import "fmt"

// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
func lastInFirstOut() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func main() {
	// A defer statement defers the execution of a function until the surrounding function returns.
	//defer fmt.Println("World")
	//
	//fmt.Println("Hello")

	//example 2
	lastInFirstOut()
}
