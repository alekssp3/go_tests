package main

import (
	"fmt"
	"regexp"
)


func main() {
	var a string
	fmt.Scan(&a)
	match, _ := regexp.MatchString("^[a-zA-Z0-9]{5,}$", a)
	if match {
		fmt.Print("Ok")
	} else {
		fmt.Print("Wrong password")
	}
}



// package main

// import (
// 	"fmt"
// 	"unicode"
// )

// func main() {
// 	var pass string
// 	fmt.Scan(&pass)

// 	if isPasswordValid(pass) {
// 		fmt.Printf("Ok")
// 	} else {
// 		fmt.Printf("Wrong password")
// 	}
// }

// func isPasswordValid(password string) bool {
// 	if len([]rune(password)) < 5 {
// 		return false
// 	}

// 	for _, char := range password {
// 		if !unicode.IsDigit(char) && !unicode.Is(unicode.Latin, char) {
// 			return false
// 		}
// 	}
// }