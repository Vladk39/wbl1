package main

import "fmt"

func main() {
	a := 1

	fmt.Println(chechtype(a))
}

func chechtype(a any) string {

	switch a.(type) {
	case string:
		return "string"
	case bool:
		return "bool"
	case int:
		return "int"
	case chan any:
		return "channel"
	default:
		return "cant find type"
	}

}
