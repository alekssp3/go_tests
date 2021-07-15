package main

import (
	"fmt"
	ll "go_tests/linkedlist"
)

func linkedListSum(l *ll.LinkedList) int {
	var res int
	for {
		if l.Next != nil {
			res += l.Val
			l = l.Next
		} else {
			res += l.Val
			break
		}
	}
	return res
}

func main(){
	var l = &ll.LinkedList{Val: 1, Next: &ll.LinkedList{Val: 2, Next: &ll.LinkedList{Val: 4}}}
	fmt.Println(linkedListSum(l))
}

