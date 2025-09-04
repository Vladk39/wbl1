package main

import "fmt"

func main() {
	fmt.Println(turnstr("главрыба"))
}

// func turnstr(str string) string {

// 	sl := []rune(str)

// 	result := make([]rune, len(sl))

// 	for i, j := 0, len(sl)-1; i < len(sl); i, j = i+1, j-1 {
// 		result[i] = sl[j]

// 	}
// 	return string(result)
// }

func turnstr(str string) string {
	sl := []rune(str)

	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}

	return string(sl)
}
