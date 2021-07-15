package main

import (
	"fmt"
	"go_tests/stack"
)

func main() {
	m := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
	inp := "(){}[]"
	s := new(stack.Stack)
	fmt.Printf("New: %T - %v\n", s, s)
	for i, ch := range inp {
		if m[string(ch)] == "" {
			if m[s.Last()] == string(ch) {
				_ = s.Pop()
			}
		} else {
			s.Push(string(ch))
		}
		// fmt.Printf("%T - %v => %t\n", m[string(ch)], m[string(ch)], m[string(ch)] == "")
		fmt.Printf("Iter %d: %T - %v\n", i, s, s)
	}
	fmt.Printf("End: %T - %v\n", s, s)
	fmt.Printf("%T - %v\n", m, m)
}
