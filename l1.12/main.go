package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(uniq(s))
}

func uniq(s []string) []string {
	sort.Strings(s)

	result := []string{}

	prev := ""

	for i, v := range s {
		if i == 0 || v != prev {
			result = append(result, v)
		}
		prev = v
	}

	return result
}

// func sortstr(s []string) []string {
// 	result := []string{}
// 	m := make(map[string]struct{})

// 	for _, v := range s {
// 		if _, ok := m[v]; !ok {
// 			result = append(result, v)
// 		}

// 		m[v] = struct{}{}

// 	}

// 	return result
// }
