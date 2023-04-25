package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	*c += WordCounter(count)
	return count, nil
}

func main() {
	var c WordCounter
	p := "   😎Hello, world 1111"
	fmt.Fprintf(&c, "%s\n", p)
	fmt.Println(c)
}
