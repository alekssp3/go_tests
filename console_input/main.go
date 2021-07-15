package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cmd () string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	text, _ := reader.ReadString('\n')
	return text[:len(text) - 1]
}

//func exit(a bool) {
//	a = false
//}
//
//var commands = map[string]func {
//	"exit" : exit,
//}

func Cmder() {
	var c = ""
	//exit := "exit quit"
	for c != "exit" {
		c = cmd()
	}
}

func main() {
	//Cmder()

	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("1 >> ")
	//text, _, _ := reader.ReadLine()
	//fmt.Printf("You input: %v. Ok?\n", text)
	//fmt.Print("2 >> ")
	//text2, _ := reader.ReadString('\n')
	//fmt.Printf("You input: %v. Ok?", text2[:len(text2) - 1])

	fmt.Println(strings.Contains("ad ad", ""))
	fmt.Println(strings.Count("ad ad", ""))
}
