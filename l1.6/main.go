package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	chint := make(chan int)
	quit := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	times := time.After(time.Duration(3) * time.Second)

	go func() {
		for {
			select {

			case <-times:
				fmt.Println("timeout")
				return
			default:
				fmt.Println(" i am work(timeout)")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stopped ctx")
				return
			default:
				fmt.Println(" i am work(ctx)")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	go func() {
		fmt.Println("Kill gorutine")
		runtime.Goexit()

	}()

	go func() {
		for v := range chint {
			fmt.Println("info drom int channel", v)
		}
	}()
	time.Sleep(500 * time.Millisecond)
	chint <- 1
	chint <- 2
	close(chint)

	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("stopped ne quit")
				return
			default:
				fmt.Println(" i am work(quit)")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	time.Sleep(2 * time.Second)
	quit <- struct{}{}

	time.Sleep(3 * time.Second)
}
