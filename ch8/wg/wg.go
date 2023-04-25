package wg

import (
	"fmt"
	"sync"
	"time"
)

func Run() {
	var wg sync.WaitGroup
	f := func(i int) {
		defer wg.Done()
		time.Sleep(5 * time.Second)
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f(i)
	}

	// ждать горутины
	wg.Wait()
}
