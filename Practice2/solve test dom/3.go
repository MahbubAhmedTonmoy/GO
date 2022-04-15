package main

import (
	"fmt"
	"math"
)

func findRoots(a, b, c float64) (float64, float64) {
	t := math.Sqrt(math.Pow(b, 2)-4*a*c)
	x1 := (-b + t) / (2 * a)
	x2 := (-b - t) / (2 * a)

	return x1, x2
}

func main() {
	x1, x2 := findRoots(2, 10, 8)
	fmt.Printf("Roots: %f, %f", x1, x2)
}
