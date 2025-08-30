package main

import (
	"fmt"
	"time"
)

// func workerpool(ch chan int, numw int) {
// 	var wg sync.WaitGroup

// 	for i := 0; i < numw; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			for v := range ch {
// 				fmt.Println(v)
// 			}
// 			wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()
// }

// func main() {
// 	ch := make(chan int)

// 	go workerpool(ch, 2)
// 	ch <- 42
// 	ch <- 22
// 	ch <- 1111
// 	close(ch)

// }

func worker(id int, ch <-chan int) {
	for v := range ch {
		fmt.Printf("worker %d, job %v", id, v)
	}
}

func main() {

	numw := 5
	ch := make(chan int)

	for i := 0; i < numw; i++ {
		go worker(i, ch)
	}

	for i := 0; i < 1111; i++ {
		ch <- i

	}
	time.Sleep(1 * time.Second)

}
