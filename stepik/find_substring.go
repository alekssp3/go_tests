package main

import (
	"fmt"
	"strings"
)


func main() {
	var a, b string
	fmt.Scan(&a, &b)
	fmt.Print(strings.Index(a, b))
}