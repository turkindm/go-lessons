package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	res, err := MakeRequest(ctx)
	if err != nil {
		log.Fatal("request timeout")
	}

	fmt.Println(string(res))
}

func MakeRequest(ctx context.Context) ([]byte, error) {
	res := make([]byte, 0)
	ch := make(chan []byte, 0)

	go doRequest(ch)

	select {
	case <-ctx.Done():
		return []byte{}, ctx.Err()
	case res = <-ch:
		return res, nil
	}
}

func doRequest(res chan []byte) {
	time.Sleep(1 * time.Second)
	res <- []byte("response body")
}
