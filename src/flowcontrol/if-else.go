package main

import (
	"fmt"
	"math"
)

//func powForElseExample(x, n, lim float64) float64 {
//	if v := math.Pow(x, n); v < lim {
//		return v
//	} else {
//		fmt.Printf("%g >= %g\n", v, lim)
//	}
//	return lim
//}

func main() {
	if num := -5; num < 0 {
		fmt.Println(num, "is negative")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	if i := 12; i%2 == 0 {
		fmt.Println(i, "is even")
	} else {
		fmt.Println(i, "is odd")
	}

	fmt.Println(
		math.Pow(3, 2),
		//powForElseExample(3, 2, 10),
		//powForElseExample(3, 3, 20),
	)
}
