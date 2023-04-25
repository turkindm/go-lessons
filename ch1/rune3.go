package main

import (
	"fmt"
)

func main() {
	s := "hello, 😎"
	for i, r := range s {
		fmt.Printf("%d\t%q\t%x\n", i, r, r)
	}
}
