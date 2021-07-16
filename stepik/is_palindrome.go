package main

import (
	"fmt"
)


func main() {
	var a string
	fmt.Scan(&a)
	var s = []rune(a)
	var s2 = []rune(a)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	if string(s) == string(s2) {
		fmt.Print("Палиндром")
	} else {
		fmt.Print("Нет")
	}
}



// package main

// import "fmt"

// func main() {
// 	var s, r string
// 	fmt.Scan(&s)
// 	for _, i := range s {
// 		r = string(i) + r
// 	}
// 	if s == r {
// 		fmt.Println("Палиндром")
// 	} else {
// 		fmt.Println("Нет")
// 	}
// }