package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type counter struct {
	count int64
}

func (c *counter) inc() {
	atomic.AddInt64(&c.count, 1)
}

func main() {
	c := counter{}
	for i := 0; i < 100; i++ {
		go c.inc()
	}
	time.Sleep(1 * time.Second)

	fmt.Println(c.count)
}
