package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	data := map[int]int{
		1: 1,
		2: 2,
	}

	for i := 0; i < 15; i++ {
		wg.Add(1)
		go worker(data, i, &mu, &wg)
	}
	wg.Wait()
	fmt.Println(data)

}

func worker(data map[int]int, i int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	data[i] = i + 1
	mu.Unlock()
}
