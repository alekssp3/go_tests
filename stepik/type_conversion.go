package main

import (
	"fmt"
	"strconv"
)


func main() {
	// fmt.Print("%d\n", uint64(math.MaxUint64))
	fmt.Print([]rune("stepik"))
	fmt.Print(string([]byte("stepik")))
	fmt.Print(strconv.FormatBool(10 == int16(float64(100/10))))
	// fmt.Print((strconv.FormatBool(true)) == (10 == int(float64(100/10)))) //error
}