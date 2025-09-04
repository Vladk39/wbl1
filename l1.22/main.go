package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(2)
	aexp := big.NewInt(33)
	resultA := new(big.Int).Exp(a, aexp, nil)
	fmt.Println(resultA)

	b := big.NewInt(2)
	bexp := big.NewInt(33)
	resultB := new(big.Int).Exp(b, bexp, nil)

	z := big.NewInt(0)

	z.Add(resultA, resultB)

	fmt.Println(z)
	z.Mul(resultA, resultB)
	fmt.Println(z)
	z.Sub(resultA, resultB)
	fmt.Println(z)

}
