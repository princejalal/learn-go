package main

import "fmt"

// https://go.dev/tour/moretypes/1
func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	//example 1
	//i, j := 42, 2701
	//p := &i
	//
	//fmt.Println(*p, j)

	//example 2
	i := 2
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
