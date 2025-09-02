package main

import "fmt"

func main() {
	s := []int{8, 3, 7, 4, 9, 2}

	sorted := quickSort(s)

	fmt.Println(sorted)
}

func quickSort(s []int) []int {
	if len(s) < 2 {
		return s
	}

	pivot := s[len(s)/2]
	i, j := 0, len(s)-1

	for i <= j {
		for s[i] < pivot {
			i++
		}
		for s[j] > pivot {
			j--
		}
		if i <= j {
			s[i], s[j] = s[j], s[i]
			i++
			j--
		}
	}

	if j > 0 {
		quickSort(s[:j+1])
	}
	quickSort(s[i:])

	return s
}
