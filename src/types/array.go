package main

import "fmt"

// The type [n]T is an array of n values of type T.

func main() {
	// declares a variable a as an array of ten integers
	// var a [10]int
	// array length cannot be resized

	//example 1
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])

	//example 2
	var x [5]int
	// By default, an array is zero-valued, which for ints means 0s.result emp: [0 0 0 0 0]
	fmt.Println("emp:", x)

	// example 3
	// Use this syntax to declare and initialize an array in one line
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	// The builtin len returns the length of an array
	fmt.Println("length:", len(primes))
}
