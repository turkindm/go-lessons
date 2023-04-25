package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello, 😎"
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\t%d\n", i, r, size)
		i += size
	}
}
