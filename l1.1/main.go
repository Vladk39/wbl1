package main

import "fmt"

type human struct {
	Action
}

type Action struct {
}

func (A *Action) Print() {
	fmt.Println("qq")
}

func main() {
	h := human{}

	h.Print()
}
