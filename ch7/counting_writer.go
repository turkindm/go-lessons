package main

import (
	"bytes"
	"fmt"
	"io"
)

type CWriter struct {
	Wrappee io.Writer
	Counter int64
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	writer := &CWriter{Wrappee: w}

	return writer, &writer.Counter
}

func (c *CWriter) Write(p []byte) (n int, err error) {
	c.Counter += int64(len(p))
	return c.Wrappee.Write(p)
}

func main() {
	buf := new(bytes.Buffer)
	w, c := CountingWriter(buf)
	fmt.Fprintln(w, "Test")
	fmt.Println(*c)
}
