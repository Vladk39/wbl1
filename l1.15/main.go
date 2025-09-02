package main

var justString string

// func someFunc() {
// //   v := createHugeString(1 &lt;&lt; 10)
//   justString = v[:100]
// }

func main() {
	// someFunc()
}

// короче мы создаем большую строку v, А потом кберем первые 100 байт, но под капотом мы будем указывать на большой массив, и гс не сможет нормально очистить память
