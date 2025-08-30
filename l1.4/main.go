package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Printf("chanel closed")
				return
			}
			fmt.Printf("worker %d, job %v\n", id, v)
		case <-ctx.Done():
			fmt.Println("worker stopped")
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()
	var wg sync.WaitGroup
	numw := 5
	ch := make(chan int)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	for i := 0; i < numw; i++ {
		wg.Add(1)
		go worker(ctx, i, ch, &wg)
	}

	for i := 0; i < 111564; i++ {
		ch <- i

	}
	close(ch)
	time.Sleep(333 * time.Millisecond)
	<-sig
	fmt.Println("signal")
	cancel()
	wg.Wait()

	fmt.Println("done")

}
