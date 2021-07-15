package main

import (
	. "fmt"
	. "strings"
)


func main() {
	var s string
	Scan(&s)
	Print(Join(Split(s, ""), "*"))
}