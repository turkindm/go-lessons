package main

import "strings"

func main() {
	s := "path/to/file.txt"
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	println(s)
}
