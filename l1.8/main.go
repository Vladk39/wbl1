package main

import "fmt"

func SetBit(n int64, i uint, value bool) int64 {
	if value {
		return n | (1 << i)
	} else {
		return n &^ (1 << i)
	}
}

func main() {
	var n int64 = 5

	n = SetBit(n, 0, false)

	n = SetBit(n, 2, true)
	fmt.Println(n)

}
