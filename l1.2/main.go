package main

import (
	"fmt"
	"sync"
)

func main() {
	s := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, value := range s {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			fmt.Println(v * v)
		}(value)
	}

	wg.Wait()

}
