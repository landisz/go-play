package main

import (
	"fmt"
)
// Sqrt is to calculate the square root of x by approaching the result in a loop.
func Sqrt(x float64) float64 {
	var z float64
	i := 1
	z = 2
	for i<1000 {
		z -= (z*z - x) / (2*z)
		i++
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
