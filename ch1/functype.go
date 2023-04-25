package main

import (
	"fmt"
	"math"
)

func mlt(a, b int) int {
	return a * b
}

func sum(a, b int) int {
	return a + b
}

func pow(a float64, b int) float64 {
	return math.Pow(a, float64(b))
}

func main() {
	f := mlt
	fmt.Println(f(10, 20))

	f = pow //так нельзя, разные параметры - разные типы
	fmt.Println(f(10, 20))
}
