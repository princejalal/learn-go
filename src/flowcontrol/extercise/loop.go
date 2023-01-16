package main

// source: https://go.dev/tour/flowcontrol/8
import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	// This is a simple version. For large numbers, more accurate calculations should be used
	// more answers  https://gist.github.com/tetsuok/2279991
	z := float64(1) // 1.0
	for n := 1.0; n < 10; n++ {
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func main() {
	//x := float64(1250000123123)
	x := float64(25)
	fmt.Println(Sqrt(x))      // newton's method
	fmt.Println(math.Sqrt(x)) // square root function
}
