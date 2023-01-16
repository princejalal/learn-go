package main

import "fmt"

func main() {
	// example 1
	// simple for loop
	//sum := 0
	//for i := 0; i < 10; i++ {
	//	sum += 1
	//}

	//example 2
	// the init and post statements are optional | for continued
	//sum := 1
	//for ; sum < 1000; {
	//	sum += sum
	//}

	//example 3
	//you can use for instead of while in C program language
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	// example 4
	// Infinite loop
	// do not run this example
	//for{
	//	sum += sum
	//}

	fmt.Println(sum)
}
