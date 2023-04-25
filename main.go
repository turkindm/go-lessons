package main

import "fmt"

type A struct {
}

type Foo interface {
}

func main() {
	var a *A
	var f Foo = a
	fmt.Println(f, f == (Foo)(nil))
}
