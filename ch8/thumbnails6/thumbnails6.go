package thumbnails6

import (
	"fmt"
	"sync"
)

func imageFile(infile string) string {
	// сделать превью из infile, положить в той же директории и вернуть имя файла превью
	return ""
}

func makeThumbnails(filenames <-chan string) int64 {
	var wg sync.WaitGroup
	sizes := make(chan int64)
	for f := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			imageFile(<-filenames)
			var s int64 = 1 //todo:вычеслить реальный размер созданного превью
			sizes <- s
		}(f)
	}

	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}

	return total
}

func Run() {
	filenames := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			filenames <- fmt.Sprintf("image%d.png", i)
		}
		close(filenames)
	}()

	total := makeThumbnails(filenames)
	fmt.Println(total)
}
