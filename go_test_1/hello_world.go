package main

import "fmt"

func main() {
	var t_int int
	var t_float float64
	var t_string string
	var t_bool bool
	var t_arr_int [3]int
	t_arr2_int := [...]int{1, 2, 3}
	var t_mat_int [3][3]int
	var t_slice []int

	fmt.Println(t_int, t_float, t_string, t_bool, t_arr_int, t_arr2_int, t_mat_int, t_slice)
	println(len(t_slice))
	println("t_slice", cap(t_slice), "t_arr2_int", cap(t_arr2_int))
}
