package main

import (
	"bytes"
	"fmt"
)

func main() {
	ints := []int{1, 2, 3, 4, 5, 6}

	b := bytes.Buffer{}
	b.WriteByte('[')
	for i, v := range ints {
		if i > 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "%d", v)
	}
	b.WriteByte(']')

	println(b.String())
}
