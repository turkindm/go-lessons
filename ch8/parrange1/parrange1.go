package parrange1

import (
	"fmt"
	"time"
)

func Run() {
	start := time.Now()

	count := 10
	ch := make(chan struct{}, count)
	for i := 0; i < count; i++ {
		go func() {
			n := 2
			fmt.Printf("Sleeping %d seconds...\n", n)
			time.Sleep(time.Duration(n) * time.Second)
			ch <- struct{}{}
			fmt.Println("Here")
		}()
	}
	for i := 0; i < count; i++ {
		<-ch
	}
	fmt.Println("Done")
	fmt.Printf("Time: %d\n", time.Since(start))
}
