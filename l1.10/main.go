package main

import (
	"fmt"
)

func main() {
	Temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	fmt.Println(group(Temp))
}

func group(temp []float64) map[int][]float64 {
	result := make(map[int][]float64)

	for _, t := range temp {

		key := int(t/10) * 10
		fmt.Println(key)

		result[key] = append(result[key], t)
	}
	return result
}
