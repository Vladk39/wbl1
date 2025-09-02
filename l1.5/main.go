package main

import (
	"fmt"
	"time"
)

func writer(ch chan int) {
	i := 0
	for {
		ch <- i
		i++
		time.Sleep(333 * time.Millisecond)
	}
}

func reader(n int) {
	ch := make(chan int)
	go writer(ch)

	time := time.After(time.Duration(n) * time.Second)

	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-time:
			fmt.Println("timeout")
			return
		}
	}

}

func main() {
	reader(4)
}
