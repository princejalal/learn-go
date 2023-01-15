package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*y + y*y))
	var z uint = uint(f)
	var i int = 42
	var e float64 = float64(i)
	u := uint(e)
	fmt.Println(x, y, z, e, u)
}
