package main

import (
	"fmt"
)

const (
	one int = 2 * iota
	two
	three
)

func main() {
	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(three)
}
