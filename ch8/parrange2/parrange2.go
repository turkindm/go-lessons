package parrange2

import (
	"fmt"
	"sync"
	"time"
)

func Run(files <-chan int) {
	start := time.Now()

	var wg sync.WaitGroup
	ch := make(chan struct{})
	for range files {
		wg.Add(1)
		go func() {
			defer wg.Done()
			n := 1
			fmt.Printf("Sleeping %d seconds...\n", n)
			time.Sleep(time.Duration(n) * time.Second)
			ch <- struct{}{}
			fmt.Println("Here")
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	fmt.Println("Done")
	fmt.Printf("Time: %d\n", time.Since(start))
}
