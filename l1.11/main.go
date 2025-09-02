package main

import (
	"fmt"
	"sort"
)

func main() {
	A := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	B := []int{2, 5, 7}

	fmt.Println(intersection(A, B))
}

func intersection(slice1, slice2 []int) []int {
	result := []int{}

	sort.Ints(slice2)

	for _, x := range slice1 {
		idx := sort.SearchInts(slice1, x)
		if idx < len(slice1) && slice1[idx] == x {
			result = append(result, x)
		}
	}

	return result
}

// func intersection(slice1, slice2 []int) []int {
// 	result := []int{}

// 	sort.Ints(slice1)
// 	sort.Ints(slice2)
// 	fmt.Println(slice1)
// 	fmt.Println(slice2)

// 	for i, j := 0, 0; i < len(slice1) && j < len(slice2); {
// 		if slice1[i] == slice2[j] {
// 			result = append(result, slice1[i])
// 			i++
// 			j++
// 		} else if slice1[i] < slice2[j] {
// 			i++
// 		} else {
// 			j++
// 		}

// 	}

// 	return result
// }

// func intersection(slice1, slice2 []int) []int {
// 	result := []int{}
// 	m := make(map[int]struct{})

// 	for _, v := range slice1 {
// 		m[v] = struct{}{}
// 	}

// 	for _, v := range slice2 {
// 		if _, ok := m[v]; ok {
// 			result = append(result, v)
// 		}
// 	}
// 	return result
// }
