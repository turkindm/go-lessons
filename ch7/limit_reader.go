package main

import (
	"fmt"
	"io"
	"strings"
)

type LimitedReader struct {
	R io.Reader
	N int64
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{R: r, N: n}
}

func (r *LimitedReader) Read(p []byte) (n int, err error) {
	if r.N <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > r.N {
		p = p[0:r.N]
	}
	n, err = r.R.Read(p)
	r.N -= int64(n)
	return
}

func main() {
	r := strings.NewReader("hello")
	lr := LimitReader(r, 4)
	buf := make([]byte, 2)
	lr.Read(buf)
	fmt.Println(string(buf))
	lr.Read(buf)
	fmt.Println(string(buf))
}
