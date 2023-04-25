package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs slapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	res, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(io.Discard, res.Body)
	res.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("error while reading %s: %v", url, err)
		return
	}

	ch <- fmt.Sprintf("%.2fs %7d %s", time.Since(start).Seconds(), nbytes, url)
}
