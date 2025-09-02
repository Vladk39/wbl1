package main

import (
	"fmt"
	"time"
)

func writer(nums []int, ch chan int) {
	for _, n := range nums {
		ch <- n
	}
	close(ch)
}

func worker(chint, resch chan int) {
	for v := range chint {
		resch <- v * 2
	}
	close(resch)
}
func result(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	chint := make(chan int)
	resch := make(chan int)
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	go writer(nums, chint)
	go worker(chint, resch)
	result(resch)

	time.Sleep(1 * time.Second)
}
